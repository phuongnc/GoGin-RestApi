ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine AS build_base
LABEL stage=build_base
RUN apk update && apk add gcc libc-dev make git --no-cache ca-certificates  && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /go/src/shapeservice
WORKDIR /go/src/shapeservice

ENV GO111MODULE=on
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# ================ copy from stage build ===========

FROM build_base AS server_builder
LABEL stage=server_builder

RUN mkdir -p /go/src/shapeservice/cmd
COPY ./cmd/entity-server /go/src/shapeservice/cmd/entity-server

RUN mkdir -p /go/src/shapeservice/internal
COPY ./internal/pkg /go/src/shapeservice/internal/pkg

RUN mkdir -p /go/src/shapeservice/configs
COPY ./configs /go/src/shapeservice/configs

RUN mkdir -p /go/src/shapeservice/storage
COPY ./storage /go/src/shapeservice/storage

WORKDIR /go/src/shapeservice/cmd/entity-server/

RUN go build  -o /entity-server .

# ================ copy from stage build ===========
FROM alpine:3.8

RUN apk update &&  apk add --no-cache ca-certificates git && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /shapeservice
WORKDIR /shapeservice

RUN mkdir -p /shapeservice/cmd/bin
RUN mkdir -p /shapeservice/configs
RUN mkdir -p /shapeservice/storage

COPY --from=server_builder /entity-server /shapeservice/cmd/bin/
COPY --from=server_builder /go/src/shapeservice/configs/config.toml /shapeservice/configs/
COPY --from=server_builder /go/src/shapeservice/storage /shapeservice/storage/

RUN chmod -R 777 /shapeservice/cmd/bin

RUN chown -R nobody:nobody /shapeservice
RUN chmod -R 755 /shapeservice

USER nobody:nobody

EXPOSE 7171
ENTRYPOINT ["/shapeservice/cmd/bin/entity-server"]