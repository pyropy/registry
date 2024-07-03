FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app cmd/app/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/deploy cmd/deploy/main.go

FROM alpine:3.20 AS deploy-runner

WORKDIR /

COPY --from=build /app/deploy /deploy

ENTRYPOINT ["/deploy"]

FROM alpine:3.20 AS app-runner

WORKDIR /

COPY --from=build /app/app /app

EXPOSE 8000

ENTRYPOINT ["/app"]