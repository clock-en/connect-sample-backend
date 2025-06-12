FROM golang:1.24-bullseye

RUN apt-get update && \
    apt-get install -y git bash make && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/air-verse/air@v1.61.7 && \
    go install github.com/bufbuild/buf/cmd/buf@v1.54.0 && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.9.3 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6 && \
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.18.1

WORKDIR /app

# Install dependencies
COPY go.mod go.sum .air.toml ./
RUN go mod download

# Copy source code
COPY . .

CMD ["air"]
