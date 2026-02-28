FROM golang:1.26-alpine AS devtool

RUN apk add --no-cache \
    git \
    curl \
    bash \
    make \
    gcc \
    musl-dev

WORKDIR /app

COPY . .

# Keep the container running for development purposes
CMD ["sleep", "infinity"]
