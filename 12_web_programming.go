// HTTP package in Go:
// - The 'net/http' package provides tools to create HTTP servers and clients

// Creating a server in Go:
// - Use 'http.ListenAndServe' to start a server listening on a specific port

// Routing in Go:
// - Use 'http.NewServeMux' to create a multiplexer for custom routing logic
// - Define routes using 'mux.HandleFunc(path, handler)' to map paths to handler functions
// - Routes respond to all HTTP methods if no method is specified

// Handling GET requests in Go:
// - Define a GET route to handle only GET requests for a specific path
// - Restricting routes to specific methods improves security by avoiding unintended request handling

// Handling POST requests in Go:
// - Define a POST route to handle only POST requests for a specific path
// - Use 'io.ReadAll' in the POST handler to read and process incoming JSON data as a byte slice

// Making HTTP GET requests in Go:
// - Use 'http.Get' to make a GET request to a URL and retrieve its response
// - Use 'http.NewRequest' for more control over requests such as adding custom headers or query parameters

// Making HTTP POST requests in Go:
// - Use 'http.Post' to make a POST request with a URL, content type, and body data
// - Convert a string to a byte slice using '[]byte' or create a buffer with 'bytes.NewBuffer' for the request body

// Forms and user input in Go:
// - Form submissions send data to the server using an HTTP POST request
// - Access and process form data in the POST handler function

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Task represents a single task with a title and description
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// In-memory storage for tasks
var tasks []Task

func main() {
	// Set up routing with 'http.ServeMux'
	mux := http.NewServeMux()

	// Define the root route to display a welcome message
	mux.HandleFunc("/", handleRoot)

	// Define the '/tasks' route to handle GET and POST requests for tasks
	mux.HandleFunc("/tasks", handleTasks)

	// Define the '/submit' route to handle form submissions for adding tasks
	mux.HandleFunc("/submit", handleForm)

	// Start the server on port 4001
	fmt.Println("Starting server on :4001...")
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// handleRoot displays a welcome message on the root endpoint
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Task Manager API!")
}

// handleTasks handles GET and POST requests for task management
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return all tasks as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		// Add a new task from JSON data in the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body.", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var newTask Task
		err = json.Unmarshal(body, &newTask)
		if err != nil {
			http.Error(w, "Invalid JSON format.", http.StatusBadRequest)
			return
		}

		tasks = append(tasks, newTask)
		fmt.Fprintln(w, "Task added successfully!")
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

// handleForm processes form submissions to add a new task
func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data to retrieve task details
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")

	// Add the task to the in-memory storage
	tasks = append(tasks, Task{Title: title, Description: description})
	fmt.Fprintln(w, "Form submitted successfully!")
}

// simulateClient demonstrates making HTTP GET and POST requests
func simulateClient() {
	// Make a GET request to fetch all tasks
	resp, err := http.Get("http://localhost:4001/tasks")
	if err != nil {
		log.Fatal("Error making GET request:", err)
	}
	defer resp.Body.Close()

	// Print the response body from the GET request
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading GET response:", err)
	}
	fmt.Println("GET /tasks response:", string(body))

	// Create a new task and prepare JSON data for a POST request
	newTask := Task{Title: "Learn Go", Description: "Complete Go tutorials."}
	jsonData, _ := json.Marshal(newTask)

	// Make a POST request to add a new task
	resp, err = http.Post("http://localhost:4001/tasks", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error making POST request:", err)
	}
	defer resp.Body.Close()

	fmt.Println("POST /tasks completed successfully!")
}

// simulateLogs demonstrates logging server activities to a file
func simulateLogs() {
	// Open or create a log file to store server activities
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	defer file.Close()

	// Write server logs to the file
	log.SetOutput(file)
	log.Println("Server started on :4001.")
	log.Println("Handled GET request for /tasks.")
	log.Println("Handled POST request for /tasks.")
}
