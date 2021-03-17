package main

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

// TasksBackupJSON is the JSON format of the Tasks backup data.
type TasksBackupJSON struct {
	// Skips some fields we don't care about.

	Data struct {
		Tasks []TaskMeta
	}
}

// TaskMeta wraps the main task data as well as some other data we don't care about.
type TaskMeta struct {
	// Skips some fields we don't care about.

	Tags []struct {
		Name string
	}
	Task Task
}

// Task is the main task data.
type Task struct {
	// Skips some fields we don't care about.

	Title          string
	Recurrence     string
	Notes          string
	DueDate        int64
	CreationDate   int64
	CompletionDate int64
	Priority       int
}

// LoadBackupFile loads a Tasks backup file.
func LoadBackupFile(filename string) (TasksBackupJSON, error) {
	file, err := os.Open(filename)
	if err != nil {
		return TasksBackupJSON{}, err
	}
	defer file.Close()

	return ParseBackup(file)
}

// ParseBackup parses Tasks backup data.
func ParseBackup(input io.Reader) (result TasksBackupJSON, err error) {
	err = json.NewDecoder(input).Decode(&result)
	return
}

// tasksDate converts a Tasks time stamp to time.Time.
func tasksDate(timestamp int64) *time.Time {
	if timestamp == 0 {
		return nil
	}
	nanoSeconds := (timestamp % 1000) * int64(time.Millisecond/time.Nanosecond)
	date := time.Unix(timestamp/1000, nanoSeconds).Local()
	return &date
}
