package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"
	"time"
)

type mockFile struct {
	*bytes.Buffer
}

func (m *mockFile) Close() error {
	return nil
}

func mockFileCreator(t *testing.T, buf *bytes.Buffer) func(name string) (io.WriteCloser, error) {
	return func(name string) (io.WriteCloser, error) {
		t.Logf("Se intentó crear el archivo: %s", name)
		return &mockFile{Buffer: buf}, nil
	}
}

func TestSaveToFile(t *testing.T) {
	t.Run("Debería guardar las tareas en formato JSON correctamente", func(t *testing.T) {
		// Preparamos los datos de prueba.
		tasks := []Task{
			{
				Id: 1,
				Title: "Test Task 1",
				Description: "Description 1",
				Completed: false,
				CreatedAt: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
		}

		// Creamos un buffer para simular el archivo.
		var buf bytes.Buffer
		
		// Llamamos a la función con nuestro mock de creación de archivo.
		err := SaveToFile(tasks, "tasks.json", mockFileCreator(t, &buf))

		// Verificamos que no haya errores.
		if err != nil {
			t.Fatalf("Se obtuvo un error inesperado: %v", err)
		}

		// Verificamos el contenido del buffer.
		var savedTasks []Task
		err = json.Unmarshal(buf.Bytes(), &savedTasks)
		if err != nil {
			t.Fatalf("No se pudo decodificar el JSON del buffer: %v", err)
		}
        
        // Comprobamos que el número de tareas sea el correcto.
		if len(savedTasks) != 1 {
			t.Errorf("Número de tareas incorrecto. Se esperaba 1, se obtuvo %d", len(savedTasks))
		}
        
        // Comprobamos que los datos se guardaron correctamente.
		if savedTasks[0].Title != "Test Task 1" {
			t.Errorf("El título no coincide. Se esperaba 'Test Task 1', se obtuvo '%s'", savedTasks[0].Title)
		}
	})

	t.Run("Debería retornar un error si no se puede crear el archivo", func(t *testing.T) {
		// Mock que simula un error de creación de archivo.
		errorCreator := func(name string) (io.WriteCloser, error) {
			return nil, errors.New("error de permisos simulado")
		}

		var tasks []Task
		err := SaveToFile(tasks, "tasks.json", errorCreator)

		if err == nil {
			t.Error("Se esperaba un error, pero se obtuvo nil.")
		}
		if !strings.Contains(err.Error(), "error de permisos simulado") {
			t.Errorf("Mensaje de error inesperado. Se esperaba un error de permisos, se obtuvo: %v", err)
		}
	})
}