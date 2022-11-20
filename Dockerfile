FROM golang:1.19-bullseye AS build_base

WORKDIR /tmp/friendly-spoon

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/friendly-spoon ./cmd/friendly-spoon


FROM debian:bullseye

RUN apt-get update && apt-get install -yq curl

COPY --from=build_base /tmp/friendly-spoon/out/friendly-spoon /app/friendly-spoon

COPY --from=build_base /tmp/friendly-spoon/sql /sql

ENV GODEBUG madvdontneed=1

CMD ["/app/friendly-spoon"]
