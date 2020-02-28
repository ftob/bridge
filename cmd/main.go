package cmd

import (
	"fmt"
	"github.com/ftob/bridge/server"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	db = viper.Get("BR_DB")
	kafka = viper.GetString("BR_KAFKA")
	kafkaGroupID = viper.GetString("BR_KAFKA_GROUP_ID")
	kafkaTopic = viper.GetString("BR_KAFKA_TOPIC")
	debug = viper.GetBool("BR_DEBUG")
	httpPort = viper.GetString("BR_HTTP")
)

func main() {
	// Configure variables
	viper.SetEnvPrefix("br")

	// Kafka config
	_ = viper.BindEnv("kafka")
	_ = viper.BindEnv("kafka_group_id")
	_ = viper.BindEnv("kafka_topic")
	_ = viper.BindEnv("db")
	viper.AutomaticEnv()

	var logger *zap.Logger

	if debug {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	record := make(chan string, 1)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": db,
		"group.id":          kafkaGroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{kafkaTopic}, nil)

	// start kafka
	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				record<-msg.Value
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()

	go func() {

	}()


	serve := server.New()

	errs := make(chan error, 2)

	logger.Info("start HTTP server")
	errs <- serve.ListenAndServe(httpPort)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Error("server terminated")
}
