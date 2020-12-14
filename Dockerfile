FROM golang:1.14.3-alpine3.11

LABEL "Version"="1.1"

RUN apk add git
RUN apk add postgresql-client

RUN go get -u github.com/lib/pq
RUN go get -u github.com/joho/godotenv

RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup



WORKDIR /app

COPY . . 
 
RUN chown -R appuser:appgroup /app

USER 1000

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
