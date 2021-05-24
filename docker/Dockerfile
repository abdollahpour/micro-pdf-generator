ARG GOLANG_VERSION=1.16
FROM golang:${GOLANG_VERSION} AS builder
ARG APP_VERSION=developement
WORKDIR /build
COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod ./
RUN go get -d -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.Version=${APP_VERSION}" -a -installsuffix cgo -o app /build/cmd/server/main.go

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

# TODO: Run the app on non-root use and make sure chrome has access to temp files
# RUN addgroup -S app && adduser -S -G app app
# RUN chown -R app:app /app
# RUN chmod -R +r /tmp
# USER app

ENTRYPOINT [ "./app" ]