FROM golang:1.11.0 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch 
COPY --from=builder /build/main /confucius/
WORKDIR /confucius
CMD ["./main"]
