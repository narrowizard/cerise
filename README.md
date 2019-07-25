# Cerise
Cerise is a task runner based on time.Ticker. extensible and flexible.

## get started
at first, write a config file to tell cerise what to do. eg:  
```yaml
tasks:
  # task name
  task1:
    # task type, specify to task runner
    type: "test-runner"
    # tick interval
    interval: 1000
    # if stop ticker when runner cause error
    stopOnError: true
    # props for spec task runner
    props:
      key: "hahhaha"
```

## add extension
1. implement your own task runner.  
1. register your task runner via `RunnerContainer#Register`
1. define your task in `config.yaml`
1. and everything is down.
