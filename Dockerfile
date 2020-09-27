
FROM golang:1.14.9-stretch AS builder
COPY . /go/src/github.com/le-garden-fox/healthchecker
WORKDIR /go/src/github.com/le-garden-fox/healthchecker
RUN  go get
# build for alpine 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o healthchecker main.go

#small image 
FROM scratch
COPY --from=builder /go/src/github.com/le-garden-fox/healthchecker/healthchecker .
EXPOSE 8080
ENTRYPOINT ["./healthchecker"]