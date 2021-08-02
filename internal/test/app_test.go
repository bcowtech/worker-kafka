package test

import (
	"context"
	"testing"
	"time"

	"github.com/bcowtech/config"
	kafka "github.com/bcowtech/worker-kafka"
)

func TestStarter(t *testing.T) {
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
