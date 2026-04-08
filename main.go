package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)


type Task struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"` 	
}

func CreateTaskJSONFile() {
	_, err := os.Stat("tasks.json")
	if err != nil {
		fmt.Println("New task file created....")
		file, errFile := os.Create("tasks.json")
		if errFile != nil {
			panic(errFile)
		}
		defer file.Close()
		
		_, errINWrite := file.Write([]byte("[]"))
		if errINWrite != nil {
			fmt.Println(errINWrite)
		}
	}
}

func ReadTaskJSONFile() {
	data, errINRead := os.ReadFile("tasks.json")
	if errINRead != nil {
		panic(errINRead)
	}

	var task []Task
	errINJson := json.Unmarshal(data, &task)
	if errINJson != nil {
		fmt.Println("Json:", errINJson)
	}
	
	var payload Task
	payload.Id = 1
	payload.Description = "CLAN EQ"
	payload.Status = "todo"
	payload.CreatedAt = "2025-02-01"
	payload.UpdatedAt = "2025-02-01"

	newTasks := append(task, payload)

	_, errINOpen := os.Open("tasks.json")
	if errINOpen != nil {
		panic(errINOpen)
	}

	jsonData, errINJsonEncode := json.Marshal(newTasks)
	if errINJsonEncode != nil {
		panic(errINJsonEncode)
	}

	errInWrite := os.WriteFile("tasks.json", jsonData, 0644)
	if errInWrite != nil {
		panic(errInWrite)
	}

	fmt.Println("Add task in file...")
}

func getFileData() []Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	var task []Task
	errINJson := json.Unmarshal(data, &task)
	if errINJson != nil {
		fmt.Println("Json:", errINJson)
	}
	
	return task
}

func AddTask(desc string){
	var task Task

	fileData := getFileData()

	task.Id = len(fileData)+1
	task.Description = string(desc)
	task.Status = "todo"
	task.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	newTasks := append(fileData, task)

	finalTaskJson, errINJsonEncode := json.Marshal(newTasks)
	if errINJsonEncode != nil {
		fmt.Println(errINJsonEncode)
	}

	errInWriteFile := os.WriteFile("tasks.json", finalTaskJson, 0644)
	if errInWriteFile != nil {
		panic(errInWriteFile)
	}

	fmt.Println("Task Added...")
}

func UpdateTask(id int, desc string){
	fileData := getFileData();
	found := false

	for i := range fileData {
		if fileData[i].Id == id {
			fileData[i].Description = desc
			fileData[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task not found")
		return
	}

	jsonData, errINJsonEncode := json.Marshal(fileData)
	if errINJsonEncode != nil {
		panic(errINJsonEncode)
	}

	errInWrite := os.WriteFile("tasks.json", jsonData, 0644)
	if errInWrite != nil {
		panic(errInWrite)
	}
	
	fmt.Println("Task updated successfully")
}

func deleteTask(id int) {
	fileData := getFileData()
	var newData []Task
	found := false

	for i := range fileData {
		if fileData[i].Id == id {
			newData = append(fileData[:i], fileData[i+1:]...)
			found = true
			break
		}
	}
	
	if !found {
		fmt.Println("No Task Found...")
		return
	}

	fmt.Println(newData)
}

func changeTaskStatus(id int, status string) {
	fileData := getFileData()
	found := false

	for i := range fileData {
		if fileData[i].Id == id {
			fileData[i].Status = status
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task not found.")
		return
	}

	jsonData, errINJSONEncode := json.Marshal(fileData)
	if errINJSONEncode != nil {
		panic(errINJSONEncode)
	}

	errInWrite := os.WriteFile("tasks.json", jsonData, 0644)
	if errInWrite != nil {
		panic(errInWrite)
	}
	
	fmt.Println("Task status updated successfully")
}

func listAllTask() {
	fileData := getFileData()

	for i := range fileData {
		fmt.Println(fileData[i].Id,"|", fileData[i].Description,"|", fileData[i].Status)
	}
}

func listAllTaskByStatus(status string) {
	fileData := getFileData()
	found := false

	for i := range fileData {
		if fileData[i].Status == status {
			fmt.Println(fileData[i].Id,"|", fileData[i].Description,"|", fileData[i].Status)
			found = true
		}
	}

	if !found {
		fmt.Println("No Task found")
	}
}

func main()  {
	
	CreateTaskJSONFile()

	action := os.Args[1]

	switch action {
	case "add" :
		taskDescription := os.Args[2]

		if taskDescription != "" {
			AddTask(taskDescription)
		}else {
			fmt.Println("Please Add Task details...")
			break
		}
	
	case "update" :
		taskId := os.Args[2]
		taskDescription := os.Args[3]
		id, _ := strconv.Atoi(taskId)

		if taskId == "" {
			fmt.Println("Please Add Task Id...")
		}else if (taskDescription == "") {
			fmt.Println("Please Add Task Description")
		}else if (taskId != "" && taskDescription != ""){
			UpdateTask(id, taskDescription)
		}
	case "mark-in-progress":
		taskId := os.Args[2]
		id, _ := strconv.Atoi(taskId)
		
		if taskId != "" {
			changeTaskStatus(id, "in-progress")
		}else {
			fmt.Println("Please Add Task Id...")
		}

	case "mark-done":
		taskId := os.Args[2]
		id, _ := strconv.Atoi(taskId)
		
		if taskId != "" {
			changeTaskStatus(id, "done")
		}else {
			fmt.Println("Please Add Task Id...")
		}

	case "delete" : 
		taskId := os.Args[2]
		id, _ := strconv.Atoi(taskId)

		if taskId != "" {
			deleteTask(id)
		}else{
			fmt.Println("Please Add Task Id...")
		}
	
	case "list" :

		if len(os.Args) > 2 {
			status := os.Args[2]
			if status != "" {
				listAllTaskByStatus(status)
			}
		}else {
			listAllTask();
		}



	default: 
		fmt.Println("Invalid command please read doc using -h command")
	}
}