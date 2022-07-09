FROM golang

EXPOSE 8080

ENV USERSERVICE_URL = sneakers.userservice 

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD [ "go", "run", "cmd/main.go" ]