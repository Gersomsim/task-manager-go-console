package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const dir = "storage"

func makeDir() {
	err := os.Mkdir(dir, 0755)
	if os.IsExist(err) {
		return
	}
	if err != nil {
		fmt.Printf("Error al crear el directorio: %v\n", err)
		time.Sleep(1 * time.Second)
		return
	}
}

func SaveToFile(tasks []Task, filename string ) error {
	fmt.Print("\033[H\033[2J")
	makeDir()
	fmt.Println("💾 Guardando tareas en el archivo...")

	// creamos el archivo o lo sobreescribimos
	file, err := os.Create(dir + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}

func LoadFromFile(filename string) ([]Task, error) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Cargando tareas desde el archivo")
	var tasks []Task
	file, err := os.Open(dir + "/" + filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return []Task{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}