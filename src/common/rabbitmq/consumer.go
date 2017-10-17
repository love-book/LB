// This example declares a durable Exchange, an ephemeral (auto-delete) Queue,
// binds the Queue to the Exchange with a binding key, and consumes every
// message published to that Exchange with that routing key.
//
package rabbitmq

import (
	"github.com/streadway/amqp"
	"fmt"
	"log"
)


type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
	Exchange     string
	ExchangeType string
	Queue        string
	BindingKey   string
	ConsumerTag  string
	Body         []byte
	HandlerTag   string
	HandlerRegister  *HandlerRegister
}

func (this *Consumer)Sub() {
	c, err := NewConsumer(uri, this)
	if err != nil {
		log.Fatalf("%s", err)
	}
	go func() {
	shuttdown:
		for {
			select {
			case  <-c.done:
				log.Println("shutting down ")
				break shuttdown
			}
		}
		log.Printf("shutting down")
		if err := c.Shutdown(); err != nil {
			log.Fatalf("error during shutdown: %s", err)
		}
	}()
}

func NewConsumer(amqpURI string,c *Consumer) (*Consumer, error) {
	// exchange, exchangeType, queueName, key, ctag
	exchange:= c.Exchange
	exchangeType:=c.ExchangeType
	queueName:= c.Exchange
	key := c.Exchange
	ctag := c.Exchange
	c.conn = nil
	c.channel = nil
	c.tag = ctag
	c.done = make(chan error)

	var err error

	log.Printf("dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	if err = c.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", queueName)
	queue, err := c.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, key)

	if err = c.channel.QueueBind(
		queue.Name, // name of the queue
		key,        // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queue.Name, // name
		c.tag,      // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, c)

	return c, nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

func handle(deliveries <-chan amqp.Delivery,c *Consumer) {
	for d := range deliveries {
		c.HandlerTag = d.ConsumerTag
		c.Body =d.Body
		Do(c)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	c.done <- nil
}


func Do(s *Consumer){
	err, handles := s.HandlerRegister.Get(s.HandlerTag)
	if err != nil {
		log.Println(err)
	}
	for _, v := range handles {
		go v.Run(s)
	}
}


