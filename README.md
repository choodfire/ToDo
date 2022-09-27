# ToDo CLI 

Simple ToDo command line app written in GoLang

---

### Usage

```bash
./ToDo [options] [value]
```

Options:

*  `-list` Display all tasks
*  `-add` (taskName) Add task *taskName*
*  `-c` (index) Mark task #index as completed
*  `-uc` (index) Mark task #index as uncompleted
*  `-delete` (index) Delete task #index
*  `-help` Display help message

---

### Build executable

Clone repository with `git clone https://github.com/choodfire/ToDo` or download it [here](https://github.com/choodfire/ToDo/releases)

Then move into directory with `cd ToDo`

Run `go build` (Make sure you have Go installed, you can do it [here](https://go.dev/dl/))

After that you can use the app