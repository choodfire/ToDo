package main

import (
	"flag"
	"simpleCLI/data"
)

func main() {
	t := data.Tasks{}

	err := t.GetFromFile()
	if err != nil {
		panic(err)
	}

	list := flag.Bool("list", false, "display list of tasks") // usage = help message
	add := flag.String("add", "", "add a task")
	complete := flag.Int("complete", 0, "mark task as completed")
	unComplete := flag.Int("unсomplete", 0, "mark task as completed")
	deleteTask := flag.Int("delete", 0, "mark task as completed")
	flag.Parse()

	if *list == true {
		t.Output()
	}

	if len(*add) > 0 {
		t.Add(*add)
	}

	if *complete != 0 {
		t.MarkDone(*complete)
	}

	if *unComplete != 0 {
		t.MarkUndone(*unComplete)
	}

	if *deleteTask != 0 {
		t.Delete(*deleteTask)
	}

	t.Output()

}

// todo add colors in output
