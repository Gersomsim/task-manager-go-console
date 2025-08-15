package task

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"task-manager/internal/cli"
	"testing"
	"time"
)

var mockInput = func(responses []string) cli.InputFunc {
	index := 0
	return func(prompt string) string {
		if index >= len(responses) {
			return ""
		}
		response := responses[index]
		index++
		return response
	}
}

func TestCompleteTask(t *testing.T) {
	completeLejend := "✅ Tarea completada correctamente"
	testCases := []struct {
		name string
		initialTasks []Task
		completeIndex int
		completedTask bool
		inputs []string
	}{
		{
			name: "Should complete the first task",
			initialTasks: mockTasks(),
			completeIndex: 0,
			completedTask: true,
			inputs: []string{
				"1",
				"s",
			},
		},
		{
			name: "Should not complete a task if the id is invalid",
			initialTasks: mockTasks(),
			completeIndex: 0,
			completedTask: false,
			inputs: []string{
				"invalid",
				"s",
			},
		},
		{
			name: "Should not complete a task if the confirmation is not 's'",
			initialTasks: mockTasks(),
			completeIndex: 0,
			completedTask: false,
			inputs: []string{
				"1",
				"n",
			},
		},
		{
			name: "Should not complete a task if the input is empty",
			initialTasks: mockTasks(),
			completeIndex: 0,
			completedTask: false,
			inputs: []string{
				"",
				"s",
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			originalOutput := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Error al crear el pipe: %v", err)
			}
			os.Stdout = w
			defer func() {
				os.Stdout = originalOutput
			}()
			mockInput := mockInput(test.inputs)
			CompleteTask(&test.initialTasks, mockInput)
			w.Close()
			output, _ := io.ReadAll(r)
			outputString := strings.TrimSpace(string(output))

			if test.completedTask {
				if !strings.Contains(outputString, completeLejend) {
					t.Errorf("Expected '%s', got '%s'", completeLejend, outputString)
				}
				if !test.initialTasks[test.completeIndex].Completed {
					t.Errorf("Expected task %d to be completed, got %v", test.completeIndex, test.initialTasks[test.completeIndex].Completed)
				}
			} else {
				if test.inputs[0] == "invalid" {
					invalidTest := "Error: ID inválido"
					if !strings.Contains(outputString, invalidTest) {
						t.Errorf("Expected '%s', got '%s'", invalidTest, outputString)
					}
					if test.initialTasks[test.completeIndex].Completed {
						t.Errorf("Expected task %d to be not completed, got %v", test.completeIndex, test.initialTasks[test.completeIndex].Completed)
					}
				}
			}
		})
	}
}

func TestListTasks(t *testing.T) {
	testCases := []struct {
		name string
		initialTasks []Task
		expectedOutput string
	}{
		{
			name: "Should list tasks",
			initialTasks: mockTasks(),
			expectedOutput: "1. Test Task 1\n2. Test Task 2\n3. Test Task 3",
		},
		{
			name: "Should list tasks with empty tasks",
			initialTasks: []Task{},
			expectedOutput: "No hay tareas",
		},
		{
			name: "Should list tasks with one task",
			initialTasks: []Task{
				mockTask,
			},
			expectedOutput: "1. Test Task 1",
		},
	}
	for _, test := range testCases{
		t.Run(test.name, func(t *testing.T){
			originalOutput := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Error al crear el pipe: %v", err)
			}
			os.Stdout = w
			defer func() {
				os.Stdout = originalOutput
			}()
			ListTasks(test.initialTasks)
			w.Close()
			output, _ := io.ReadAll(r)
			outputString := strings.TrimSpace(string(output))
			title := "Título"
			if !strings.Contains(outputString, title) {
				t.Errorf("Expected '%s', got '%s'", title, outputString)
			}
			if strings.Contains(outputString, test.expectedOutput) {
				t.Errorf("Expected '%s', got '%s'", test.expectedOutput, outputString)
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	testCases := []struct {
		name string
		id int
		expectedTask Task
		expectedError error
	}{
		{
			name: "Should add a task",
			id: 1,
			expectedTask: Task{
				Id: 1,
				Title: "Test Task 1",
				Description: "Test Task 1 description",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name: "Should not add a task if the title is empty",
			id: 1,
			expectedTask: Task{
				Id: 1,
				Title: "",
				Description: "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: errors.New("título no válido"),
		},
		
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			originalOutput := os.Stdout
			r, w, err :=os.Pipe()
			if err != nil {
				t.Fatalf("Error al crear el pipe: %v", err)
			}
			os.Stdout = w
			defer func() {
				os.Stdout = originalOutput
			}()

			responses := []string{
				test.expectedTask.Title,
				test.expectedTask.Description,
			}
			mockInput := mockInput(responses)

			task, err := AddTask(test.id, mockInput)
			w.Close()
			output, _ := io.ReadAll(r)
			outputString := strings.TrimSpace(string(output))
			message := fmt.Sprintf("Agregar tarea %d o precione enter para salir", test.id)

			if !strings.Contains(outputString, message) {
				t.Errorf("Expected output to contain '%s', got '%s'", message, outputString)
			}
			
			if test.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error: %v, got nil", test.expectedError)
				} else if err.Error() != test.expectedError.Error() {
					t.Errorf("Expected error: %v, got %v", test.expectedError, err)
				}
			} else if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if task.Title != test.expectedTask.Title {
				t.Errorf("Expected title %s, got %s", test.expectedTask.Title, task.Title)
			}
			if task.Description != test.expectedTask.Description {
				t.Errorf("Expected description %s, got %s", test.expectedTask.Description, task.Description)
			}
		})
	}
}