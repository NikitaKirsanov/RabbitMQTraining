package main

import(
	"fmt"
	"github.com/streadway/amqp"
)

func main(){

	conn, err := amqp.Dial("amqp://guest:guest@yourhost:5672/")
	if err != nil{
		fmt.Println(fmt.Sprintf("Connection to RabbitMQ failed: %v", err))
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		fmt.Println(fmt.Sprintf("Declaring channel failed: %v", err))
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestingQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		fmt.Println(fmt.Sprintf("Consuming messeges error: %v", err))
		panic(err)
	}

	forever := make(chan bool)
	go func(){
		for d := range msgs {
			fmt.Printf("Reseived message: %s\n", d.Body)
		}
	}()
	<- forever
}