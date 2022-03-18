FROM golang:1.9.3-alpine3.7 AS build
RUN apk add -U g++ make
COPY . /go/src/github.com/hinshun/genplugin
RUN go install github.com/hinshun/genplugin/cmd/genplugin

FROM alpine:3.15
COPY --from=build /go/bin/genplugin bin/
WORKDIR /opt/plugin
COPY plugin /opt/plugin
ENTRYPOINT ["genplugin"]
CMD ["--help"]
