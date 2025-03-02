package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YakovlevIgA/anotherTest/database"
	"github.com/YakovlevIgA/anotherTest/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.InitDB()
	app := fiber.New()
	router.SetupRoutes(app)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Запуск приложения...")
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()

	<-stop
	fmt.Println("Получен сигнал завершения работы. Откат миграции...")
	err := database.RollbackMigration()
	if err != nil {
		log.Fatalf("Ошибка при откате миграции: %v", err)
	}

	fmt.Println("Приложение завершило свою работу.")
}
