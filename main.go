package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"feature_mvvm_gen/stacks"
)

var availableStacks = []stacks.Stack{
	stacks.FlutterStack{},
	stacks.KotlinStack{},
	stacks.SwiftStack{},
	stacks.ReactStack{},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	stack := promptStack(reader)
	featureName := promptString(reader, "Feature name: ")
	if featureName == "" {
		fatal("feature name cannot be empty")
	}
	includeModels := promptBool(reader, "Include models? (y/n): ")
	includeWidgets := promptBool(reader, "Include widgets? (y/n): ")

	cwd, err := os.Getwd()
	if err != nil {
		fatal("cannot determine current directory: %v", err)
	}

	opts := stacks.Options{
		IncludeModels:  includeModels,
		IncludeWidgets: includeWidgets,
	}

	fmt.Printf("\nGenerating %s feature '%s'...\n", stack.Name(), featureName)

	if err := stack.Generate(cwd, featureName, opts); err != nil {
		fatal("generation failed: %v", err)
	}

	fmt.Println("Done! Feature scaffold created successfully.")
}

func promptStack(r *bufio.Reader) stacks.Stack {
	fmt.Println("Select a stack:")
	for i, s := range availableStacks {
		fmt.Printf("  %d. %s\n", i+1, s.Name())
	}
	for {
		fmt.Print("Enter number: ")
		input := strings.TrimSpace(readLine(r))
		for i, s := range availableStacks {
			if input == fmt.Sprintf("%d", i+1) {
				return s
			}
		}
		fmt.Printf("Invalid choice '%s'. Please enter a number between 1 and %d.\n", input, len(availableStacks))
	}
}

func promptString(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	return strings.TrimSpace(readLine(r))
}

func promptBool(r *bufio.Reader, prompt string) bool {
	for {
		fmt.Print(prompt)
		input := strings.ToLower(strings.TrimSpace(readLine(r)))
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		}
		fmt.Println("Please enter 'y' or 'n'.")
	}
}

func readLine(r *bufio.Reader) string {
	line, err := r.ReadString('\n')
	if err != nil {
		fatal("failed to read input: %v", err)
	}
	return strings.TrimRight(line, "\r\n")
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}
