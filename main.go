package main

import "simpleCLI/data"

func main() {
	t := data.Tasks{}

	err := t.GetFromJSON()
	if err != nil {
		panic(err)
	}

	//t.Add("First")
	//t.Add("Second")
	//t.Add("package main")
	//t.Add("Новобергеровское братство")
	//t.Add("big guy started a call. — Today at 9:50 PM")
	//
	//t.Output()
}
