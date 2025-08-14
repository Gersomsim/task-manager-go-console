package task

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"task-manager/internal/cli"
	"task-manager/menu"
	"text/tabwriter"
	"time"
)



func addTask(nextId int) (Task, error) {
	var task Task
	task.Id = nextId
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Agregar tarea %d o precione enter para salir \n", nextId)
	fmt.Println("--------------------------------")
	task.Title = cli.Input("Ingrese el titulo de la tarea: ")
	if task.Title == "" {
		return task, errors.New("título no válido")
	}
	task.Description = cli.Input("Ingrese la descripcion de la tarea: ")

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return task, nil
}

func listTasks(tasks []Task) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("---------------")
	fmt.Println("📝 LISTA DE TAREAS")
	fmt.Println("---------------")
	// Crear un escritor de tabla para formatear la salida en consola
	// Parámetros: minwidth=0, tabwidth=0, padding=2, padchar=' ', flags=Debug
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	
	// Escribir el encabezado de la tabla con columnas separadas por tabulaciones
	fmt.Fprintln(w, "ID\tTítulo\tDescripción\tCompletada\t")
	
	// Iterar sobre cada tarea para mostrar sus detalles
	for _, task := range tasks {
		// Determinar el estado de completado de la tarea
		completed := "No"
		if task.Completed {
			completed = "Sí"
		}
		
		// Escribir una fila de la tabla con los datos de la tarea
		// Formato: título, descripción, estado de completado
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n", task.Id, task.Title, task.Description, completed)
	}
	
	// Forzar la escritura de todos los datos pendientes en el buffer
	w.Flush()
}

func completeTask(tasks *[]Task) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Completar tarea")
	fmt.Println("--------------------------------")
	fmt.Printf("Ingrese el id de la tarea a completar: ")
	var id int
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
			fmt.Println("Error: ID inválido")
			time.Sleep(1 * time.Second)
			return
	}
	for idx, task := range *tasks {
		if task.Id == id {
			fmt.Printf("¿Está seguro de completar la tarea %s? s/n: ", task.Title)
			var confirm string
			fmt.Scanf("%s", &confirm)
			if strings.ToLower(confirm) == "s" {
				(*tasks)[idx].Completed = true
				(*tasks)[idx].UpdatedAt = time.Now()
				fmt.Println("✅ Tarea completada correctamente")
				time.Sleep(1 * time.Second)
				return
			}
		}
	}
	
}

func Handler(option int, tasks *[]Task) {

	switch option {
	case menu.AddTask:
		newTask, err := addTask(len(*tasks) + 1)
		if err == nil {
			*tasks = append(*tasks, newTask)
			fmt.Println("✅ Tarea agregada correctamente")
			time.Sleep(1 * time.Second)
		}
	case menu.ListTasks:
		listTasks(*tasks)
		fmt.Println(" ")
		fmt.Println("Presione enter para continuar")
		fmt.Scanln()
	case menu.CompleteTask:
		completeTask(tasks)
	}
}

