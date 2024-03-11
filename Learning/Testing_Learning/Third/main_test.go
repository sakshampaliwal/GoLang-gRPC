package todolist

import (
	"reflect"
	"testing"
)

func TestTodoList_AddTask(t *testing.T) {
	todo := NewTodoList()
	todo.AddTask("Buy groceries")
	todo.AddTask("Finish project")

	if len(todo.tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(todo.tasks))
	}
}

func TestTodoList_MarkTaskCompleted(t *testing.T) {
	todo := NewTodoList()
	todo.AddTask("Buy groceries")
	todo.AddTask("Finish project")
	err := todo.MarkTaskCompleted(2)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !todo.tasks[1].Completed {
		t.Error("Task should be marked as completed")
	}
}

func TestTodoList_DeleteTask(t *testing.T) {
	todo := NewTodoList()
	todo.AddTask("Buy groceries")
	todo.AddTask("Finish project")
	err := todo.DeleteTask(1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(todo.tasks) != 1 {
		t.Errorf("Expected 1 task after deletion, got %d", len(todo.tasks))
	}

	if !reflect.DeepEqual(todo.tasks[0], &Task{ID: 2, Title: "Finish project", Completed: false}) {
		t.Error("Incorrect task remaining after deletion")
	}
}

func TestTodoList_ListTasks(t *testing.T) {
	todo := NewTodoList()
	todo.AddTask("Buy groceries")
	todo.AddTask("Finish project")

	tasks := todo.ListTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	expectedTasks := []*Task{
		{ID: 1, Title: "Buy groceries", Completed: false},
		{ID: 2, Title: "Finish project", Completed: false},
	}

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Errorf("Tasks not as expected. Got: %v, Expected: %v", tasks, expectedTasks)
	}
}
