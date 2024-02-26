package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"rinha_backend2024_q1/internal/transacao"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var wg sync.WaitGroup
	var err error

	dsn := "host=localhost user=admin password=123 dbname=rinha port=5432 sslmode=disable TimeZone=Europe/Lisbon"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	transacao.NewStore(db)

	app := fiber.New()
	transacao.RegistraHandlers(app)
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
