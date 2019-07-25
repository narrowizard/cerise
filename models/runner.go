package models

// Runner represent a task runner
type Runner interface {
	Run(props map[string]interface{}) error
}
