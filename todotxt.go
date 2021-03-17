package main

import (
	"bytes"
	"io"
	"sort"
	"strings"
	"time"
)

const todoTxtDateFormat = "2006-01-02"

// TodoTxtTask represents a todo.txt task (see https://github.com/todotxt/todo.txt)
type TodoTxtTask struct {
	Completed      bool
	Priority       byte // 0 = None, 1-26 = (A)-(Z)
	CompletionDate *time.Time
	CreationDate   *time.Time
	Description    string
	Tags           []string
	Contexts       []string
	CustomTags     map[string]string
}

func (t *TodoTxtTask) String() string {
	b := &bytes.Buffer{}
	_ = t.WriteTo(b)
	return strings.TrimSpace(b.String())
}

type writer interface {
	io.ByteWriter
	io.StringWriter
}

// WriteTo writes a string representation of this object to the writer.
func (t *TodoTxtTask) WriteTo(w writer) error {
	if t.Completed {
		w.WriteString("x ")
	}

	if t.Priority != 0 {
		w.WriteByte('(')
		w.WriteByte('A' + t.Priority - 1)
		w.WriteString(") ")
	}

	// Helper function for writing out optional dates.
	writeDate := func(date *time.Time) bool {
		if date != nil {
			w.WriteString(date.Format(todoTxtDateFormat))
			w.WriteByte(' ')
		}
		return date != nil
	}

	// Completion date is only allowed if creation date is present.
	if t.CreationDate != nil {
		writeDate(t.CompletionDate)
	}
	writeDate(t.CreationDate)

	// Description
	w.WriteString(t.Description)

	// Tags
	for _, tag := range t.Tags {
		w.WriteString(" +")
		w.WriteString(tag)
	}

	// Context tags
	for _, tag := range t.Contexts {
		w.WriteString(" @")
		w.WriteString(tag)
	}

	// Add the custom tags, in alphabetical order
	keys := []string{}
	for key := range t.CustomTags {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		w.WriteByte(' ')
		w.WriteString(key)
		w.WriteByte(':')
		w.WriteString(t.CustomTags[key])
	}

	return w.WriteByte('\n')
}
