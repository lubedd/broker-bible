# Golang Message Broker Bible

## Как использовать
 Consumer
 ```
 
    // Подключение к брокеру
	consumer, err := bible.GetConsumerClient("127.0.0.1", "5300")
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	var id uint64
	// Запрос на обмен данными
	err = consumer.Send("a", domain.ConsumersOpenConnection, id)
	if err != nil {
		log.Fatalf("can not send %v", err)
	}

    // Чтение из канала и подтверждение получения сообщения
	for {
		resp, err := consumer.Recv()
		if err != nil {
			continue
		}
		id = resp.Id
		DoSomething(resp.Message)

		err = consumer.Send(resp.RoutingKey, domain.ConsumersTaskAccepted, id)
		if err != nil {
			continue
		}
	}

```

Consumer interface
```
    type Consumer interface {
        Send(routerKey, message string, id uint64) error
        Recv() (*pb.Consumer, error)
    }
```

### Producer
```
    // Подключение к брокеру
	producer, err := bible.GetProducerClient("127.0.0.1", "5300")
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
    // Отправление новой задачи
	response, err := producer.NewTask("a", "Very important message")
	if err != nil || response.Message != domain.ProducerTaskAccepted {
		log.Fatalf("can not create new task %v", err)
	}
```
Producer interface
```
    type Producer interface {
        NewTask(routingKey, messageText string) (*pb.ResponseProducer, error)
    }
```