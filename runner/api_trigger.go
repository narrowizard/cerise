package runner

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/narrowizard/cerise/models"
)

// APITrigger handle api-trigger task
type APITrigger struct{}

// NewAPITrigger create api trigger
func NewAPITrigger() *APITrigger {
	return &APITrigger{}
}

// Run start task
func (at *APITrigger) Run(props map[string]interface{}) error {
	var m models.APITriggerProps
	var err = mapstructure.Decode(props, &m)
	if err != nil {
		return err
	}
	var client = &http.Client{}
	req, err := http.NewRequest(m.Method, m.Protocol+"//"+m.Host+m.Pathname, strings.NewReader(m.Params))
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode <= 300 {
		return nil
	}
	// 出错
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return fmt.Errorf("http error %d", res.StatusCode)
}
