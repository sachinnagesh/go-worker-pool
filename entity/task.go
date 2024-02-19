package entity

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Task struct {
	ID int
}

func (t *Task) Process() {
	log.Info("Processing task ", t.ID)

	time.Sleep(2 * time.Second)
}
