# JC Interview Assignment

This program requires a PORT to listen on. It can be passed as a command line argument or set as an environment variable. The command line argument will take precidence. Here are some usage examples:

    go run main.go routes.go 8080

    export JC_INTERVIEW_ASSIGNMENT_PORT=8080
    go run main.go routes.go

## Running The Program

You can use go run with the following command `go run main.go routes.go 8080`.

You can also build it first, `go build`, then `./jc-interview-assignment 8080`.
