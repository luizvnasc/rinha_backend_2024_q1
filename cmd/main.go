package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var wg sync.WaitGroup
	var err error
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	shutdownError := make(chan error)

	go func() {

		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit
		log.Printf("shuting down the server: %s", s)

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

		defer cancel()

		err := app.ShutdownWithContext(ctx)
		if err != nil {
			shutdownError <- err
		}

		wg.Wait()

		shutdownError <- nil
	}()

	log.Println("Starting the server")

	if err = app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
	log.Println("Stopping the server")
}
