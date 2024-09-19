FROM golang:1.23-alpine AS build

WORKDIR /TimeBookingAPI

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o app ./cmd/time-booking

FROM alpine:latest
WORKDIR /root/

COPY --from=build /TimeBookingAPI/app /TimeBookingAPI/app
COPY ./config/config.yaml ./config/config.yaml
COPY ./logs ./logs

EXPOSE 8080

CMD ["/TimeBookingAPI/app"]
