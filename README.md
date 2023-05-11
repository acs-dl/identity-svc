# identity-svc

## Description

Identity service for access management service

## Install

  ```bash
  git clone github.com/acs-dl/identity-svc
  cd identity-svc
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main migrate up
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```bash
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.


## Running from docker 
  
Make sure that docker installed.
{%_ if (handleHTTP) { _%}
use `docker run ` with `-p 8080:80` to expose port 80 to 8080

{%_ } _%}

    ```bash
    docker build -t github.com/acs-dl/identity-svc .
    docker run -e KV_VIPER_FILE=/config.yaml github.com/acs-dl/identity-svc
    ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `migrate up` command to create database schema
* Launch the service with `run service` command


### Database
For services, we do use ***PostgresSQL*** database. 
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).


### Third-party services


## Contact

Responsible Dima Kudriavtsev
The primary contact for this project is  @slandyman
