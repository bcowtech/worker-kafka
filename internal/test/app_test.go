package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/bcowtech/config"
	kafka "github.com/bcowtech/worker-kafka"
)

func TestStarter(t *testing.T) {
	err := setupTestStarter()
	if err != nil {
		t.Fatal(err)
	}

	var (
		KafkaNameFormatter kafka.NameTransformProc = kafka.NameFormatter("%s-${KAFKA_TOPIC_SUFFIX}")
	)

	app := App{}
	starter := kafka.Startup(&app).
		Middlewares(
			kafka.UseGroupIDTransformer(KafkaNameFormatter),
			kafka.UseTopicGateway(&TopicGateway{}).
				TopicTransformer(KafkaNameFormatter),
			kafka.UseErrorHandler(func(err kafka.Error) (disposed bool) {
				t.Logf("catch err: %v", err)
				return false
			}),
		).
		ConfigureConfiguration(func(service *config.ConfigurationService) {
			service.
				LoadEnvironmentVariables("").
				LoadYamlFile("config.yaml").
				LoadCommandArguments()
		})

	runCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := starter.Start(runCtx); err != nil {
		t.Error(err)
	}

	select {
	case <-runCtx.Done():
		if err := starter.Stop(context.Background()); err != nil {
			t.Error(err)
		}
	}

	// assert app.Config
	{
		conf := app.Config
		if conf.BootstrapServers == "" {
			t.Errorf("assert 'Config.BootstrapServers':: should not be an empty string")
		}
		var expectedGroupID string = "gotest-group"
		if conf.GroupID != expectedGroupID {
			t.Errorf("assert 'Config.GroupID':: expected '%v', got '%v'", expectedGroupID, conf.GroupID)
		}
		var expectedPollingTimeout time.Duration = 30 * time.Millisecond
		if conf.PollingTimeout != expectedPollingTimeout {
			t.Errorf("assert 'Config.PollingTimeout':: expected '%v', got '%v'", expectedPollingTimeout, conf.PollingTimeout)
		}
	}
}

func setupTestStarter() error {
	f, err := kafka.NewForwarder(&kafka.ForwarderOption{
		FlushTimeout: 3 * time.Second,
		PingTimeout:  3 * time.Second,
		ConfigMap: &kafka.ConfigMap{
			"client.id":         "gotest",
			"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		},
	})
	if err != nil {
		return err
	}
	defer f.Close()

	topic := "myTopic"
	realTopic := fmt.Sprintf(os.ExpandEnv("%s-${KAFKA_TOPIC_SUFFIX}"), topic)
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		f.WriteMessage(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &realTopic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	return nil
}
