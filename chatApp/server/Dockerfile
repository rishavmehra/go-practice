# Build stage
FROM golang:1.22.4-alpine AS build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
RUN chmod +x /app/main

# Final stage
FROM alpine:latest
COPY --from=build /app/main /main
ENTRYPOINT ["/main"]


# run back as docker container: docker run -p 8080:8080 backend
