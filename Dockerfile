FROM golang:1.11.0

WORKDIR /go/src/github.com/igorpo/confusius
COPY . .

RUN go install .
CMD "confusius"
