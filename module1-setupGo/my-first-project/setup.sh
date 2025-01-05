#!bin/bash
git init

go mod init github.com/soypetetech/my-first-go-project

go mod tidy

nvim main.go
