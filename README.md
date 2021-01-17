# init project
go mod init go-api

# install
go mod download

# add dependencie
go mod init github.com/callicoder/go-docker
go get -u github.com/codegangsta/gin
go get -u golang.org/x/lint/golint
go get -u github.com/labstack/echo/v4
go get -u github.com/valyala/fasthttp


# run project
go run main.go
gin run main.go

# build project
go build

docker run --rm -d --name nicehash-miner -e BTC_PAYMENT_ADDRESS=3Khe6ZchukMDERASSRMc6pJuWnWbQMwDsM caspernaudts/nicehash-miner:1.0