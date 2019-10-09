package models

// Task represent a auto-invoke task
type Task struct {
	Type        string // api-trigger
	StartAt     string // first trigger time, yyyy-MM-dd hh:mm:ss
	Interval    int64  // ms
	StopOnError bool
	Props       map[string]interface{} // props
}

// APITriggerProps api trigger task props model
type APITriggerProps struct {
	Protocol string
	Host     string
	Pathname string
	Method   string
	Params   string
}
