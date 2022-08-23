package data

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	_ "github.com/alexeyco/simpletable"
	"time"
)

type task struct {
	title         string
	isDone        bool
	timeCreated   time.Time
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

func (t *task) getIsDone() string {
	if t.isDone == true {
		return "yes"
	}
	return "no"
}

func (t *task) getCreatedTime() string {
	//t.timeCreated
	return fmt.Sprintf("%.2d.%.2d.%.2d %.2d:%.2d:%.2d", t.timeCreated.Day(),
		t.timeCreated.Month(),
		t.timeCreated.Year(),
		t.timeCreated.Hour(),
		t.timeCreated.Minute(),
		t.timeCreated.Second(),
	)
}

func (t *task) getCompletedTime() string {
	return fmt.Sprintf("%.2d.%.2d.%.2d %.2d:%.2d:%.2d", t.timeCompleted.Day(),
		t.timeCompleted.Month(),
		t.timeCompleted.Year(),
		t.timeCompleted.Hour(),
		t.timeCompleted.Minute(),
		t.timeCompleted.Second(),
	)
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

func (t *Tasks) GetCount() int {
	return len(t.Tasks)
}

func (t *Tasks) Output() {
	//StyleDefault = &Style{
	//	Border: &BorderStyle{
	//		TopLeft:            "+",
	//		Top:                "-",
	//		TopRight:           "+",
	//		Right:              "|",
	//		BottomRight:        "+",
	//		Bottom:             "-",
	//		BottomLeft:         "+",
	//		Left:               "|",
	//		TopIntersection:    "+",
	//		BottomIntersection: "+",
	//	},
	//	Divider: &DividerStyle{
	//		Left:         "+",
	//		Center:       "-",
	//		Right:        "+",
	//		Intersection: "+",
	//	},
	//	Cell: "|",
	//}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created at"},
			{Align: simpletable.AlignCenter, Text: "Completed at"},
		},
	}

	i := 1
	for _, task := range t.Tasks {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i)},
			{Text: task.getTitle()},
			{Align: simpletable.AlignCenter, Text: task.getIsDone()},
			{Align: simpletable.AlignLeft, Text: task.getCreatedTime()},
			{Align: simpletable.AlignLeft, Text: task.getCompletedTime()},
		}

		table.Body.Cells = append(table.Body.Cells, r)

		i += 1
	}

	fmt.Println(table.String())
}
