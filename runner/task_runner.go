package runner

import (
	"errors"
	"fmt"
	"time"

	"github.com/narrowizard/cerise/models"
)

// RunnerContainer global runner and task container
type RunnerContainer struct {
	runners map[string]models.Runner
	tasks   map[string]chan bool
}

// NewRunnerContainer create runner container
func NewRunnerContainer() *RunnerContainer {
	return &RunnerContainer{
		runners: make(map[string]models.Runner),
		tasks:   make(map[string]chan bool),
	}
}

// RegisterRunner register a task runner
func (rc *RunnerContainer) RegisterRunner(name string, r models.Runner) {
	rc.runners[name] = r
}

// RunTask start a task
func (rc *RunnerContainer) RunTask(name string, t models.Task) error {
	var runner, ok = rc.runners[t.Type]
	if !ok {
		return errors.New("not found specific task runner")
	}
	_, ok = rc.tasks[name]
	if ok {
		return errors.New("task already exists, task name conflict")
	}
	// start a ticker
	var ticker = time.NewTicker(time.Duration(t.Interval) * time.Millisecond)
	var stop chan bool
	go func() {
		for range ticker.C {
			select {
			case <-stop:
				fmt.Println("task " + name + " stopped")
				ticker.Stop()
			default:
				// run task
				runner.Run(t.Props)
			}
		}
	}()
	rc.tasks[name] = stop
	return nil
}

// StopTask manually stop a task
func (rc *RunnerContainer) StopTask(name string) error {
	var stop, ok = rc.tasks[name]
	if !ok {
		return errors.New("task not found")
	}
	stop <- true
	delete(rc.tasks, name)
	return nil
}
