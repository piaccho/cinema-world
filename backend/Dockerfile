FROM golang:alpine3.19 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o app

FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]