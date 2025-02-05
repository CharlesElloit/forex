FROM golang:1.23-bullseye AS build-base

WORKDIR /app 

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

FROM build-base AS dev

RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

CMD ["air", "-c", ".air.toml"]

FROM build-base AS build-production

RUN useradd -u 1001 nonroot

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o healthcheck \
    ./healthcheck/healthcheck.go

RUN echo "Hello"

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o api.forex

FROM scratch

# ENV GIN_MODE=release

WORKDIR /

COPY --from=build-production /etc/passwd /etc/passwd

COPY --from=build-production /app/healthcheck/healthcheck healthcheck

COPY --from=build-production /app/api.forex api.forex

USER nonroot

EXPOSE 4000

CMD ["/api.forex"]