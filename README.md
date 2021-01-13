go run main.go
go build

go mod init github.com/callicoder/go-docker
go mod download
go get -u golang.org/x/lint/golint