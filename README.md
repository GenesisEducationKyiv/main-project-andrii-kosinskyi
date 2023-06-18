# bitcoin_checker_api

This repo consists of
_env
cmd
    main.go
config
    config.go
internal
    handlers
    models
    repositories

In _env folder we have configuration.
In config folder we have model for config
In handlers folder we have handler and utils
In models we have model for user and convertor
In repositories we have interface for internal-storage

In this project I used https://gin-gonic.com/
Information about bitcoin I got from  

For check on duplicate user record I used map, but I think better it is binary search.

Our entry point in this repo is cmd/main.go in this file we initialiase config, repository an handlers after we use 
framework GIN for routing