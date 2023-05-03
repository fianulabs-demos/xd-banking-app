package main

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
)

func receive(ctx context.Context, event cloudevents.Event) {
	log.Print("ping\n")

	log.Println("antoher test dd")

	log.Printf("%s", event.Data())

	result, _ := process(event.Data()) // output of process function
	log.Printf("output\n%s\\n", result)
	kotest()
	test2()
}

func process(d []byte) ([]byte, error) {
	return d, nil
}

func kotest() {
	log.Println("some function")
}

func test2() {
	log.Println("noah")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("bing")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("noah")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
}

func test() {
	log.Println("noah")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("bing")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("noah")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("noah")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
	log.Println("test")
}

func test2() {
	log.Println("test")
}

func main() {

	ctx := context.Background()

	p, err := cloudevents.NewHTTP()

	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p)

	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("listening on :8080\n")
	log.Println("test")
	test()
	test2()
	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, receive))
}
