# Task Tracker CLI

## Usage

```txt
Usage of ./cmd/cli/:
  -action string
     	action to perform [create|read|update|delete|mark-todo|mark-in-progress|mark-done|list|list-todo|list-in-progress|list-done|list-deleted|list-all|clean]
  -id uint
     	id of the task to read/update/delete
  -description string
     	description of the task
```

## Try It Out

```sh
# Create tasks
go run ./cmd/cli/ -action=create -description="Buy groceries"
go run ./cmd/cli/ -action=create -description="Cook dinner"
go run ./cmd/cli/ -action=create -description="Repair front porch"
go run ./cmd/cli/ -action=create -description="Service car"

# Read task
go run ./cmd/cli/ -action=read -id=1

# Update task
go run ./cmd/cli/ -action=update -id=1 -description="Buy groceries and cook dinner"

# Delete task
go run ./cmd/cli/ -action=delete -id=2

# Mark tasks
go run ./cmd/cli/ -action=mark-in-progress -id=1
go run ./cmd/cli/ -action=mark-done -id=1
go run ./cmd/cli/ -action=mark-todo -id=1

# List tasks
go run ./cmd/cli/ -action=list
go run ./cmd/cli/ -action=list-todo
go run ./cmd/cli/ -action=list-in-progress
go run ./cmd/cli/ -action=list-done
go run ./cmd/cli/ -action=list-deleted
go run ./cmd/cli/ -action=list-all

# Clean tasks
go run ./cmd/cli/ -action=clean
```

> Take 🎂, Folks! 🌮 🐴 💨
