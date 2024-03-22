package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define mock routes related to banking

	// Mock route for creating an account
	app.Post("/api/account", func(c *fiber.Ctx) error {
		// Mock response for creating an account
		return c.SendString("Account created successfully")
	})

	// Mock route for getting account balance
	app.Get("/api/account/:id/balance", func(c *fiber.Ctx) error {
		// Mock response for getting account balance
		return c.SendString("Account balance: $1000")
	})

	// Mock route for depositing funds into an account
	app.Post("/api/account/:id/deposit", func(c *fiber.Ctx) error {
		// Mock response for depositing funds
		return c.SendString("Funds deposited successfully")
	})

	// Mock route for withdrawing funds from an account
	app.Post("/api/account/:id/withdraw", func(c *fiber.Ctx) error {
		// Mock response for withdrawing funds
		return c.SendString("Funds withdrawn successfully")
	})

	// Mock route for transferring funds between accounts
	app.Post("/api/transfer", func(c *fiber.Ctx) error {
		// Mock response for transferring funds
		return c.SendString("Funds transferred successfully")
	})

	// Mock route for getting transaction history of an account
	app.Get("/api/account/:id/transactions", func(c *fiber.Ctx) error {
		// Mock response for getting transaction history
		return c.SendString("Transaction history retrieved successfully")
	})

	// Add more mock routes as needed...

	// Start the server on port 3000
	app.Listen(":8080")
}
