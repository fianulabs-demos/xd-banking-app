package main

import (
	"context"
	"database/sql"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	user     = "dbuser"
	password = "s3cretp4ssword"
)

func connect() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/pqgotest", user, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}
	return db
}

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
}

func noah() {
	log.Println("test")
}

func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	token := ""
	for i := 0; i < 10; i++ {
		token += string(rand.Intn(26) + 'a')
	}
	return token
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
	token := generateToken()
	log.Printf("yoour secret is: %s\n", token)
	test()
	noah()
	connect()

	log.Println(os.Getenv("GITHUB_ACCESS_TOKEN"))

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, receive))
}
