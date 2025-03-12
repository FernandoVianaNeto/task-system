# build image
FROM public.ecr.aws/docker/library/golang:1.23-alpine3.19 AS builder

RUN apk add --update \
    git \
    build-base \
    coreutils \
    gcc \
    make \
    musl-dev \
    ca-certificates \
    && mkdir /go/src/app \
    && git config --global --add safe.directory /usr/src/app

WORKDIR /usr/src/app

COPY go.* ./
RUN go mod tidy

COPY . .

RUN make build

# final image
FROM public.ecr.aws/docker/library/alpine:3.17
ENV TZ="America/Sao_Paulo"

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ca-certificates.conf /etc/ca-certificates.conf
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/ca-certificates /usr/share/ca-certificates

COPY --from=builder /usr/src/app/cmd/config /configs

COPY --from=builder /usr/src/app/bin/application /bin/application

CMD ["/bin/application"]