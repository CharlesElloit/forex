FROM golang:1.19-buster AS build

WORKDIR /app

RUN useradd -u 1001 noneroot

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags metgo \
    -o api.forex.com

###
FROM scratch

ENV key=value

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/api.forex.com /api.forex.com

USER noneroot

EXPOSE 4000

CMD [ "/api.forex.com" ]

