FROM registry.semaphoreci.com/golang:1.18 as builder

WORKDIR /myApp
COPY ./ .

RUN go mod download
RUN go build -o serverRunner main.go

FROM registry.semaphoreci.com/golang:1.18

WORKDIR /
COPY --from=builder /myApp/serverRunner .
USER root

ENTRYPOINT ["/serverRunner"]

