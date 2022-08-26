package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"simpleCLI/data"
	"syscall"
)

func displayHelpMessage() {
	fmt.Println("	-list				Display all tasks")
	fmt.Println("	-add \"Taskname\"			Add new task")
	fmt.Println("	-c \"Task number\"		Mark task as completed")
	fmt.Println("	-uc \"Task number\"		Mark task as uncompleted")
	fmt.Println("	-delete \"Task number\"		Delete task")
	fmt.Println("	-help				Display help message")
}

func main() {
	t := data.Tasks{}

	err := t.GetFromFile()
	if err != nil {
		if errors.Is(err, syscall.ERROR_FILE_NOT_FOUND) {
			fmt.Println("TODO list is empty, add a new one using -add \"Taskname\"")
			os.Exit(1)
		}
		//panic(err)
	}

	list := flag.Bool("list", false, "display list of tasks") // usage = help message
	add := flag.String("add", "", "add a task")
	complete := flag.Int("c", 0, "mark task as completed")
	unComplete := flag.Int("uc", 0, "mark task as uncompleted")
	deleteTask := flag.Int("delete", 0, "delete task")
	help := flag.Bool("help", false, "display help message")
	flag.Parse()

	if *list == true {
		t.Output()
	}

	if len(*add) > 0 {
		t.Add(*add)
		fmt.Printf("Task \"%s\" added to list\n", *add)
		t.Output()
	}

	if *complete != 0 {
		err := t.MarkDone(*complete)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("#%d marked as completed\n", *complete)
		t.Output()
	}

	if *unComplete != 0 {
		err := t.MarkUndone(*unComplete)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("#%d marked as uncompleted\n", *unComplete)
		t.Output()
	}

	if *deleteTask != 0 {
		err := t.Delete(*deleteTask)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("#%d deleted\n", *deleteTask)
		t.Output()
	}

	if *help == true {
		displayHelpMessage()
	}

}

// todo change default message
