FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

EXPOSE 8000


RUN go build -o /items-api

CMD ["/items-api"]