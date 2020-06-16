FROM golang:1.14-alpine as build-env

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

RUN mkdir /work

WORKDIR /work
COPY ./go.mod .
COPY ./go.sum .
ENV GOPROXY https://goproxy.io
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/overtime-backend

VOLUME ["config", "log"]

FROM alpine
COPY --from=build-env /go/bin/overtime-backend /go/bin/overtime-backend
ENTRYPOINT ["go/bin/overtime-backend"]
CMD ["--config-path=/config"]