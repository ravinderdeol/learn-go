// Managing directories in Go:
// - Use 'os.Mkdir(name, perm)' to create a directory with a specified name and permissions
// - Set 'perm' to '0777' for open read, write, and execute access
// - Use 'os.Getwd()' to retrieve the current working directory
// - Use 'os.Chdir(name)' to change the current working directory

// File creation in Go:
// - Use 'os.Create("name")' to create a new file
// - Use 'defer File.Close()' to release system resources and save changes after file operations

// Reading files in Go:
// - Use 'os.ReadFile("name")' to read the entire contents of a file
// - Convert file content to a string using 'string(content)' before printing it with 'fmt.Println'

// Writing files in Go:
// - Use 'File.WriteString("string")' to append a string to a file
// - Use 'File.Write(data)' to write byte data such as JSON objects into a file

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User represents a user profile with name, age, and email
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// createDirectory creates a new directory with full permissions
func createDirectory(dirName string) error {
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dirName, err)
	}
	fmt.Printf("Directory '%s' created successfully.\n", dirName)
	return nil
}

// changeDirectory changes the current working directory to the specified directory
func changeDirectory(dirName string) error {
	err := os.Chdir(dirName)
	if err != nil {
		return fmt.Errorf("failed to change directory to '%s': %w", dirName, err)
	}
	fmt.Printf("Changed to directory: %s\n", dirName)
	return nil
}

// getWorkingDirectory returns the current working directory and prints it
func getWorkingDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}
	fmt.Printf("Current working directory: %s\n", cwd)
	return cwd, nil
}

// createFile creates a file and writes the specified content to it
func createFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file '%s': %w", fileName, err)
	}

	fmt.Printf("File '%s' created with initial content.\n", fileName)
	return nil
}

// writeUserData writes a User struct as JSON data to a file
func writeUserData(fileName string, user User) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", fileName, err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to serialize user to JSON: %w", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write JSON data to file '%s': %w", fileName, err)
	}

	fmt.Printf("User data written to '%s'.\n", fileName)
	return nil
}

// readFile reads the content of a file and prints it
func readFile(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %w", fileName, err)
	}
	fmt.Printf("Contents of '%s':\n%s\n", fileName, string(content))
	return nil
}

func main() {
	// Create a directory for user profiles
	err := createDirectory("user_profiles")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Change the current working directory to the new directory
	err = changeDirectory("user_profiles")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the current working directory
	_, err = getWorkingDirectory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a file and write initial content to it
	err = createFile("welcome.txt", "Welcome to the User Profiles Directory!")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a User struct and write its data to a JSON file
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	err = writeUserData("user_data.json", user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and print the content of the JSON file
	err = readFile("user_data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
