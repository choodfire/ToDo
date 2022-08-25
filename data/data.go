package data

import (
	"encoding/json"
	"fmt"
	"github.com/alexeyco/simpletable"
	_ "github.com/alexeyco/simpletable"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type task struct {
	Title         string    `json:"title"`
	IsDone        bool      `json:"isDone"`
	TimeCreated   time.Time `json:"timeCreated"`
	TimeCompleted time.Time `json:"timeCompleted"`
}

func NewTask(title string) *task {
	return &task{title,
		false,
		time.Now(),
		time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func (t *task) getTitle() string {
	return t.Title
}

func (t *task) getIsDone() string {
	if t.IsDone == true {
		return "yes"
	}
	return "no"
}

func (t *task) getCreatedTime() string {
	return fmt.Sprintf("%.2d.%.2d.%.2d %.2d:%.2d:%.2d", t.TimeCreated.Day(),
		t.TimeCreated.Month(),
		t.TimeCreated.Year(),
		t.TimeCreated.Hour(),
		t.TimeCreated.Minute(),
		t.TimeCreated.Second(),
	)
}

func (t *task) getCompletedTime() string {
	return fmt.Sprintf("%.2d.%.2d.%.2d %.2d:%.2d:%.2d", t.TimeCompleted.Day(),
		t.TimeCompleted.Month(),
		t.TimeCompleted.Year(),
		t.TimeCompleted.Hour(),
		t.TimeCompleted.Minute(),
		t.TimeCompleted.Second(),
	)
}

func (t *task) setTitle(newTitle string) {
	t.Title = newTitle
}

func (t *task) markCompleted() {
	t.IsDone = true
	t.TimeCompleted = time.Now()
}

func (t *task) markUnCompleted() {
	t.IsDone = false
	t.TimeCompleted = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
}

///

type Tasks struct {
	Tasks []task `json:"tasks"`
}

func saveToJSON(t Tasks) {
	if _, err := os.Stat("./data/data.json"); !os.IsNotExist(err) {
		err := os.Remove("./data/data.json")
		if err != nil {
			panic(err)
		}
	}

	file, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./data/data.json", file, 0644)
	if err != nil {
		panic(err)
	}
}

func (t *Tasks) GetFromJSON() error {
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	if json.Valid(data) == false {
		log.Fatal("JSON file isn't valid")
	}

	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tasks) Add(title string) {
	t.Tasks = append(t.Tasks, *NewTask(title))

	saveToJSON(*t)
}

func (t *Tasks) Delete(index int) { // todo check index
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
}

func (t *Tasks) MarkDone(index int) { // todo check index
	t.Tasks[index].markCompleted()
}

func (t *Tasks) MarkUndone(index int) { // todo check index
	t.Tasks[index].markUnCompleted()
}

func (t *Tasks) getCount() int {
	return len(t.Tasks)
}

func (t *Tasks) Output() {
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
