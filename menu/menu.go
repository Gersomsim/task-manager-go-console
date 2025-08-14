package menu

import (
	"errors"
	"fmt"
)

const (
	AddTask = 1
	ListTasks = 2
	CompleteTask = 3
	Exit = 4
)

func ShowMenu() (int, error) {
	var option int
	fmt.Println("--------------------------------")
	fmt.Println(" ")
	fmt.Println("Seleccione una opción:")
	fmt.Println("1. 💾 Agregar tarea")
	fmt.Println("2. 📝 Listar tareas")
	fmt.Println("3. ✅ Marcar tarea como completada")
	fmt.Println("4. 🚪 Salir")
	fmt.Println(" ")
	fmt.Println("--------------------------------")
	fmt.Printf("Ingrese : ")
	// Verificamos que el tipo de dato sea int
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("La opcion ingresada no es valida, debe ser un numero del 1 al 4")
	}

	if option < AddTask || option > Exit {
		return 0, errors.New("La opcion ingresada no es valida")
	}

	return option, nil
}