package menu

import (
	"io"
	"os"
	"strings"
	"task-manager/internal/cli"
	"testing"
)

func mockInput(value string) cli.InputFunc {
	return func(prompt string) string {
		return value
	}
}

func TestShowMenu(t *testing.T) {
	// Simular la salida  (os.Stdout) para capturar el mensaje
	originalStdout := os.Stdout
	_, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error al crear el pipe: %v", err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	// Ejecutar la funcion a testear
	option, err := ShowMenu(mockInput("1"))
	if option != AddTask {
		t.Errorf("ShowMenu() = %s, se esperaba: %s", option, AddTask)
	}
	if err != nil {
		t.Errorf("ShowMenu() = %v, se esperaba: nil", err)
	}
}

func TestShowMenuInvalidOption(t *testing.T) {
	expectedError := "la opcion ingresada no es valida, debe ser un numero del 1 al 4"
	invalidOptions := []struct {
		name string
		input string
	}{
		{name: "Entrada vacia", input: "", },
		{name: "Entrada con solo espacios", input: "   ", },
		{name: "Entrada con caracteres especiales", input: "!@#$%^&*()", },
		{name: "Entrada con numeros", input: "1234567890", },
		{name: "Entrada con letras", input: "abc", },
		{name: "Entrada con letras y numeros", input: "abc123", },
		{name: "Entrada con letras y caracteres especiales", input: "abc!@#", },
		{name: "Entrada con numeros y caracteres especiales", input: "123!@#", },
		{name: "Entrada con letras, numeros y caracteres especiales", input: "abc123!@#", },
	}

	for _, test := range invalidOptions {
		t.Run(test.name, func(t *testing.T) {
			originalStdout := os.Stdout
			_, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Error al crear el pipe: %v", err)
			}
			os.Stdout = w
			defer func() {
				os.Stdout = originalStdout
			}()

			_, err = ShowMenu(mockInput(test.input))
			if err == nil {
				t.Errorf("ShowMenu() = nil, se esperaba: %s", expectedError)
			}
			if err.Error() != expectedError {
				t.Errorf("ShowMenu() = %v, se esperaba: %s", err, expectedError)
			}
		})
	}
}

func TestShowMenuValidOption(t *testing.T) {
	validOptions := []struct {
		name string
		input string
		expected string
	}{
		{name: "Entrada con numero 1", input: "1", expected: AddTask},
		{name: "Entrada con numero 2", input: "2", expected: ListTasks},
		{name: "Entrada con numero 3", input: "3", expected: CompleteTask},
		{name: "Entrada con numero 4", input: "4", expected: Exit},
	}

	for _, test := range validOptions {
		t.Run(test.name, func(t *testing.T) {
			originalStdout := os.Stdout
			_, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Error al crear el pipe: %v", err)
			}
			os.Stdout = w
			defer func() {
				os.Stdout = originalStdout
			}()

			option, err := ShowMenu(mockInput(test.input))
			if option != test.expected {
				t.Errorf("ShowMenu() = %s, se esperaba: %s", option, test.expected)
			}
			if err != nil {
				t.Errorf("ShowMenu() = %v, se esperaba: nil", err)
			}
		})
	}
}

func TestShowMenuOutput(t *testing.T) {
	// Capturar la salida del os.Stdout
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error al crear el pipe: %v", err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	// Ejecutar la funcion a testear
	ShowMenu(mockInput("1"))

	// Leer la salida capturada
	w.Close()
	salida, _ := io.ReadAll(r)
	salidaStr := string(salida)
	
	if !strings.Contains(salidaStr, "Seleccione una opción:") {
		t.Error("El menú no contiene el encabezado")
	}
	if !strings.Contains(salidaStr, "1. 💾 Agregar tarea") {
		t.Error("El menú no contiene la opción de agregar tarea")
	}
	if !strings.Contains(salidaStr, "2. 📝 Listar tareas") {
		t.Error("El menú no contiene la opción de listar tareas")
	}
	if !strings.Contains(salidaStr, "3. ✅ Marcar tarea como completada") {
		t.Error("El menú no contiene la opción de marcar tarea como completada")
	}
	if !strings.Contains(salidaStr, "4. 🚪 Salir") {
		t.Error("El menú no contiene la opción de salir")
	}
}