# QUICKSTART A GO CLI APP

## Install Go
`sudo apt install golang.go`

## Set Go Path
`export GOPATH="~/go"`

## Create Project Directory
`mkdir <your-directory-name>`

## Initialize Git
```
cd <your-directory-name>
git init
```

## Initialize GO Module
`go mod init github.com/<user-name>/<app_name>`

## Install Cobra
`go get -u github.com/spf13/cobra`
More? Maybe need cobra-cli?

## Initialize Cobra
`cobra-cli init`

# Running a Go Application

Can either compile and run the application directly via `go run main.go` or compile with `go build` and run the executable with `./<app-name>`.