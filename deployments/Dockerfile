FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache git ca-certificates make gcc musl-dev

COPY go.mod go.sum ./
COPY api/go.mod api/go.sum ./api/
RUN go mod download

COPY . .

RUN make build
RUN make test

FROM alpine:3.19 AS runtime

WORKDIR /app

RUN apk add --no-cache tzdata
RUN addgroup -S nonroot
RUN adduser -S nonroot -G nonroot

COPY --from=build /app/bin/inventory /app/inventory
COPY configs /app/configs

USER nonroot:nonroot

ENTRYPOINT ["/app/inventory"]
