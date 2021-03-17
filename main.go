package main

import (
	"bufio"
	"fmt"
	"os"
)

var appName = "tasks2todotxt"

func main() {
	if len(os.Args) != 0 {
		appName = os.Args[0]
	}

	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "usage: %s tasks-backup-file.json [Todo.txt]\n", appName)
		os.Exit(1)
	}

	sink := os.Stdout
	if len(os.Args) > 2 {
		var err error
		sink, err = os.Create(os.Args[2])
		handleError(err)
		defer sink.Close()
	}
	output := bufio.NewWriter(sink)
	defer output.Flush()

	backupData, err := LoadBackupFile(os.Args[1])
	handleError(err)

	todoTxt := backupData.ConvertToTodoTxt()

	for _, line := range todoTxt {
		err := line.WriteTo(output)
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", appName, err)
		os.Exit(2)
	}
}
