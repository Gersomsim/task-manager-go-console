package task

import "time"

type Task struct {
	Id int
	Title string
	Description string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}
