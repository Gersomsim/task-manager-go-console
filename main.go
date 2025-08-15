package main

import (
	"fmt"
	"io"
	"os"
	"task-manager/internal/cli"
	"task-manager/menu"
	"task-manager/task"
	"time"
)

const filename = "tasks.json"

func main() {
	var option string
	var tasks []task.Task
	tasks, err := task.LoadFromFile(filename)
	if err != nil {
		fmt.Printf("Error al cargar las tareas: %v\n", err)
		time.Sleep(1 * time.Second)
		return
	}
	for option != menu.Exit {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Bienvenido al gestor de tareas")
		optionSelected, error := menu.ShowMenu(cli.Input)
		if error != nil {
			fmt.Printf("Error: %v\n", error)
			time.Sleep(1 * time.Second)
			continue
		}
		option = optionSelected
		task.Handler(option, &tasks, task.Dependencies{
			AddTask: task.AddTask,
			ListTasks: task.ListTasks,
			CompleteTask: task.CompleteTask,
		})
	}
	err = task.SaveToFile(tasks, filename, func(name string) (io.WriteCloser, error) {
		return os.Create(name)
	})
	if err != nil {
		fmt.Println("El trabajo se perderá")
		fmt.Printf("Error al guardar las tareas: %v\n", err)
		time.Sleep(1 * time.Second)
		return
	}

	fmt.Println("Gracias por usar el gestor de tareas")






	



}