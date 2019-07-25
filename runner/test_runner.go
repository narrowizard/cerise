package runner

import "fmt"

type TestRunner struct{}

func NewTestRunner() *TestRunner {
	return &TestRunner{}
}

func (tr *TestRunner) Run(props map[string]interface{}) error {
	for k, v := range props {
		fmt.Println(k)
		fmt.Println(v)
	}
	return nil
}
