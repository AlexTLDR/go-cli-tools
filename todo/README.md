# Usage:

```
go run main.go -task "name of your task"
go run main.go -list
go run main.go -complete 1
```

## Using env variables
```
export TODO_FILENAME=new-todo.json
unset TODO_FILENAME
```

## Capturing Input from STDIN
```
go run main.go -add "item from Args"
echo "this one is from STDIN via the echo command" | go run main.go -add
```