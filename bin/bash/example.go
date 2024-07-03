package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Paths to the required files
	participantSolutionPath := "/Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-26-code"
	jurySolutionPath := "/Users/amangupta/Desktop/cf-stress/backend/bin/contest/1746/b/solution"
	generatorPath := "/Users/amangupta/Desktop/cf-stress/backend/bin/contest/1746/b/generator"
	inputFilePath := "/Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-26-input.txt"
	participantOutputFilePath := "/Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-26-output-participant.txt"
	juryOutputFilePath := "/Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-26-output-jury.txt"

	// Arguments to pass to the bash script
	args := []string{
		fmt.Sprintf("%s.cpp", participantSolutionPath),
		fmt.Sprintf("%s.cpp", jurySolutionPath),
		fmt.Sprintf("%s.cpp", generatorPath),
		inputFilePath,
		participantOutputFilePath,
		juryOutputFilePath,
		participantSolutionPath,
		jurySolutionPath,
		generatorPath,
	}

	// Command to execute the script
	cmd := exec.Command("./checker.sh", args...)

	// To get combined output (stdout and stderr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running stress test:", err)
	}
}
