package todolist

import "errors"

// Task represents a task in the todo list.
type Task struct {
    ID       int
    Title    string
    Completed bool
}

// TodoList represents a collection of tasks.
type TodoList struct {
    tasks []*Task
}

// NewTodoList creates a new todo list.
func NewTodoList() *TodoList {
    return &TodoList{}
}

// AddTask adds a new task to the todo list.
func (t *TodoList) AddTask(title string) {
    task := &Task{
        ID:       len(t.tasks) + 1,
        Title:    title,
        Completed: false,
    }
    t.tasks = append(t.tasks, task)
}

// MarkTaskCompleted marks a task with the given ID as completed.
func (t *TodoList) MarkTaskCompleted(taskID int) error {
    for _, task := range t.tasks {
        if task.ID == taskID {
            task.Completed = true
            return nil
        }
    }
    return errors.New("task not found")
}

// DeleteTask deletes a task with the given ID from the todo list.
func (t *TodoList) DeleteTask(taskID int) error {
    for i, task := range t.tasks {
        if task.ID == taskID {
            t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}

// ListTasks returns a list of all tasks in the todo list.
func (t *TodoList) ListTasks() []*Task {
    return t.tasks
}
