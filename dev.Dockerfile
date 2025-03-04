FROM golang:1.23.5-alpine3.21

WORKDIR /app

ENV TZ="Asia/Tehran" PATH="/app:${PATH}"

RUN go install github.com/bokwoon95/wgo@v0.5.7
RUN go install gotest.tools/gotestsum@v1.12.0

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go mod download

COPY . .

ENTRYPOINT ["wgo", "go", "run", "."]
