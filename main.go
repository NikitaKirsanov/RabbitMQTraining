package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main(){
	/*
	* Creation of connection to RabbitMQ instance
	*/
	conn, err := amqp.Dial("amqp://guest:guest@yourhost:5672/")
	if err != nil{
		fmt.Println(fmt.Sprintf("Connection to RabbitMQ failed: %v", err))
		panic(err)
	}
	defer conn.Close()


	/*
	* Creation of Channel to interract with queue
	*/ 
	ch, err := conn.Channel()
	if err != nil{
		fmt.Println(fmt.Sprintf("Creation of channel err: %v", err))
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestingQueue",
		false,
		false,
		false, 
		false,
		nil,
	)
	if err != nil{
		fmt.Println(fmt.Sprintf("Creation of queue err: %v", err))
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestingQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("Here would be messages from the Client"),
		},
	)
	if err != nil{
		fmt.Println(fmt.Sprintf("Sending message to queue err: %v", err))
		panic(err)
	}
}