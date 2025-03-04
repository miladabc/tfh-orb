FROM golang:1.23.5-alpine3.21 AS build

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go mod download

COPY . .

RUN go build \
	-ldflags="-linkmode external -extldflags -static" \
	-tags netgo \
	-o orb

FROM alpine:3.21

RUN apk update && apk upgrade --no-cache && apk add --update --no-cache tzdata ca-certificates

ENV TZ="Asia/Tehran" PATH="/app:${PATH}"

RUN adduser -S -u 1001 guser
RUN mkdir /app && chown guser /app
WORKDIR /app
USER guser

COPY --from=build --chown=guser /app/orb ./orb

ENTRYPOINT ["orb"]
