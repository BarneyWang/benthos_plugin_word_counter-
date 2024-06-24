FROM golang:1.22 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN useradd -u 10001 benthos

WORKDIR /build/
COPY . /build/

RUN go mod tidy && go mod vendor
RUN go build -mod=vendor

FROM busybox AS package

LABEL maintainer="di.wang <di.wang@justeattakeaway.com>"

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /build/benthos_plugin_word_counter .

COPY ./config_counter.yaml /benthos.yaml

USER benthos

EXPOSE 4195

ENTRYPOINT ["/benthos_plugin_word_counter"]

CMD ["-c", "/benthos.yaml"]
