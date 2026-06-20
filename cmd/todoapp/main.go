package main

import (
	"fmt"
	"todoapp/internal/features/tasks/repository"
	"todoapp/internal/features/tasks/service"
)

func main() {
	repo := repository.NewTaskRepository("tasks.json")
	taskService := service.NewTaskService(repo)

	for {
		fmt.Println("\n1. Показать задачи\n2. Добавить задачу\n3. Выполнить задачу\n4. Выйти")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			list := taskService.GetTasks()
			fmt.Println("--- СПИСОК ЗАДАЧ ---")
			for i, task := range list {
				status := "[ ]"
				if task.IsDone {
					status = "[X]"
				}
				fmt.Printf("%d. %s %s\n", i, task.Title, status)
			}
		case 2:
			var taskTitle string
			fmt.Println("Введите название задачи:")
			fmt.Scan(&taskTitle)

			err := taskService.CreateTask(taskTitle)
			if err != nil {
				fmt.Println("❌ Ошибка:", err)
			} else {
				fmt.Println("🎉 Задача успешно добавлена!")
			}
		case 3:
			var id int
			fmt.Println("Введите номер задачи:")
			fmt.Scan(&id)

			err := taskService.CompleteTask(id)
			if err != nil {
				fmt.Println("❌ Ошибка:", err)
			} else {
				fmt.Println("🎉 Задача выполнена!")
			}
		case 4:
			fmt.Println("Пока!")
			return
		default:
			fmt.Println("Неверный ввод, попробуй еще раз.")
		}
	}
}
