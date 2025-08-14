package menu

import (
	"errors"
	"fmt"
	"task-manager/internal/cli"
)

const (
	AddTask = "1"
	ListTasks = "2"
	CompleteTask = "3"
	Exit = "4"
)

func ShowMenu(input cli.InputFunc) (string, error) {
	var option string
	validOptions := []string{AddTask, ListTasks, CompleteTask, Exit}
	fmt.Println("--------------------------------")
	fmt.Println(" ")
	fmt.Println("Seleccione una opción:")
	fmt.Println("1. 💾 Agregar tarea")
	fmt.Println("2. 📝 Listar tareas")
	fmt.Println("3. ✅ Marcar tarea como completada")
	fmt.Println("4. 🚪 Salir")
	fmt.Println(" ")
	fmt.Println("--------------------------------")

	option = input("Ingrese la opcion: ")

	for _, opt := range validOptions {
		if option == opt {
			return option, nil
		}
	}

	return "", errors.New("la opcion ingresada no es valida, debe ser un numero del 1 al 4")
}