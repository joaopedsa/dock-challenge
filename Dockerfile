FROM golang:1.18

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
RUN go build -o ./app

CMD ["./app"]