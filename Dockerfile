FROM golang:latest AS build-stage
# ENV DOCKERIZE_VERSION v0.6.0
# RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
#     && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
#     && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# CMD dockerize -wait tcp://mysql-development:3308 -timeout 720s
LABEL app="capital-adequacy"
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Because of https://github.com/docker/docker/issues/14914
# required by go get
# RUN apk add --update --no-cache alpine-sdk bash python ca-certificates \
    #   libressl \
    #   tar \
    #   git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev coreutils

WORKDIR $GOPATH/src/airflow-report/capital-adequacy

COPY . .

# build the application
# RUN chmod +x init.sh
# RUN ./init.sh
RUN go test ./...
RUN GOOS=linux go build -a -o capital-adequacy .

