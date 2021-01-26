// Example function-based high-level Apache Kafka consumer
package main
// go run consumer.go localhost:9092 zgroup senz
/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// consumer_example implements a consumer using the non-channel Poll() API
// to retrieve messages and events.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n",
			os.Args[0])
		os.Exit(1)
	}

	broker := os.Args[1]
	group := os.Args[2]
	topics := os.Args[3:]
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(topics, nil)

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/Shopify/sarama"
// 	"github.com/wvanbergen/kafka/consumergroup"
// )

// const (
// 	zookeeperConn = "localhost:2181"
// 	cgroup        = "zgroup"
// 	topic         = "senz"
// )

// func main() {
// 	// setup sarama log to stdout
// 	// sarama.Logger = log.New(os.Stdout, "", log.Ltime)
// 	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

// 	// init consumer
// 	cg, err := initConsumer()
// 	if err != nil {
// 		fmt.Println("Error consumer goup: ", err.Error())
// 		os.Exit(1)
// 	}
// 	defer cg.Close()

// 	// run consumer
// 	consume(cg)
// }

// func initConsumer() (*consumergroup.ConsumerGroup, error) {
// 	// consumer config
// 	config := consumergroup.NewConfig()
// 	config.Offsets.Initial = sarama.OffsetOldest
// 	config.Offsets.ProcessingTimeout = 10 * time.Second

// 	// join to consumer group
// 	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{topic}, []string{zookeeperConn}, config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cg, err
// }

// func consume(cg *consumergroup.ConsumerGroup) {
// 	for {
// 		select {
// 		case msg := <-cg.Messages():
// 			// messages coming through chanel
// 			// only take messages from subscribed topic
// 			if msg.Topic != topic {
// 				continue
// 			}

// 			fmt.Println("Topic: ", msg.Topic)
// 			fmt.Println("Value: ", string(msg.Value))

// 			// commit to zookeeper that message is read
// 			// this prevent read message multiple times after restart
// 			err := cg.CommitUpto(msg)
// 			if err != nil {
// 				fmt.Println("Error commit zookeeper: ", err.Error())
// 			}
// 		}
// 	}
// }
