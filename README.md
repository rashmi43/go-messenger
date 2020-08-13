# go-messenger

## Architecture

## How to clone and run

git clone https://github.com/rashmi43/go-messenger.git

cd go-messenger/

go get -v -t -d ./...

go build ./...

go run main.go

## How to build
Using Dockerfile:

docker build -t "go-messenger" .
docker run --publish 8001:8001 --name test go-messenger:latest

Using CLI:
cd github.com/go-messenger
go run main.go

## How to access
Use HTTP protocol and port 8001 to access the application, type in your browser
http://ip:8081/v1/messages

## REST API
### Create a message
curl -X POST 'http://localhost:8001/v1/messages' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ID": "2",
    "text": "My favorite passtime is tweeting",
    "submitter": "rashmi"
}'

### GET All Messages
curl --location --request GET 'http://localhost:8001/v1/messages' \
--header 'Content-Type: application/json'

### GET  message by id
curl --location --request GET 'http://localhost:8001/v1/messages/2' \
--header 'Content-Type: application/json'

