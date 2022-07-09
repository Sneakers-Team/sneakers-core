FROM golang

EXPOSE 8080

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD [ "go", "run", "cmd/main.go" ]