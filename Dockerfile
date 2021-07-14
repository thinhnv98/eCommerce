FROM golang:latest as builder
LABEL stage=builder

WORKDIR /app
COPY env/prod.env ./env/.env
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

FROM golang:latest as final
LABEL stage=Final

WORKDIR /root

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/env/.env ./env/.env

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]