# Dockerfile
FROM alpine:latest

# Install netcat-openbsd, curl and Go
RUN apk update && \
    apk add --no-cache \
    netcat-openbsd \
    curl \
    go

# Install go dependencies
RUN go install github.com/bootdotdev/bootdev@latest

# Add go binary to PATH
ENV PATH="/root/go/bin:${PATH}"

WORKDIR /app

# Copy source code
COPY ./src .

# Copy bootdev login info
COPY ./.bootdev/* /root
