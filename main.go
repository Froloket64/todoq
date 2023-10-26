package main

import (
    "strings"
	"errors"
	"fmt"
	"os"
)

func pop[T any](slice *[]T) (T, error) {
    if len(*slice) == 0 {
        return *new(T), errors.New("Empty slice")
    }

    popped := (*slice)[0]
    *slice = (*slice)[1:]

    return popped, nil
}

func pushTask(name string, tasks *[]string) {
    var new_task string

    if name == "" {
        fmt.Print("New task: ")
        fmt.Scanln(&new_task)
    } else {
        new_task = name
    }

    *tasks = append(*tasks, new_task)
}

func popTask(tasks *[]string) {
    popped, err := pop(tasks)

    if err == nil {
        fmt.Printf("Task %#v completed!\n", popped)
    } else {
        fmt.Println("Queue is empty.")
    }
}

func listTasks(tasks []string) {
    fmt.Println("Current task queue:")

    for _, task := range tasks {
        fmt.Printf("- %v\n", task)
    }

    fmt.Println("")
}

func loadTasks(filename string, tasks *[]string) {
    tasks_save, err := os.ReadFile(os.ExpandEnv(filename))

    if err == nil {
        tasks_fmt := string(tasks_save)
        tasks_fmt = strings.TrimSpace(tasks_fmt)

        if tasks_fmt != "" {
            *tasks = strings.Split(tasks_fmt, " ")
        }
    } else  {
        tasks = nil
    }
}

func saveTasks(filename string, tasks []string) {
    var tasks_fmt string
    tasks_fmt = strings.Join(tasks, " ") + "\n"

    os.WriteFile(os.ExpandEnv("$HOME/todoq.tsk"), []byte(tasks_fmt), 0644)
}

func main() {
    var tasks []string
    var cmd string
    var arg1 string

    loadTasks("$HOME/todoq.tsk", &tasks)

    Loop: for {
        fmt.Print("> ")
        fmt.Scanln(&cmd, &arg1)

        switch cmd {
        case "list":
            listTasks(tasks)

        case "push":
            pushTask(arg1, &tasks)

        case "pop":
            popTask(&tasks)

        case "exit", "quit", "q":
            saveTasks("$HOME/todoq.tsk", tasks)

            break Loop

        case "": // Ignore

        default:
            fmt.Printf("Unknown command: %v\n", cmd)
        }

        cmd = ""
        arg1 = ""
    }

    fmt.Println("")
}
