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

func swap[T any](slice *[]T) error {
    if len(*slice) == 0 {
        return errors.New("Empty slice")
    }

    (*slice)[1], (*slice)[0] =
        (*slice)[0], (*slice)[1]

    return nil
}

func flip[T any](slice *[]T) error {
    if len(*slice) == 0 {
        return errors.New("Empty slice")
    }

    (*slice)[len(*slice)-2], (*slice)[len(*slice)-1] =
        (*slice)[len(*slice)-1], (*slice)[len(*slice)-2]

    return nil
}

func deferTask[T any](slice *[]T) error {
    if len(*slice) == 0 {
        return errors.New("Empty slice")
    }

    x, xs := (*slice)[0], (*slice)[1:]
    (*slice) = append(xs, x)

    return nil
}

func undeferTask[T any](slice *[]T) error {
    if len(*slice) == 0 {
        return errors.New("Empty slice")
    }

    xs, x := (*slice)[0:len(*slice)-1], (*slice)[len(*slice)-1]
    (*slice) = append([]T{x}, xs...)

    return nil
}

func pushTask(name string, tasks *[]string) {
    var new_task string

    if name == "" {
        fmt.Print("New task: ")
        fmt.Scanln(&new_task)

        if new_task == "" {
            return
        }
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

    for i, task := range tasks {
        fmt.Printf("%v. %v\n", i+1, task)
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

        case "swap":
            swap(&tasks)

        case "flip":
            flip(&tasks)

        case "defer":
            deferTask(&tasks)

        case "undefer":
            undeferTask(&tasks)

        case "edit": // TODO?: Add more stuff to edit
            var task_id int
            var new_name string

            fmt.Printf("Number of task to edit: ")
            fmt.Scanln(&task_id)

            fmt.Printf("New name for task n.%v: ", task_id)
            fmt.Scanln(&new_name)

            if new_name != "" {
                tasks[task_id-1] = new_name
            }

        case "clear":
            tasks = nil // NOTE: Is this GC-compliant?

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
