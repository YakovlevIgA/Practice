package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Ошибка конфигурации подключения: %v\n", err)
	}

	ctx := context.Background()
	DB, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v\n", err)
	}

	fmt.Println("База данных успешно подключена!")

	err = applyMigration(ctx)
	if err != nil {
		log.Fatalf("Ошибка при применении миграции: %v\n", err)
	}
}

func applyMigration(ctx context.Context) error {
	migrationSQL, err := os.ReadFile("migrations/01_table.up.sql")
	if err != nil {
		return fmt.Errorf("ошибка чтения миграции: %v", err)
	}

	_, err = DB.Exec(ctx, string(migrationSQL))
	if err != nil {
		return fmt.Errorf("ошибка выполнения миграции: %v", err)
	}

	fmt.Println("Миграция успешно выполнена!")
	return nil
}

func RollbackMigration() error {

	rollbackSQL, err := os.ReadFile("migrations/01_table.down.sql")
	if err != nil {
		return fmt.Errorf("ошибка чтения миграции для отката: %v", err)
	}

	_, err = DB.Exec(context.Background(), string(rollbackSQL))
	if err != nil {
		return fmt.Errorf("ошибка выполнения отката миграции: %v", err)
	}

	fmt.Println("Выполнен откат БД!")
	return nil
}
