package main

// Task - represent structure of a single task
type Task struct {
	ID          int
	Title       string
	Description string
	Command     string
	Filename    string
}

// Manifest - represent structure of tutorial manifest
type Manifest struct {
	Title string
	Repo  string
	Tasks []Task
}
