FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app/

RUN go build -o desafio

# EXPOSE 8080

# CMD [ "/app/desafio" ]

FROM golang:1.16-alpine

WORKDIR /

COPY --from=build /app/desafio /desafio

EXPOSE 8080

ENTRYPOINT ["/desafio"]