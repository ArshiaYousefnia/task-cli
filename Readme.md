## task-cli is a cli utility to manage tasks

### commands
```bash
task-cli add <task-description>
task-cli update <task-id> <new-description>
task-cli delete <taask_id>
task-cli mark-in-progress <task_id>
task-cli mark-done <task_id>
task-cli list
task-cli list done
task-cli list todo
task-cli list in-progress
```

### usage
1. build via ```go build /path-to-task-cli.go```
2. (optional) add executive file to PATH