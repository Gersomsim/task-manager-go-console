package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
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

func mockFileOpener(content string, err error) func(name string) (io.ReadCloser, error) {
	return func(name string) (io.ReadCloser, error) {
			if err != nil {
					return nil, err
			}
			return &mockFile{Buffer: bytes.NewBufferString(content)}, nil
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

func TestLoadFromFile(t *testing.T) {
	tests := []struct {
	name              string
	fileContent       string
	openFileErr       error
	expectedTasksCount int
	expectedError     error
}{
	{
			name:              "Debería cargar las tareas correctamente",
			fileContent:       `[{"Id": 1, "Title": "Tarea 1"}]`,
			openFileErr:       nil,
			expectedTasksCount: 1,
			expectedError:     nil,
	},
	{
			name:              "Debería retornar un slice vacío si el archivo no existe",
			fileContent:       "",
			openFileErr:       os.ErrNotExist,
			expectedTasksCount: 0,
			expectedError:     nil,
	},
	{
			name:              "Debería retornar un error si el JSON es inválido",
			fileContent:       `{"Id": 1`, // JSON inválido
			openFileErr:       nil,
			expectedTasksCount: 0,
			expectedError:     errors.New("error al decodificar el archivo"),
	},
	{
			name:              "Debería retornar un error genérico",
			fileContent:       "",
			openFileErr:       errors.New("error de permisos"),
			expectedTasksCount: 0,
			expectedError:     errors.New("error al abrir el archivo"),
	},
}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Creacion de mock de apertura de archivo
			mockOpener := mockFileOpener(test.fileContent, test.openFileErr)

			// Llamada a la funcion a testear
			tasks, err := LoadFromFile("tasks.json", mockOpener)

			// Verificacion del error esperado 
			if err != nil && err.Error() != test.expectedError.Error() {
				t.Errorf("Expected error: %v, got: %v", test.expectedError, err)
			}

			// Verificacion del numero de tareas esperado
			if len(tasks) != test.expectedTasksCount {
				t.Errorf("Expected %d tasks, got %d", test.expectedTasksCount, len(tasks))
			}
			
		})
	}
}

func mockMkdirAlreadyExists(name string, perm os.FileMode) error {
	return os.ErrExist // Retorna el error específico de "ya existe".
}

// mockMkdirWithError simula un error de creación diferente.
func mockMkdirWithError(name string, perm os.FileMode) error {
	return errors.New("simulated permission denied") // Retorna un error genérico.
}

func TestMakeDir(t *testing.T) {
	// Capturamos la salida para verificar los mensajes impresos.
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	t.Run("No debería retornar un error si el directorio ya existe", func(t *testing.T) {
		// Llamamos a la función con el mock para el caso de "ya existe".
		makeDir(mockMkdirAlreadyExists)

		// Cerramos el pipe de escritura y leemos la salida.
		w.Close()
		output, _ := ioutil.ReadAll(r)
		outputStr := strings.TrimSpace(string(output))

		// Verificamos que no se imprimió nada, ya que el error se ignoró.
		if outputStr != "" {
			t.Errorf("Se imprimió una salida inesperada: %s", outputStr)
		}
	})

	t.Run("Debería imprimir un error si la creación falla", func(t *testing.T) {
		// Capturamos la salida de nuevo para este caso de prueba.
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		defer func() {
			os.Stdout = oldStdout
		}()

		// Llamamos a la función con el mock para el caso de error.
		makeDir(mockMkdirWithError)

		w.Close()
		output, _ := ioutil.ReadAll(r)
		outputStr := strings.TrimSpace(string(output))

		// Verificamos que el mensaje de error se imprimió.
		expectedOutput := "Error al crear el directorio: simulated permission denied"
		if !strings.Contains(outputStr, expectedOutput) {
			t.Errorf("Mensaje de error inesperado. Se esperaba que contuviera '%s', se obtuvo '%s'", expectedOutput, outputStr)
		}
	})
}