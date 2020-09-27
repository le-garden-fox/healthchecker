#!/usr/bin/env bash

function main() {

    case $1 in
    "docker-build")
        docker build -t le-garden-fox/healthchecker -f Dockerfile .
        ;;

    "docker-run")
        docker run -d -p 8080:8080 le-garden-fox/healthchecker 
        ;;
    "run")
        go run main.go
        ;;

    "build")
        go build -o healthchecker main.go
        ;;

    *)
        echo "Command $1 not found"
        ;;

    esac
}

main "$@"
