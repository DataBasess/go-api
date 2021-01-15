# init project
go mod init go-api

# install
go mod download

# add dependencie
go get -u golang.org/x/lint/golint
go mod init github.com/callicoder/go-docker

# run project
go run main.go

# build project
go build