FROM golang:1.15 AS builder
WORKDIR /build
COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod ./
RUN go get -d -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app /build/cmd/mpg/main.go

FROM alpine:3.13

RUN apk update && apk upgrade && apk add --no-cache bash git ca-certificates chromium

# Installs latest Chromium package.
RUN echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories \
    && echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories \
    && apk add --no-cache \
    harfbuzz@edge \
    nss@edge \
    freetype@edge \
    ttf-freefont@edge \
    && rm -rf /var/cache/* \
    && mkdir /var/cache/apk

WORKDIR /app/
COPY --from=builder /build/app .
COPY templates ./

RUN addgroup -S app && adduser -S -G app app
RUN chown -R app:app /app
USER app

CMD ./app