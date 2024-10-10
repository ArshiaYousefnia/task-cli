package main

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/commands"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
func root(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: task-cli <command> [<args>]")
		return
	}
	command := args[0]

	if command == "add" {
		if len(args) < 2 {
			fmt.Println("task description not provided")
			return
		}
		description := args[1]

		_, err := commands.AddTask(description)

		if err != nil {
			fmt.Println("could not add task")
			return
		}

		fmt.Println("task added successfully")
	} else if command == "update" {
		if len(args) < 3 {
			fmt.Println("new task description missing")
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("could not parse id, update failed")
			return
		}
		description := args[2]

		i, err := commands.UpdateTask(id, description)

		if err != nil {
			fmt.Println("could not update task")
			return
		}

		fmt.Printf("task %d updated successfully\n", i)
	} else if command == "delete" {
		if len(args) < 2 {
			fmt.Println("task id to delete not provided")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("could not parse id, delete failed")
			return
		}

		i, err := commands.DeleteTask(id)

		if err != nil {
			fmt.Println("could not delete task")
			return
		}

		fmt.Printf("task %d deleted successfully\n", i)
	} else if command == "mark-in-progress" {
		if len(args) < 2 {
			fmt.Println("task id not provided")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("could not parse id, mark-in-progress failed")
			return
		}

		i, err := commands.MarkTaskAsInProgress(id)

		if err != nil {
			fmt.Println("could not mark task as in progress")
			return
		}

		fmt.Printf("task %d marked as in progress", i)
	} else if command == "mark-done" {
		if len(args) < 2 {
			fmt.Println("task id not provided")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("could not parse id, mark done failed")
			return
		}

		i, err := commands.MarkTaskAsDone(id)

		if err != nil {
			fmt.Println("could not mark task as done")
			return
		}

		fmt.Printf("task %d marked as done successfully", i)
	} else if command == "list" {
		if len(args) < 2 {
			fmt.Print(commands.ListAllTasks())
			return
		}
		criteria := args[1]

		if criteria == "done" {
			fmt.Print(commands.ListAllDoneTasks())
		} else if criteria == "todo" {
			fmt.Print(commands.ListAllTODOTasks())
		} else if criteria == "in-progress" {
			fmt.Print(commands.ListAllInProgressTasks())
		}
	} else {
		fmt.Println("invalid command")
	}
}

func main() {
	root(os.Args[1:])
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
