package task

import (
	"fmt"
	"task-manager/menu"
	"time"
)

type Dependencies struct {
	AddTask func(nextId int) (Task, error)
	ListTasks func([]Task)
	CompleteTask func(*[]Task)
}



func Handler(option string, tasks *[]Task, deps Dependencies) {

	switch option {
	case menu.AddTask:
		newTask, err := deps.AddTask(len(*tasks) + 1)
		if err == nil {
			*tasks = append(*tasks, newTask)
			fmt.Println("✅ Tarea agregada correctamente")
			time.Sleep(1 * time.Second)
		}
	case menu.ListTasks:
		deps.ListTasks(*tasks)
		fmt.Println(" ")
		fmt.Println("Presione enter para continuar")
		fmt.Scanln()
	case menu.CompleteTask:
		deps.CompleteTask(tasks)
	}
}

