FROM golang:alpine AS builder

ENV GO111MODULE=on

# alpine image does not have git in it
RUN apk update && apk add --no-cache git

WORKDIR /app

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch

COPY --from=builder /app/bin/main .

CMD ["./main"]