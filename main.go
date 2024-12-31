package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

func readTasks() ([]Task, error) {
	f, err := os.OpenFile("tasks.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	resp, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return tasks, err
	}
	err = json.Unmarshal(resp, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func writeTasks(tasks []Task) error {
	f, err := os.OpenFile("tasks.json", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	rs, err := json.MarshalIndent(tasks, " ", "  ")
	if err != nil {
		return err
	}
	f.Write(rs)
	return nil
}

func addTask(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("you just need to inform description! no more information")
		return
	}
	if len(args) == 0 {
		fmt.Println("Not enough arguments to call this function!")
		return
	}
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	newtask := &Task{
		Id:          len(tasks) + 1,
		Description: args[0],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Date(1, 1, 1, 0, 0, 0, 1, time.UTC),
	}
	tasks = append(tasks, *newtask)
	err = writeTasks(tasks)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("Task added successfuly")
}

func updateTask(cmd *cobra.Command, args []string) {
	if len(args) > 2 {
		fmt.Println("you just need to inform id and new description! no more information")
		return
	}
	if len(args) < 2 {
		fmt.Println("Not enough arguments to call this function!")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The id you informed must be integer!")
		return
	}
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	for i, task := range tasks {
		if task.Id == id {
			if task.DeletedAt.After(task.CreatedAt) {
				fmt.Println("Task deleted before! no more actions allowed!")
				return
			}
			newtask := task
			newtask.Description = args[1]
			newtask.UpdatedAt = time.Now()
			secpart := tasks[i+1:]
			tasks = append(tasks[:i], newtask)
			tasks = append(tasks, secpart...)
			err = writeTasks(tasks)
			if err != nil {
				fmt.Println("Error occured: ", err)
				return
			}
			fmt.Println("task updated successfuly")
			return
		}
	}
	fmt.Println("There is no task with given id")
}

func deleteTask(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("you just need to inform id! no more information")
		return
	}
	if len(args) == 0 {
		fmt.Println("Not enough arguments to call this function!")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The id you informed must be integer!")
		return
	}
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	for i, task := range tasks {
		if task.Id == id {
			if task.DeletedAt.After(task.CreatedAt) {
				fmt.Println("Task deleted before! no more actions allowed!")
				return
			}
			newtask := task
			newtask.DeletedAt = time.Now()
			secpart := tasks[i+1:]
			tasks = append(tasks[:i], newtask)
			tasks = append(tasks, secpart...)
			err = writeTasks(tasks)
			if err != nil {
				fmt.Println("Error occured: ", err)
				return
			}
			fmt.Println("task deleted successfuly")
			return
		}
	}
	fmt.Println("There is no task with given id")
}

func showstatus(status string) {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("Showing tasks with status:", status)
	fmt.Println()
	for _, task := range tasks {
		if task.Status == status && !task.DeletedAt.After(task.CreatedAt) {
			fmt.Println(task.Id, task.Description)
		}
	}
}

func listTasks(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("Too many aguments to call this function!")
		return
	}
	if len(args) == 1 {
		status := strings.ToLower(args[0])
		switch status {
		case "done":
			showstatus("done")
		case "todo":
			showstatus("todo")
		case "in-progress":
			showstatus("in-progress")
		default:
			fmt.Println("Given status is unavailable!")
		}
		return
	}
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("Showing all tasks:")
	fmt.Println()
	for _, task := range tasks {
		if !task.DeletedAt.After(task.CreatedAt) {
			fmt.Printf("%d %s with status of %s\n", task.Id, task.Description, task.Status)
		}
	}
}

func markDone(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("you just need to inform id! no more information")
		return
	}
	if len(args) == 0 {
		fmt.Println("Not enough arguments to call this function!")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The id you informed must be integer!")
		return
	}
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	for i, task := range tasks {
		if task.Id == id {
			if task.DeletedAt.After(task.CreatedAt) {
				fmt.Println("Task deleted before! no more actions allowed!")
				return
			}
			newtask := task
			newtask.Status = "done"
			newtask.UpdatedAt = time.Now()
			secpart := tasks[i+1:]
			tasks = append(tasks[:i], newtask)
			tasks = append(tasks, secpart...)
			err = writeTasks(tasks)
			if err != nil {
				fmt.Println("Error occured: ", err)
				return
			}
			fmt.Println("task updated successfuly")
			return
		}
	}
	fmt.Println("There is no task with given id")
}

func markInProgress(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("you just need to inform id! no more information")
		return
	}
	if len(args) == 0 {
		fmt.Println("Not enough arguments to call this function!")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The id you informed must be integer!")
		return
	}
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	for i, task := range tasks {
		if task.Id == id {
			if task.DeletedAt.After(task.CreatedAt) {
				fmt.Println("Task deleted before! no more actions allowed!")
				return
			}
			newtask := task
			newtask.Status = "in-progress"
			newtask.UpdatedAt = time.Now()
			secpart := tasks[i+1:]
			tasks = append(tasks[:i], newtask)
			tasks = append(tasks, secpart...)
			err = writeTasks(tasks)
			if err != nil {
				fmt.Println("Error occured: ", err)
				return
			}
			fmt.Println("task updated successfuly")
			return
		}
	}
	fmt.Println("There is no task with given id")
}

func main() {
	rootcmd := &cobra.Command{Use: "task-cli"}
	addcmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to app",
		Run:   addTask,
	}
	updatecmd := &cobra.Command{
		Use:   "update",
		Short: "Update task description",
		Run:   updateTask,
	}
	deletecmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Run:   deleteTask,
	}
	listcmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks with or without filter",
		Run:   listTasks,
	}
	mdcmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark the task as done",
		Run:   markDone,
	}
	micmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark the task as in-progress",
		Run:   markInProgress,
	}
	rootcmd.AddCommand(addcmd, updatecmd, deletecmd, listcmd, mdcmd, micmd)
	if err := rootcmd.Execute(); err != nil {
		println("Error starting command: ", err)
	}
}
