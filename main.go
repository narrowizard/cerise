package main

import (
	"fmt"

	"github.com/narrowizard/cerise/config"
	"github.com/narrowizard/cerise/runner"
)

func main() {
	var tasks, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}
	var rc = runner.NewRunnerContainer()
	rc.RegisterRunner("api-trigger", runner.NewAPITrigger())
	rc.RegisterRunner("test-runner", runner.NewTestRunner())

	for k, v := range tasks {
		err = rc.RunTask(k, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	// broke main thread
	select {}
}
