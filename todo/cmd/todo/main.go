package main

import (
	"fmt"
	"github.com/AlexTLDR/go-cli-tools/todo"
	"os"
	"strings"
)

const todoFilename = ".todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}