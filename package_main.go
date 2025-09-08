package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoList struct {
	tasks []Task
	mutex sync.Mutex
	nextID int
}

var todoList TodoList

func main() {
	todoList = TodoList{tasks: []Task{}, nextID: 1}

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/tasks", HandleTasks)
	http.HandleFunc("/health", HealthCheck)

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Todo List App!")
}

// HandleTasks manages different HTTP methods for task operations
func HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTasks(w, r)
	case "POST":
		addTask(w, r)
	case "PUT":
		updateTask(w, r)
	case "DELETE":
		deleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getTasks retrieves all tasks from the todo list
func getTasks(w http.ResponseWriter, r *http.Request) {
	todoList.mutex.Lock()
	defer todoList.mutex.Unlock()
	json.NewEncoder(w).Encode(todoList.tasks)
}

// addTask creates a new task and adds it to the todo list
func addTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todoList.mutex.Lock()
	task.ID = todoList.nextID
	todoList.nextID++
	todoList.tasks = append(todoList.tasks, task)
	todoList.mutex.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// updateTask modifies an existing task in the todo list
func updateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todoList.mutex.Lock()
	defer todoList.mutex.Unlock()
	for i, task := range todoList.tasks {
		if task.ID == updatedTask.ID {
			todoList.tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// deleteTask removes a task from the todo list
func deleteTask(w http.ResponseWriter, r *http.Request) {
	var taskID struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&taskID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todoList.mutex.Lock()
	defer todoList.mutex.Unlock()
	for i, task := range todoList.tasks {
		if task.ID == taskID.ID {
			todoList.tasks = append(todoList.tasks[:i], todoList.tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}