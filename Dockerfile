FROM --platform=linux/amd64 golang:1.20-alpine
MAINTAINER OD.Ice
WORKDIR /go/src/survey_backend
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]
