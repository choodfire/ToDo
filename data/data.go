package data

import "time"

type task struct {
	title         string
	isDone        bool
	timeAdded     time.Time
	timeCompleted time.Time
}

func NewTask(title string) *task {
	return &task{title,
		false,
		time.Now(),
		time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func (t *task) getTitle() string {
	return t.title
}

func (t *task) getIsDone() bool {
	return t.isDone
}

func (t *task) getAddedTime() time.Time {
	return t.timeAdded
}

func (t *task) getCompletedTime() time.Time {
	return t.timeCompleted
}

func (t *task) setTitle(newTitle string) {
	t.title = newTitle
}

func (t *task) markCompleted() {
	t.isDone = true
	t.timeCompleted = time.Now()
}

func (t *task) markUnCompleted() {
	t.isDone = false
	t.timeCompleted = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
}

///

type Tasks struct {
	Tasks []task
}

func (t *Tasks) Add(title string) {
	t.Tasks = append(t.Tasks, *NewTask(title))
}

func (t *Tasks) Delete(index int) {
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
}

func (t *Tasks) MarkDone(index int) {
	t.Tasks[index].markCompleted()
}

func (t *Tasks) MarkUndone(index int) {
	t.Tasks[index].markUnCompleted()
}
