FROM golang:1.19-bullseye AS build_base

WORKDIR /tmp/gateway

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/gateway ./cmd/gateway


FROM debian:bullseye

RUN apt-get update && apt-get install -yq curl

COPY --from=build_base /tmp/gateway/out/gateway /app/gateway

ENV GODEBUG madvdontneed=1

CMD ["/app/gateway"]
