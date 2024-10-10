package commands

import (
	"bytes"
	"task-cli/api"
	"task-cli/model"
	"time"
)

func AddTask(description string) (int, error) {
	tasks, err := api.GetTasks()

	if err != nil {
		return -1, err
	}

	idToAllocate := 0
	length := len(tasks.Tasks)
	if length != 0 {
		idToAllocate = tasks.Tasks[length-1].Id + 1
	}

	currentTime := time.Now()
	newTask := model.Task{
		Id:          idToAllocate,
		Description: description,
		Status:      model.TODO,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	tasks.Tasks = append(tasks.Tasks, newTask)

	err = api.WriteTasks(tasks)

	if err != nil {
		return -1, err
	}

	return newTask.Id, nil
}

func UpdateTask(id int, description string) (int, error) {
	tasks, err := api.GetTasks()

	if err != nil {
		return -1, err
	}

	currentTime := time.Now()

	var taskToUpdate *model.Task

	found := false

	for i := 0; i < len(tasks.Tasks); i++ {
		if tasks.Tasks[i].Id == id {
			taskToUpdate = &tasks.Tasks[i]
			found = true
			break
		}
	}
	if !found {
		return -1, nil
	}
	taskToUpdate.Description = description
	taskToUpdate.UpdatedAt = currentTime

	err = api.WriteTasks(tasks)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func DeleteTask(id int) (int, error) {
	tasks, err := api.GetTasks()

	if err != nil {
		return -1, err
	}

	for i := 0; i < len(tasks.Tasks); i++ {
		if tasks.Tasks[i].Id == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)

			err = api.WriteTasks(tasks)

			if err != nil {
				return -1, err
			}

			return id, nil
		}
	}

	return -1, nil
}

func changeStatus(id int, status model.Status) (int, error) {
	tasks, err := api.GetTasks()

	if err != nil {
		return -1, err
	}

	for i := 0; i < len(tasks.Tasks); i++ {
		if tasks.Tasks[i].Id == id {
			tasks.Tasks[i].Status = status

			err = api.WriteTasks(tasks)

			if err != nil {
				return -1, err
			}

			return id, nil
		}
	}

	return -1, nil
}

func MarkTaskAsDone(id int) (int, error) {
	return changeStatus(id, model.DONE)
}

func MarkTaskAsInProgress(id int) (int, error) {
	return changeStatus(id, model.InProgress)
}

func listTasks(criteria ...model.Status) (string, error) {
	tasks, err := api.GetTasks()

	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	if len(criteria) == 0 {
		for i := 0; i < len(tasks.Tasks); i++ {
			buffer.WriteString(model.TaskFormatter(tasks.Tasks[i]))
		}
	} else {
		for i := 0; i < len(tasks.Tasks); i++ {
			if tasks.Tasks[i].Status == criteria[0] {
				buffer.WriteString(model.TaskFormatter(tasks.Tasks[i]))
			}
		}
	}

	return buffer.String(), nil
}

func ListAllTasks() (string, error) {
	return listTasks()
}

func ListAllDoneTasks() (string, error) {
	return listTasks(model.DONE)
}

func ListAllInProgressTasks() (string, error) {
	return listTasks(model.InProgress)
}

func ListAllTODOTasks() (string, error) {
	return listTasks(model.TODO)
}
