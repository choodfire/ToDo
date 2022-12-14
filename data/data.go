package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"io"
	"os"
	"time"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	//colorYellow = "\033[33m"
	//colorBlue   = "\033[34m"
	//colorPurple = "\033[35m"
	//colorCyan   = "\033[36m"
	colorGray = "\033[37m"
	//colorWhite  = "\033[97m"
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
	if t.IsDone == true {
		return fmt.Sprintf(string(colorGreen) + t.Title + string(colorReset))
	}
	return fmt.Sprintf(string(colorGray) + t.Title + string(colorReset))
}

func (t *task) getIsDone() string {
	if t.IsDone == true {
		return fmt.Sprintf(string(colorGreen) + "yes" + string(colorReset))
	}
	return fmt.Sprintf(string(colorRed) + "no" + string(colorReset))
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
	if t.TimeCompleted.Equal(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return fmt.Sprintf(string(colorRed) + "not completed yet" + string(colorReset))
	}
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

func saveToFile(t Tasks) {
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

	err = os.WriteFile("./data/data.json", file, 0644)
	if err != nil {
		panic(err)
	}
}

func (t *Tasks) GetFromFile() error {
	jsonFile, err := os.Open("data/data.json")
	if err != nil {
		return err
	}

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tasks) Add(title string) {
	t.Tasks = append(t.Tasks, *NewTask(title))

	saveToFile(*t)
}

func (t *Tasks) Delete(index int) error {
	index = index - 1
	if index > -1 && index < t.getLength() {
		t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
		saveToFile(*t)

		return nil
	} else {
		return errors.New("Wrong index!")
	}

}

func (t *Tasks) MarkDone(index int) error {
	index = index - 1
	if index > -1 && index < t.getLength() {
		t.Tasks[index].markCompleted()
		saveToFile(*t)

		return nil
	} else {
		return errors.New("Wrong index!")
	}
}

func (t *Tasks) MarkUndone(index int) error {
	index = index - 1
	if index > -1 && index < t.getLength() {
		t.Tasks[index].markUnCompleted()
		saveToFile(*t)

		return nil
	} else {
		return errors.New("Wrong index")
	}
}

func (t *Tasks) getLength() int {
	return len(t.Tasks)
}

func (t *Tasks) Output() {
	if len(t.Tasks) == 0 {
		return
	}
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
			{Align: simpletable.AlignLeft, Text: task.getTitle()},
			{Align: simpletable.AlignLeft, Text: task.getIsDone()},
			{Align: simpletable.AlignLeft, Text: task.getCreatedTime()},
			{Align: simpletable.AlignLeft, Text: task.getCompletedTime()},
		}

		table.Body.Cells = append(table.Body.Cells, r)

		i += 1
	}

	fmt.Println(table.String())
}
