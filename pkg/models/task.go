package models

// import (
//     "time"
// )

type Task struct {
    ID          string    `json:"id"`
    Description string    `json:"description"`
    // DueDate     time.Time `json:"due_date"`
    Completed   bool      `json:"completed"`
}

type TaskList struct {
    Tasks []Task `json:"tasks"`
}

