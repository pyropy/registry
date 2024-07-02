FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/app/main.go

# Deploy the application binary into a lean image
FROM alpine:3.20 AS runner

WORKDIR /

COPY --from=build /api /api

EXPOSE 8000

ENTRYPOINT ["/api"]