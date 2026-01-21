FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache git ca-certificates make

COPY go.mod go.sum ./
COPY api/go.mod api/go.sum ./api/
RUN go mod download

COPY . .

RUN make build
RUN make test

FROM gcr.io/distroless/base-debian12 AS runtime

WORKDIR /app

COPY --from=build /app/bin/inventory /app/inventory
COPY configs /app/configs

USER nonroot:nonroot

ENTRYPOINT ["/app/inventory"]
