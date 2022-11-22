FROM golang:1.19.3-alpine3.16

ENV \
    CGO_ENABLED=0 \
    GOOS=linux

RUN apk add postgresql-client

RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
COPY main.go /app
COPY sql /app
COPY static /app
COPY templates /app
COPY .env_db /app
COPY .env_app /app

RUN chown -R appuser:appgroup /app
RUN chown -R appuser:appgroup /go/bin
USER 1000

RUN go mod download 

RUN echo ${PGPASSFILE} > /home/appuser/.pgpass && \
    chown appuser:appgroup /home/appuser/.pgpass && \
    chmod 0600 /home/appuser/.pgpass

# Build the Go app
RUN go build -ldflags "-s -w" -o /go/bin/cycling_blog -buildvcs=false

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["cycling_blog"]
