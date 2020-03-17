package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
)

func startTesting() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("There was an error running the tests: %s", err)
	}

	// read tutorial-mapping.json and start executing tests
	tutorialFilePath := path.Join(pwd, "tutorial-manifest.json")
	// manifest struct instance
	var metadata Manifest

	// open the manifest file and store its content in data
	data, err := ioutil.ReadFile(tutorialFilePath)
	if err != nil {
		log.Fatalf("There was an error reading the manifest file: %s", err)
	}

	// unmarhsall the data obtained into a Manifest struct
	err = json.Unmarshal(data, &metadata)
	if err != nil {
		log.Fatalf("There was an error in the manifest file: %s", err)
	}

	// Inform user about start of tests
	printTitle(fmt.Sprintf("Running tests for \"%s\"\n", metadata.Title))

	// run test for each task
	for index, task := range metadata.Tasks {
		test(pwd, task, index+1)
		fmt.Println() // Print newline
	}
}

func test(pwd string, task Task, index int) {
	fileName := path.Join(pwd, "tasks", strconv.Itoa(index), task.Filename)

	cmd := exec.Command(task.Command, fileName)

	// Set stderr and stdout variables
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	// run command and capture error
	err := cmd.Run()

	// print status of test
	if err != nil {
		printFailed()
	} else {
		printPassed()
	}

	// print test title
	fmt.Printf("  %s\n", task.Title)

	// print stderr from failing test
	if err != nil {
		printStderr(stdErr.String())
	}
}
