package main

import (
	"fmt"
	"task-manager/menu"
	"task-manager/task"
	"time"
)

const filename = "tasks.json"

func main() {
	var option int
	var tasks []task.Task
	tasks, err := task.LoadFromFile(filename)
	if err != nil {
		fmt.Printf("Error al cargar las tareas: %v\n", err)
		time.Sleep(1 * time.Second)
		return
	}
	for option != 4 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Bienvenido al gestor de tareas")
		optionSelected, error := menu.ShowMenu()
		if error != nil {
			fmt.Printf("Error: %v\n", error)
			time.Sleep(1 * time.Second)
			continue
		}
		option = optionSelected
		task.Handler(option, &tasks)
	}
	err = task.SaveToFile(tasks, filename)
	if err != nil {
		fmt.Println("El trabajo se perderá")
		fmt.Printf("Error al guardar las tareas: %v\n", err)
		time.Sleep(1 * time.Second)
		return
	}

	fmt.Println("Gracias por usar el gestor de tareas")






	



}