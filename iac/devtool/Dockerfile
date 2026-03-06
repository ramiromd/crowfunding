FROM golang:1.26-alpine AS devtool

RUN apk add --no-cache \
    git \
    curl \
    bash \
    make \
    gcc \
    lcov \
    musl-dev

WORKDIR /app

RUN go install github.com/jandelgado/gcov2lcov@v1.1.1

COPY . .

EXPOSE 8080
# Keep the container running for development purposes
CMD ["sleep", "infinity"]
