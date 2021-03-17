package main

import (
	"strings"
)

// ConvertToTodoTxt converts these tasks to the Todo.txt format.
func (t TasksBackupJSON) ConvertToTodoTxt() []TodoTxtTask {
	var result []TodoTxtTask
	for _, task := range t.Data.Tasks {
		result = append(result, task.ConvertToTodoTxt())
	}

	return result
}

const tagDateFormat = "2006-01-02"
const tagDateTimeFormat = "2006-01-02T15:04"

// ConvertToTodoTxt converts this task to the Todo.txt format.
func (t TaskMeta) ConvertToTodoTxt() TodoTxtTask {
	task := t.Task

	result := TodoTxtTask{
		Completed:      task.CompletionDate != 0,
		Priority:       byte(task.Priority + 1),
		CompletionDate: tasksDate(task.CompletionDate),
		CreationDate:   tasksDate(task.CreationDate),
		Description:    task.Title,
		Contexts:       nil,
		CustomTags:     make(map[string]string),
	}

	if task.Notes != "" {
		notes := strings.TrimSpace(task.Notes)
		result.Description += " ::: " + strings.ReplaceAll(notes, "\n", "\\n")
	}

	for _, tag := range t.Tags {
		result.Tags = append(result.Tags, tag.Name)
	}

	if dueDate := tasksDate(task.DueDate); dueDate != nil {
		format := tagDateFormat
		if dueDate.Hour() != 0 || dueDate.Minute() != 0 {
			format = tagDateTimeFormat
		}
		result.CustomTags["due"] = dueDate.Format(format)
	}

	if task.Recurrence != "" {
		result.CustomTags["recurrence"] = task.Recurrence
	}

	return result
}
