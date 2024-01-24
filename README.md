# project.template.go

Template for Go backend projects

## Install dependencies

```shell
$ make install
```

## Run the project

Start the database container:

```shell
$ make db-start
```

Run application:

```shell
$ make start
```


## Run test
```shell
$ make test
```

## Features
- **Gin Web Framework**: Utilizes the Gin framework, a fast and lightweight web framework for Go, to handle routing, middleware, and more.

- **Dependency Injection**: Demonstrates a simple setup for dependency injection to manage dependencies throughout your application.

- **Error Handling**: Includes basic error handling patterns and demonstrates how to handle errors gracefully in your application.

- **Configuration**: Provides a basic configuration setup to manage environment-specific configurations.

- **Logging**: Illustrates logging implementation using a logging library compatible with Gin.

- **Testing**: Includes a basic structure for writing tests, making it easier to ensure the reliability of your code.
