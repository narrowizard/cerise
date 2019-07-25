package runner

// APITrigger handle api-trigger task
type APITrigger struct{}

// NewAPITrigger create api trigger
func NewAPITrigger() *APITrigger {
	return &APITrigger{}
}

// Run start task
func (at *APITrigger) Run(props map[string]interface{}) error {
	return nil
}
