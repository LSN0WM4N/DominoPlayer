FROM golang:1.18-alpine 

ADD ./app /app 

CMD go run app/main.go