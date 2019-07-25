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
	fmt.Println("task " + name + " started")
	go func() {
		// manually tick at start
		var err = runner.Run(t.Props)
		if err != nil && t.StopOnError {
			rc.StopTask(name)
		}
		for range ticker.C {
			select {
			case <-stop:
				ticker.Stop()
			default:
				// run task
				var err = runner.Run(t.Props)
				if err != nil && t.StopOnError {
					rc.StopTask(name)
				}
			}
		}
	}()

	rc.tasks[name] = stop
	return nil
}

// StopTask manually stop a task
func (rc *RunnerContainer) StopTask(name string) error {
	fmt.Println("task " + name + " stopped")
	var stop, ok = rc.tasks[name]
	if !ok {
		return errors.New("task not found")
	}
	stop <- true
	delete(rc.tasks, name)
	return nil
}
