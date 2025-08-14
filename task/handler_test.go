package task

import (
	"io"
	"os"
	"strings"
	"task-manager/menu"
	"testing"
	"time"
)

var mockTask = Task{
	Id: 3,
	Title: "Task mock",
	Description: "Description mock",
	Completed: false,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func mockAddTask(t *testing.T) func(nextId int) (Task, error) {
	return func(nextId int) (Task, error) {
		t.Log("mockAddTask called")
		return mockTask, nil
	}
}

func mockListTasks(t *testing.T, called *bool) func([]Task) {
	return func(tasks []Task) {
		t.Log("mockListTasks called")
		*called = true
	}
}

func mockCompleteTask(t *testing.T, called *bool) func(*[]Task) {
	return func(tasks *[]Task) {
		t.Log("mockCompleteTask called")
		*called = true
	}
}


func mockTasks() []Task {
	return []Task{
		{
			Id: 1,
			Title: "Test Task 1",
			Description: "Test Description 1",
			Completed: false,
			CreatedAt: time.Now(),
		},
		{
			Id: 2,
			Title: "Test Task 2",
			Description: "Test Description 2",
			Completed: false,
			CreatedAt: time.Now(),
		},
	}
}

func TestHandler(t *testing.T) {
	testCases := []struct {
		name string
		option string
		initialTasks []Task
		expectedTasks []Task
		expectedOutput string
		deps Dependencies
		calledFlags map[string]*bool
	}{
		{
			name: "Should call addTask and add a task",
			option: menu.AddTask,
			initialTasks: mockTasks(),
			expectedTasks: append(mockTasks(), mockTask),
			expectedOutput: "✅ Tarea agregada correctamente",
			deps: Dependencies{
				AddTask: mockAddTask(t),
				ListTasks: nil,
				CompleteTask: nil,
			},
		},
		{
			name: "Should call listTasks and list tasks",
			option: menu.ListTasks,
			initialTasks: mockTasks(),
			expectedTasks: mockTasks(),
			expectedOutput: "Presione enter para continuar",
			deps: Dependencies{
				AddTask: nil,
				ListTasks: mockListTasks(t, new(bool)),
				CompleteTask: nil,
			},
		},{
			name: "Should call completeTask and complete a task",
			option: menu.CompleteTask,
			initialTasks: mockTasks(),
			expectedTasks: mockTasks(),
			expectedOutput: "",
			deps: Dependencies{
				AddTask: nil,
				ListTasks: nil,
				CompleteTask: mockCompleteTask(t, new(bool)),
			},
		},
	}

	for _, test :=range testCases {
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
			tasks := test.initialTasks
			deps := test.deps

			listTasksCalled := false
			completeTaskCalled := false
			deps.ListTasks = mockListTasks(t, &listTasksCalled)
			deps.CompleteTask = mockCompleteTask(t, &completeTaskCalled)

			Handler(test.option, &tasks, deps)
			w.Close()
			output, _ := io.ReadAll(r)
			outputString := strings.TrimSpace(string(output))
			if outputString != test.expectedOutput {
				t.Errorf("Expected '%s', got '%s'", test.expectedOutput, outputString)
			}
			switch test.option {
			case menu.AddTask:
				if len(tasks) != len(test.expectedTasks) {
					t.Errorf("Expected %d tasks, got %d", len(test.expectedTasks), len(tasks))
				}
			case menu.ListTasks:
				if !listTasksCalled {
					t.Errorf("Expected listTasks to be called, but it wasn't")
				}
			case menu.CompleteTask:
				if !completeTaskCalled {
					t.Errorf("Expected completeTask to be called, but it wasn't")
				}
			}
		})
	}
}