FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

COPY api/ ./api
COPY model/ ./model

RUN CGO_ENABLED=0 GOOS=linux go build -o /dynamic-dns-service

CMD [ "/dynamic-dns-service" ]
