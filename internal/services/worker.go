package services

import (
	"log"
	"time"

	"github.com/techymj/task-manager/internal/repositories"
)

type Worker struct {
	Repo        repositories.TaskRepository
	Queue       chan string
	DelayMinute int
}

func NewWorker(repo repositories.TaskRepository, queue chan string, delay int) *Worker {

	return &Worker{
		Repo:        repo,
		Queue:       queue,
		DelayMinute: delay,
	}
}

func (w *Worker) Start() {
	go func() {
		for taskID := range w.Queue {
			go w.handleTask(taskID)
		}
	}()
}

func (w *Worker) handleTask(taskID string) {
	log.Println("Worker received:", taskID)
	log.Println("Sleeping for", w.DelayMinute, "minutes")

	time.Sleep(time.Duration(w.DelayMinute) * time.Minute)

	log.Println("Woke up for:", taskID)

	task, err := w.Repo.GetByID(taskID)
	if err != nil {
		log.Println("Task not found:", taskID)
		return
	}

	log.Println("Current status:", task.Status)

	if task.Status == "pending" || task.Status == "in_progress" {
		log.Println("Auto completing task:", taskID)
		err := w.Repo.UpdateStatus(taskID, "completed")
		if err != nil {
			log.Println("Update failed:", err)
		} else {
			log.Println("Task marked completed")
		}
	}
}
