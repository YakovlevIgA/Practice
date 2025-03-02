package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/YakovlevIgA/anotherTest/database"
	"github.com/YakovlevIgA/anotherTest/models"
	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {

	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	if task.Status == "" {
		task.Status = "new"
	}

	ctx := c.Context()
	query := `INSERT INTO tasks (title, description, status, created, updated) 
			VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, created, updated`
	err := database.DB.QueryRow(ctx, query, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.Created, &task.Updated)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Task successfully created",
		"Task":    task,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating task",
		})
		return err
	}
	return nil
}

func ShowTasks(c *fiber.Ctx) error {

	ctx := context.Background()

	rows, err := database.DB.Query(ctx, "SELECT id, title, description, status, created, updated FROM tasks ORDER BY id")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created, &task.Updated)
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
			return err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
		return err
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"tasks":   tasks,
		"message": "All tasks returned successfully",
	})
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	fmt.Println(id)
	commandTag, err := database.DB.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {

		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	log.Printf("Deleting task with id: %s", id)

	if commandTag.RowsAffected() == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Task deleted successfully",
	})
}
func EditTask(c *fiber.Ctx) error {
	id := c.Params("id") //
	ctx := context.Background()

	if id == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Task ID is required",
		})
	}
	fmt.Println("Got id!")
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if task.Title == "" || task.Description == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Title and Description are required",
		})
	}

	query := `
		UPDATE tasks 
		SET title = $1, description = $2, status = $3, updated = NOW()
		WHERE id = $4
	`
	commandTag, err := database.DB.Exec(ctx, query, task.Title, task.Description, task.Status, id)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to update task",
			"error":   err.Error(),
		})
	}

	if commandTag.RowsAffected() == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Task updated successfully",
	})
}
