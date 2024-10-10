package api

import (
	"encoding/json"
	"os"
	"task-cli/model"
)

var filePath = "tasks.json"

func GetTasks() (model.Tasks, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return model.Tasks{Tasks: make([]model.Task, 0)}, nil
	}

	var tasks model.Tasks

	err = json.Unmarshal(fileContent, &tasks)

	if err != nil {
		return model.Tasks{}, err
	}

	return tasks, nil
}

func WriteTasks(tasks model.Tasks) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)

	if err != nil {
		return err
	}

	return nil
}
