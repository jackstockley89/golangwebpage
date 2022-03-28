FROM golang:1.17.8-alpine3.15

RUN apk add git
RUN apk add postgresql-client

RUN go get -u github.com/lib/pq
RUN go get -u github.com/joho/godotenv

RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

WORKDIR /app

COPY . . 

RUN echo ${PGPASSFILE} > /home/appuser/.pgpass && \
    chown appuser:appgroup /home/appuser/.pgpass && \
    chmod 0600 /home/appuser/.pgpass
 
RUN chown -R appuser:appgroup /app
RUN chown -R appuser:appgroup /go

USER 1000

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
