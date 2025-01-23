// Making directories:
// - Use os.Mkdir(name, perm) to create directories in Go.
// - The "name" parameter specifies the directory name, and "perm" sets permission bits.
// - Use 0777 for open read, write, and execute access.

// Changing directories:
// - Use os.Getwd() to get the current working directory.
// - Use os.Chdir(name) to change the current working directory.

// File creation:
// - Use os.Create("name") to create a new file in Go.
// - Use defer File.Close() to release system resources and save changes after file operations.

// Reading from files:
// - Use os.ReadFile("name") to read the entire contents of a file.
// - Convert the content to a string using string(content) before printing it with fmt.Println.

// Writing to files:
// - Use File.WriteString("string") to append a string to a file.
// - Use File.Write(data) to write byte data, such as JSON objects, into a file.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User represents a user profile.
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// createDirectory creates a new directory.
func createDirectory(dirName string) error {
	err := os.Mkdir(dirName, 0777) // Create a directory with full permissions.
	if err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dirName, err)
	}
	fmt.Printf("Directory '%s' created successfully.\n", dirName)
	return nil
}

// changeDirectory changes the current working directory.
func changeDirectory(dirName string) error {
	err := os.Chdir(dirName)
	if err != nil {
		return fmt.Errorf("failed to change directory to '%s': %w", dirName, err)
	}
	fmt.Printf("Changed to directory: %s\n", dirName)
	return nil
}

// getWorkingDirectory prints the current working directory.
func getWorkingDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}
	fmt.Printf("Current working directory: %s\n", cwd)
	return cwd, nil
}

// createFile creates a file and writes initial content.
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

// writeUserData writes a User struct to a JSON file.
func writeUserData(fileName string, user User) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", fileName, err)
	}
	defer file.Close()

	// Serialize user to JSON and write it to the file.
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

// readFile reads and prints the content of a file.
func readFile(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %w", fileName, err)
	}
	fmt.Printf("Contents of '%s':\n%s\n", fileName, string(content))
	return nil
}

func main() {
	// Step 1: Create a directory for user profiles.
	err := createDirectory("user_profiles")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 2: Change to the new directory.
	err = changeDirectory("user_profiles")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 3: Print the current working directory.
	_, err = getWorkingDirectory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 4: Create a file and write initial content.
	err = createFile("welcome.txt", "Welcome to the User Profiles Directory!")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 5: Create and write user data to a JSON file.
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	err = writeUserData("user_data.json", user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 6: Read and print the content of the JSON file.
	err = readFile("user_data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
