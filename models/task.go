package models

// Task represent a auto-invoke task
type Task struct {
	Type     string                 // api-trigger
	Interval int64                  // ms
	Props    map[string]interface{} // props
}

// APITriggerProps api trigger task props model
type APITriggerProps struct {
	Protocol string
	Host     string
	Pathname string
	Method   string
	Params   string
}
