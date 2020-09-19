FROM golang:latest

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN chmod +x ./deploy-golang-app-in-heroku

EXPOSE 8080

CMD ./deploy-golag-app-in-heroku