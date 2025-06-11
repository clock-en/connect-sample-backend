FROM golang:1.24-bullseye

RUN apt-get update && \
    apt-get install -y git bash make && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/air-verse/air@v1.61.7

WORKDIR /app

# Install dependencies
COPY go.mod go.sum .air.toml ./
RUN go mod download

# Copy source code
COPY . .

CMD ["air"]
