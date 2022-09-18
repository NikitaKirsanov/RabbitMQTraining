package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io/ioutil"
)


func main(){
	/*
	* Creating new GorillaMux router for REST API
	* Создаем новый роутер GorillaMux для REST API
	*/
	r := mux.NewRouter()
	r.HandleFunc("/msg", Publish).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
	
}

func Publish(resp http.ResponseWriter,r *http.Request) {
	/*
	* Creation of connection to RabbitMQ instance
	* Создаем подключение к инстансу RabbitMQ
	*/
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil{
		fmt.Println(fmt.Sprintf("Connection to RabbitMQ failed: %v", err))
		panic(err)
	}
	defer conn.Close()

	/*
	* Declaring amqp channel to interract with queue
	* Объявляем amqp канал для взаимодействия с очередью
	*/
	ch, err := conn.Channel()
	if err != nil{
		fmt.Println(fmt.Sprintf("Declaring channel failed: %v", err))
		panic(err)
	}
	defer ch.Close()	

	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("Message reading  failed: %v", err))
	}

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
			Body: []byte(msg),
		},
	)
	if err != nil{
		fmt.Println(fmt.Sprintf("Sending message to queue err: %v", err))
		panic(err)
	}

}