package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var taskList []Task

type Task struct {
	ID                   string
	Name                 string
	Status               string
	Deadline             time.Time
	CreatedAt, UpdatedAt time.Time
}

type Person struct {
	position string
}

func (p *Person) give_task(str string, task_step string) (string, error) {
	if p.position != "director" {
		fmt.Println("Do not have permission")
		return "-1", errors.New("Do not have persmission!")
	}
	id := uuid.New()

	task := Task{ID: id.String(), Name: str, Status: task_step, CreatedAt: time.Now(), UpdatedAt: time.Now(), Deadline: time.Date(2022, time.November, 11, 23, 0, 0, 0, time.UTC)}

	taskList = append(taskList, task)

	return id.String(), nil
}

func (p *Person) deligate_task(id string, task_step string) (string, error) {
	if p.position != "team lead" {
		fmt.Println(p.position, "do not have permission!")
		return "-1", errors.New("Do not have permission!")
	} else if p.position == "team lead" && task_step != "dev" && task_step != "initial" {
		fmt.Println("You cannot change the step of task!")
		return "-1", errors.New("You cannot change the step of task!")
	}

	isExist := false

	for idx, val := range taskList {
		if val.ID == id {
			if val.Status == "initial" && task_step == "done" {
				fmt.Println("You can not change step from initial to done!")
				return "-1", errors.New("You can not change step from initial to done!")
			}
			isExist = true
			taskList[idx].Status = task_step
		}
	}

	if !isExist {
		fmt.Println("Task not found!")
		return "-1", errors.New("Task not found!")
	}

	return task_step, nil

}

func (p *Person) develop(id string, task_step string) (string, error) {
	if p.position != "developer" {
		return "-1", errors.New("Do not have permission!")
	} else if p.position == "developer" && task_step != "done" && task_step != "test" {
		fmt.Println("You cannot change the step of task!")
		return "-1", errors.New("You cannot change the step of task!")
	}
	isExist := false

	for idx, val := range taskList {
		if val.ID == id {
			if val.Status == "initial" && task_step == "done" {
				fmt.Println("You can not change step from initial to done!")
				return "-1", errors.New("You can not change step from initial to done!")
			}
			isExist = true
			taskList[idx].Status = task_step
		}
	}

	if !isExist {
		fmt.Println("Task not found!")
		return "-1", errors.New("Task not found!")
	}

	return task_step, nil
}

func OopTask() {
	director := Person{position: "director"}
	developer := Person{position: "developer"}
	teamLead := Person{position: "team lead"}

	id, _ := director.give_task("Login and sign up page!", "initial")
	teamLead.deligate_task(id, "dev")
	developer.develop(id, "test")
	developer.develop(id, "done")
	fmt.Println(taskList)

}
