FROM golang:1.14-alpine as builder

# ENV GO111MODULE=on

LABEL maintainer="Anton Kovalenko <CaribbeanLegend@mail.ru>"

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./main ./

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main /app/.env ./

ENTRYPOINT ["./main"]