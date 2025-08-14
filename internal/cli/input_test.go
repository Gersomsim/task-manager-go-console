package cli

import (
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	// 1 entrada simulada
	entradaSimulada := "Hola Go!\n"
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error al crear el pipe: %v", err)
	}
	// guardar el StdIn original
	stdinOriginal := os.Stdin
	defer func() {
		os.Stdin = stdinOriginal
	}()

	os.Stdin = r
	// escribir la entrada simulada
	go func() {
		defer w.Close()
		w.WriteString(entradaSimulada)
	}()
	// ejecutar la funcion a testear
	message := "Ingrese un texto: "
	resultado := Input(message)

	// verificar el resultado
	esperado := "Hola Go!"
	if resultado != esperado {
		t.Errorf("Input('%s') = '%s', se esperaba: %s", message, resultado, esperado)
	}	
}

// Table Driven Test
func TestInputTable(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
	}{
		{name: "Entrada simple", input: "Hola\n", expected: "Hola"},
		{name: "Entrada con espacios", input: "  Hola Go  \n", expected: "Hola Go"},
		{name: "Entrada con caracteres especiales", input: " correo@mail.com \n", expected: "correo@mail.com"},
		{name: "Entrada con numeros", input: "1234567890 \n", expected: "1234567890"},
		{name: "Entrada con caracteres alfanumericos", input: "Hola 123 \n", expected: "Hola 123"},
		{name: "Entrada vacia", input: "\n", expected: ""},
		{name: "Entrada con solo espacios", input: "       \n", expected: ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			// simular entrada
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Error al crear el pipe: %v", err)
			}
			// guardar el StdIn original
			stdinOriginal := os.Stdin
			defer func() {
				os.Stdin = stdinOriginal
			}()
			os.Stdin = r
			// escribir la entrada simulada
			go func() {
				defer w.Close()
				w.WriteString(test.input)
			}()
			// ejecutar la funcion a testear
			resultado := Input("Ingrese un texto: ")
			// verificar el resultado
			if resultado != test.expected {
				t.Errorf("Input('%s') = '%s', se esperaba: %s", test.input, resultado, test.expected)
			}
		})
	}
}