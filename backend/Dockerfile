# Development image - live reloading with CompileDaemon
 
FROM golang:alpine3.19
# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0
RUN mkdir /app
COPY .. /app
WORKDIR /app
# Get CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
# The build flag sets how to build after a change has been detected in the source code
# The command flag sets how to run the app after it has been built
ENTRYPOINT CompileDaemon -build="go build -o api" -command="./api"