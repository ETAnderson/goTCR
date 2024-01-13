package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "gotcr",
  Short: "Run tests and commit if they pass, revert if they fail",
  Run: func(cmd *cobra.Command, args []string) {
    // Run tests
    if !runTests() {
      fmt.Println("Tests failed. Reverting changes.")
      revertChanges("Revert changes due to test failure")
      return
    }

    fmt.Println("Tests passed. Committing changes.")
    commitChanges("Commit changes after successful tests")
  },
}

func main() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func runTests() bool {
  // Run 'go test' command
  cmd := exec.Command("go", "test")
  err := cmd.Run()

  // Check if the command was successful (exit code 0)
  return err == nil
}

func commitChanges(message string) {
  // Add all changes to the staging area
  cmd := exec.Command("git", "add", ".")
  err := cmd.Run()
  if err != nil {
    fmt.Println("Error adding changes to the staging area:", err)
    return
  }

  // Commit the changes
  cmd = exec.Command("git", "commit", "-m", message)
  err = cmd.Run()
  if err != nil {
    fmt.Println("Error committing changes:", err)
    return
  }

  fmt.Println("Changes committed successfully.")
}

func revertChanges(message string) {
  // Revert changes
  cmd := exec.Command("git", "reset", "--hard", "HEAD")
  err := cmd.Run()
  if err != nil {
    fmt.Println("Error reverting changes:", err)
    return
  }

  fmt.Println("Changes reverted successfully.")
}
