# ----------------------- Stage building -----------------------
FROM golang:alpine AS builder
ENV LANG=C.UTF-8

WORKDIR /app

# Install packages
RUN apk update &&\
    apk add --update --no-cache bash curl \
    gcc g++ vips vips-dev vips-poppler pkgconf \
    poppler-dev musl-dev libc6-compat

# Install dependencies and build the app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Build the app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./.build/main ./cmd/main.go



# ----------------------- Stage serve -----------------------
FROM alpine:latest AS cdn
ENV LANG=C.UTF-8

WORKDIR /app

# Install packages
RUN apk update &&\
    apk add --update --no-cache bash \
    gcc g++ vips vips-dev vips-poppler pkgconf \
    poppler-dev musl-dev libc6-compat

# Copy the binary and assets to the container
COPY --from=builder /app/.build/ ./
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/.env.example ./app.env
RUN chmod +x /app/main

EXPOSE 3100

ENTRYPOINT ["/app/main"]
