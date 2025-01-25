// HTTP Package:
// - The net/http package provides tools to create HTTP servers and clients.

// Creating a Server:
// - Use http.ListenAndServe to start a server that listens on a specific port.

// Routing:
// - Use http.NewServeMux to create a multiplexer for managing custom routing logic.
// - Define routes using mux.HandleFunc(path, handler) to map paths to handler functions.
// - A route will respond to all HTTP methods if no method is specified.

// GET Requests:
// - Define a GET route to handle only GET requests for a specific path.
// - Restricting routes to specific methods improves security by avoiding unintended request handling.

// POST Requests:
// - Define a POST route to handle only POST requests for a specific path.
// - Use io.ReadAll in the POST handler to read and process incoming JSON data as a byte slice.

// HTTP GET:
// - Use http.Get to make a GET request to a URL and retrieve its response.
// - Use http.NewRequest for more control over HTTP requests (including custom headers and query parameters).

// HTTP POST:
// - Use http.Post to make a simple POST request with a URL, content type, and body data.
// - Convert a string to a byte slice using []byte or create a buffer with bytes.NewBuffer for the request body.

// Forms and User Input:
// - Form submissions send data to the server using an HTTP POST request.
// - Access and process form data in the POST handler function.

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

// Task represents a single task.
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// In-memory storage for tasks.
var tasks []Task

func main() {
	// Set up routing using http.ServeMux.
	mux := http.NewServeMux()

	// Define routes.
	mux.HandleFunc("/", handleRoot)       // Root route.
	mux.HandleFunc("/tasks", handleTasks) // Handles GET and POST requests for tasks.
	mux.HandleFunc("/submit", handleForm) // Handles form submissions.

	// Start the server on port 4001.
	fmt.Println("Starting server on :4001...")
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// handleRoot handles the root endpoint.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Task Manager API!")
}

// handleTasks handles GET and POST requests for tasks.
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return all tasks as JSON.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		// Add a new task from JSON data.
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

// handleForm handles form submissions to add a new task.
func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data.
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		return
	}

	// Retrieve form values.
	title := r.FormValue("title")
	description := r.FormValue("description")

	// Add the task to in-memory storage.
	tasks = append(tasks, Task{Title: title, Description: description})
	fmt.Fprintln(w, "Form submitted successfully!")
}

// simulateClient demonstrates making GET and POST requests.
func simulateClient() {
	// Make a GET request to fetch tasks.
	resp, err := http.Get("http://localhost:4001/tasks")
	if err != nil {
		log.Fatal("Error making GET request:", err)
	}
	defer resp.Body.Close()

	// Read and print the response.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading GET response:", err)
	}
	fmt.Println("GET /tasks response:", string(body))

	// Prepare JSON data for a POST request.
	newTask := Task{Title: "Learn Go", Description: "Complete Go tutorials."}
	jsonData, _ := json.Marshal(newTask)

	// Make a POST request to add a new task.
	resp, err = http.Post("http://localhost:4001/tasks", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error making POST request:", err)
	}
	defer resp.Body.Close()

	fmt.Println("POST /tasks completed successfully!")
}

// simulateLogs demonstrates logging server activities to a file.
func simulateLogs() {
	// Create or open a log file.
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	defer file.Close()

	// Log server activities.
	log.SetOutput(file)
	log.Println("Server started on :4001.")
	log.Println("Handled GET request for /tasks.")
	log.Println("Handled POST request for /tasks.")
}
