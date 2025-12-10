FROM golang:1.24-alpine as build

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o /graphql-service .


FROM alpine:latest  

WORKDIR /root/
COPY --from=build /graphql-service .

EXPOSE 8080
CMD ["./graphql-service"]
