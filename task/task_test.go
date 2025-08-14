package task

import (
	"testing"
	"time"
)

func TestMarkAsCompleted(t *testing.T) {
	task := Task{
		Id: 1,
		Title: "Test Task",
		Description: "Test Description",
		Completed: false,
		CreatedAt: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	originalUpdatedAt := task.UpdatedAt
	originalCreatedAt := task.CreatedAt

	task.MarkAsCompleted()

	if !task.Completed {
		t.Error("La tarea no se marcó como completada")
	}
	if !task.UpdatedAt.After(originalUpdatedAt) {
		t.Error("La fecha de actualización no se actualizó")
	}
	if !task.CreatedAt.Equal(originalCreatedAt) {
		t.Error("La fecha de creación no se mantuvo")
	}

}