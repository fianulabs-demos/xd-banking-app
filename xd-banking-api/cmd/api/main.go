package main

import (
	"context"
	"database/sql"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/miekg/dns"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

// createJWT creates a new JWT token using jwt-go
func createJWT() string {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "example_user"
	claims["exp"] = 15000

	// Sign and get the complete encoded token as a string using a secret key
	tokenString, _ := token.SignedString([]byte("mysecretkey"))
	return tokenString
}

// parseJWT parses a JWT token string
func parseJWT(tokenString string) (*jwt.Token, error) {
	// Parse the token
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecretkey"), nil
	})
}

func main() {

	// Example 1: Using jwt-go to create and parse a token
	token := createJWT()
	fmt.Println("Generated JWT:", token)

	parsedToken, err := parseJWT(token)
	if err != nil {
		log.Fatalf("Error parsing JWT: %v", err)
	}
	fmt.Println("Parsed JWT:", parsedToken)

	// Example 2: Using gin-gonic/gin to create a simple web server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go r.Run(":8080") // Run Gin server in a goroutine

	// Example 3: Using satori/go.uuid to generate a UUID
	id := uuid.NewV4()
	fmt.Println("Generated UUID:", id)

	// Example 4: Using miekg/dns to create a simple DNS message
	msg := new(dns.Msg)
	msg.SetQuestion("example.com.", dns.TypeA)
	fmt.Printf("Created DNS message: %v\n", msg)

	// Example 5: Using golang.org/x/crypto/bcrypt to hash a password
	password := "mysecretpassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	fmt.Println("Hashed password:", string(hashedPassword))

	// Keep the server running
	select {}

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
	token = generateToken()
	log.Printf("yoour secret is: %s\n", token)
	test()
	noah()
	connect()

	log.Println(os.Getenv("GITHUB_ACCESS_TOKEN"))

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, receive))
}
