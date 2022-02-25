# Tasks2Todo.txt

This is a simple utility to convert [Tasks](https://tasks.org) backup files to the [Todo.txt file format](https://github.com/todotxt/todo.txt).

## Exporting a backup file

In Tasks, go to Menu ⇒ Settings ⇒ Backups and [create a backup](https://tasks.org/docs/backups/).

Copy the resulting file to your computer.

If you have Tasks configured to copy backups to your Google Drive account, it can be found in your "Tasks Backups" folder. Automatic backups are named `auto.<date>.<time>.json` and manual backups are named `user.<date>.<time>.json`. You'll probably want the most recent one.

If you don't have that configured (and don't want to do so), you'll need to figure out another way to get the file to your computer.

## Usage

Make sure you have [Go](https://go.dev) installed.

Then, from this folder, run `go build`. This will produce the executable you need.

It accepts two command-line parameters:

- the input file: a backup file from Tasks in JSON format (see above on how to export it).
- (optionally) the output file, which will be a text file in Todo.txt format. If this parameter is omitted the output is sent to the standard output (the console, unless redirected).

## Output

Tasks are exported in the order they appear in the backup file, which appears to be sorted by creation date.

The Todo.txt format does not support all features of Tasks. This is handled as follows:

- For tasks with notes they will be appended to the title, separated by `:::`. Newlines in notes are replaced by `\n` because tasks must be a single line.
- For repeating tasks, a custom `recurrence` tag is added which specifies the [recurrence rule](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10). (This is how Tasks saves it in the backup file)
- For tasks that have a due date, a custom `due` tag is added containing this date. If the task is not due at midnight the time will be included, separated from the date by a `T`.
- Some information in the backup file that cannot easily be represented in the Todo.txt format is ignored.
