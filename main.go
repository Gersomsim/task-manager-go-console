package main

import (
	"fmt"
	"task-manager/menu"
	"task-manager/task"
	"time"
)



func main() {
	var option int
	var tasks []task.Task
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
	fmt.Println("Gracias por usar el gestor de tareas")






	



}