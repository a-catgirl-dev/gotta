package cli

import (
    "fmt"
    "slices"
    "strconv"
    "strings"
    "time"

    "github.com/a-catgirl-dev/gotta/pkg/models"
    "github.com/a-catgirl-dev/gotta/pkg/storage"
)

type CLI struct {
    storage storage.Storage
}

func NewCLI(storage storage.Storage) *CLI {
    return &CLI{
        storage: storage,
    }
}

func (c *CLI) Run(args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("gotta do what?\nfor usage guide, execute `gotta help`")
    }

    switch strings.ToLower(args[0]) {
        case "help":
            return fmt.Errorf("figure it out.")
        case "do":
            return c.addTask(args)
        case "list":
            return c.listTasks()
        case "checkoff": // gotta x?
            return c.completeTask(args)
        // gotta rm? gotta drop?
        case "drop":
            return c.dropTask(args)
        default:
            return fmt.Errorf("unknown command: %s", args[0])
    }
}

func (c *CLI) addTask(args []string) error {
    if len(args) != 3 {
        return fmt.Errorf("usage: gotta do <description> <when>")
    }

    task := models.Task{
        ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
        Description: args[1],
        // DueDate:     TODO: implement (im lazy)
        Completed:   false,
    }

    tasks, err := c.storage.LoadTasks()
    if err != nil {
        return err
    }
    tasks = append(tasks, task)

    return c.storage.SaveTasks(tasks)
}

func (c *CLI) dropTask(args []string) error {
    if len(args) != 2 {
        return fmt.Errorf("usage: gotta drop <index>")
    }

    tasks, err := c.storage.LoadTasks()
    if err != nil {
        return err
    }

    index, err := strconv.Atoi(args[1])
    if err != nil {
        return fmt.Errorf("invalid index: %v", err)
    }

    if index < 0 || index >= len(tasks) {
        return fmt.Errorf("task index %d out of range (valid indices: 0-%d)", 
            index, len(tasks)-1)
    }

    tasks = slices.Delete(tasks, index, index+1)

    return c.storage.SaveTasks(tasks)
}

func (c *CLI) completeTask(args []string) error {
    if len(args) != 2 {
        return fmt.Errorf("usage: gotta checkoff <index>")
    }

    tasks, err := c.storage.LoadTasks()
    if err != nil {
        return err
    }

    index, err := strconv.Atoi(args[1])
    if err != nil {
        return fmt.Errorf("invalid index: %v", err)
    }

    if index < 0 || index >= len(tasks) {
        return fmt.Errorf("task index %d out of range (valid indices: 0-%d)", index, len(tasks)-1)
    }

    tasks[index].Completed = true

    return c.storage.SaveTasks(tasks)
}

func (c *CLI) listTasks() error {
    tasks, err := c.storage.LoadTasks()
    if err != nil {
        return err
    }

    fmt.Println("Current Tasks:")
    // fmt.Println("Index | Status | Description | Due date")
    fmt.Println("Index | Status | Description")
    fmt.Println("----- | ------ | -----------")

    for i, task := range tasks {
        status := "[✓]"
        if !task.Completed {
            status = "[ ]"
        }
        fmt.Printf("%5d | %6s | %s\n", i, status, task.Description)
    }
    return nil
}

