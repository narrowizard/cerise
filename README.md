# Cerise
[![Build Status](https://travis-ci.org/narrowizard/cerise.svg?branch=master)](https://travis-ci.org/narrowizard/cerise)  
Cerise is a task runner based on time.Ticker. extensible and flexible.

## get started
at first, write a config file to tell cerise what to do. eg:  
```yaml
tasks:
  # task name
  task1:
    # task type, specify to task runner
    type: "test-runner"
    # specify first task starts time
    startAt: "2019-10-09 16:00:00"
    # tick interval
    interval: 1000
    # if stop ticker when runner cause error
    stopOnError: true
    # props for spec task runner
    props:
      key: "hahhaha"
```

## docker
```yaml
# docker-compose.yml
version: "3"
services:
  cerise:
    image: "cerise"
    volumes:
     # volume your config.yaml
     - /home/data/cerise/config:/root/config
    environment:
     # set timezone
     - TIMEZONE=Africa/Abidjan
```


## add extension
1. implement your own task runner.  
1. register your task runner via `RunnerContainer#Register`
1. define your task in `config.yaml`
1. and everything is down.
