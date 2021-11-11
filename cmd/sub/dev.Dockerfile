FROM golang:1.17-alpine AS go-builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add --no-cache upx && \
    go version && \
    go mod download
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/k-yomo/pubsub_cli
COPY . .

CMD cd cmd/sub && air -c .air.toml
