# Kapeta Readme
This file contains some structural information about this service.

This file will be overwritten every time you change the service definition in Kapeta.

## Structure
This service is structured as follows:
* ```generated```: Contains generated code that shouldn't be edited directly
* ```pkg```: Contains your own and generated code that you can and should edit

In the sub folder structure you'll find the following:
* ```generated/entities```: Contains the entities used by the service.
* ```generated/queues```: Contains anything related to queues
* ```generated/rest_api```: Contains the REST API routes.
* ```generated/services```: Contains the REST API service interfaces.
* ```pkg/services```: Contains the services used by the REST API routes. This is where you should add your business logic

## REST API 
To edit the REST API handlers edit the services found here:
[pkg/services](pkg/services/)

The REST layer itself is generated for you - so your service
will be called as specified within the REST API definition in Kapeta.

You just need to worry about the logic.

The service files will only be generated if they don't already exist - or if they have not
changed since the last time they were generated.

## RabbitMQ Subscriber
To consume messages from a RabbitMQ queue a consumer is generated for you for each resource.

Use the constructor function to create a new consumer.

Below is an example of how to use the consumer to listen for messages on the queue:
```go
package main

import (
  kapeta "github.com/kapetacom/sdk-go-config"
  "github.com/kapetacom/my-project/generated/entities"
  "github.com/kapetacom/my-project/generated/queues"
  "github.com/rabbitmq/amqp091-go"
)

main() {
    config, err := kapeta.Init(".")
    if err != nil {
        panic(err)
    }
	
    eventsConsumer, err := queues.CreateEventsConsumer(config, 
        func(message entities.Event, delivery amqp091.Delivery) (rabbitmq.Action, error) {
            data, err := json.Marshal(message)
            if err != nil {
                return rabbitmq.NackDiscard, err
            }
            log.Printf("Received event message: %s", data)
            return rabbitmq.Ack, nil
        }
    )
	
    if err != nil {
        panic(err)
    }
    defer eventsConsumer.Close()
	
}
```

