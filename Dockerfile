FROM golang:1.5-alpine:3.12 AS gobuilder-stage

MAINTAINER 
LABEL "purpose"="Service Deployment using Multi-stage builds."

WORKDIR /usr/src/goapp

COPY goapp.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/gostart

FROM scratch AS runtime-stage

COPY --from=gobuilder-stage /usr/local/bin/gostart /usr/local/bin/gostart

CM ["/usr/local/bin/gostart"]
