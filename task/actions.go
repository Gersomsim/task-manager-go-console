package task

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-manager/internal/cli"
	"text/tabwriter"
	"time"
)


func AddTask(nextId int, input cli.InputFunc) (Task, error) {
	var task Task
	task.Id = nextId
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Agregar tarea %d o precione enter para salir \n", nextId)
	fmt.Println("--------------------------------")
	task.Title = input("Ingrese el titulo de la tarea: ")
	if task.Title == "" {
		return task, errors.New("título no válido")
	}
	task.Description = input("Ingrese la descripcion de la tarea: ")

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return task, nil
}

func ListTasks(tasks []Task) {
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

func CompleteTask(tasks *[]Task, input cli.InputFunc) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Completar tarea")
	fmt.Println("--------------------------------")
	id := input("Ingrese el id de la tarea a completar: ")
	if id == "" {
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: ID inválido")
		time.Sleep(1 * time.Second)
		return
	}
	
	for idx, task := range *tasks {
		if task.Id == idInt {
			fmt.Printf("¿Está seguro de completar la tarea %s?\n", task.Title)
			confirm := input("s/n: ")
			if strings.ToLower(confirm) == "s" {
				(*tasks)[idx].MarkAsCompleted()
				fmt.Println("✅ Tarea completada correctamente")
				time.Sleep(1 * time.Second)
				return
			}
		}
	}
	
}