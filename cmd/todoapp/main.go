package main

import (
	"context"
	"fmt"
	"os"
	"todoapp/internal/features/tasks/repository"
	"todoapp/internal/features/tasks/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbPass == "" {
		dbPass = "postgres"
	} // твой пароль от базы
	if dbName == "" {
		dbName = "todo_db"
	}

	// 2. Собираем строку подключения
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", dbUser, dbPass, dbName)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		fmt.Printf("Не удалось подключиться к базе: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// 4. Передаем пул в твой новый репозиторий
	repo := repository.NewTaskRepository(pool)
	taskService := service.NewTaskService(repo)

	for {
		fmt.Println("\n--- МЕНЮ ---")
		fmt.Println("1. Показать задачи")
		fmt.Println("2. Добавить задачу")
		fmt.Println("3. Выполнить задачу")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")

		var choice int
		fmt.Print("Выберите пункт: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			list, err := taskService.GetTasks()
			if err != nil {
				fmt.Println("❌ Не удалось загрузить задачи:", err)
				continue
			}

			fmt.Println("--- СПИСОК ЗАДАЧ ---")
			for i, task := range list {

				fmt.Printf("%d. %s (ID в базе: %d)\n", i+1, task.Title, task.ID)
			}

		case 2:
			var taskTitle string
			fmt.Print("Введите название задачи: ")
			fmt.Scan(&taskTitle)

			err := taskService.CreateTask(taskTitle)

			if err != nil {
				fmt.Println("Ошибка при добавлении задачи", err)
				continue
			}

			fmt.Print("Задача добавлена")
		case 3:
			list, err := taskService.GetTasks()
			if err != nil {
				fmt.Println("❌ Ошибка при загрузке данных:", err)
				continue
			}
			if len(list) == 0 {
				fmt.Println("Список пуст!")
				continue
			}

			var num int
			fmt.Print("Введите номер задачи: ")
			fmt.Scan(&num)

			idx := num - 1

			if idx < 0 || idx >= len(list) {
				fmt.Println("❌ Нет такой задачи в списке!")
				continue
			}

			targetID := list[idx].ID

			err = taskService.CompleteTask(targetID)

			if err != nil {
				fmt.Println("❌ Ошибка:", err)
			} else {
				fmt.Println("🎉 Задача выполнена!")
			}

		case 4:
			list, err := taskService.GetTasks()
			if err != nil {
				fmt.Println("❌ Ошибка при загрузке данных:", err)
				continue
			}
			if len(list) == 0 {
				fmt.Println("Список пуст!")
				continue
			}

			var num int
			fmt.Print("Введите номер задачи: ")
			fmt.Scan(&num)

			idx := num - 1

			if idx < 0 || idx >= len(list) {
				fmt.Println("❌ Нет такой задачи в списке!")
				continue
			}

			targetID := list[idx].ID

			err = taskService.DeleteTask(targetID)

			if err != nil {
				fmt.Println("❌ Ошибка:", err)
			} else {
				fmt.Println("🎉 Задача удалена!")
			}
		case 5:
			fmt.Println("Пока!")
			return

		default:
			fmt.Println("Неверный ввод, попробуй еще раз.")
		}
	}

}
