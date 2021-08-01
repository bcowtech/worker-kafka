package test

import (
	"fmt"
	"time"

	kafka "github.com/bcowtech/worker-kafka"
)

type (
	App struct {
		Host            *Host
		Config          *Config
		ServiceProvider *ServiceProvider
	}

	Host kafka.Worker

	Config struct {
		// kafka
		BootstrapServers string        `env:"*KAFKA_BOOTSTRAP_SERVERS"   yaml:"-"`
		GroupID          string        `env:"-"                          yaml:"groupID"`
		PollingTimeout   time.Duration `env:"-"                          yaml:"pollingTimeout"`
	}

	ServiceProvider struct {
		ResourceName string
	}

	TopicGateway struct {
		*MyTopicMessageHandler   `topic:"myTopic"`
		*UnhandledMessageHandler `topic:"?"`
	}
)

func (provider *ServiceProvider) Init(conf *Config) {
	fmt.Println("ServiceProvider.Init()")
	provider.ResourceName = "demo resource"
}

func (h *Host) Init(conf *Config) {
	h.PollingTimeout = conf.PollingTimeout
	h.PingTimeout = time.Second * 3
	h.ConfigMap = &kafka.ConfigMap{
		"bootstrap.servers": conf.BootstrapServers,
		"group.id":          conf.GroupID,
		"auto.offset.reset": "earliest",
	}
}
