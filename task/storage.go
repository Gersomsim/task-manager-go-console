package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type FileCreator func(name string) (io.WriteCloser, error)

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



func SaveToFile(tasks []Task, filename string, fileCreator FileCreator) error {
	fmt.Print("\033[H\033[2J")
	makeDir()
	fmt.Println("💾 Guardando tareas en el archivo...")

	// creamos el archivo o lo sobreescribimos
	file, err := fileCreator(dir + "/" + filename)
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

type FileOpener func(name string) (io.ReadCloser, error)

func LoadFromFile(filename string, fileOpener FileOpener) ([]Task, error) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Cargando tareas desde el archivo")
	var tasks []Task
	file, err := fileOpener(dir + "/" + filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return []Task{}, errors.New("error al abrir el archivo")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return tasks, errors.New("error al decodificar el archivo")
	}
	return tasks, nil
}