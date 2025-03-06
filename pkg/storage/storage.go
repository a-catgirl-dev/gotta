package storage

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/a-catgirl-dev/gotta/pkg/models"
)

type Storage struct {
    filename string
}

func NewStorage(filename string) *Storage {
    return &Storage{
        filename: filename,
    }
}

func (s *Storage) SaveTasks(tasks []models.Task) error {
    data, err := json.MarshalIndent(tasks, "", "\t")
    if err != nil {
        return fmt.Errorf("failed to marshal tasks: %v", err)
    }
    return os.WriteFile(s.filename, data, 0644)
}

func (s *Storage) LoadTasks() ([]models.Task, error) {
    data, err := os.ReadFile(s.filename)
    if err != nil {
        return []models.Task{}, nil
    }
    var tasks []models.Task
    err = json.Unmarshal(data, &tasks)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal tasks: %v", err)
    }
    return tasks, nil
}

