FROM golang:alpine AS build

WORKDIR /go/src/app



COPY . .

RUN go build -o /go/bin/app

FROM alpine:latest

COPY --from=build /go/bin/app /app

EXPOSE 8080

CMD ["/app"]
