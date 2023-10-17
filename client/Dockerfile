FROM registry.semaphoreci.com/golang:1.18 as builder

WORKDIR /myApp
COPY ./ .

RUN go mod download
RUN go build -o clientRunner main.go

FROM registry.semaphoreci.com/golang:1.18

WORKDIR /
COPY --from=builder /myApp/clientRunner .
USER root

ENTRYPOINT ["/clientRunner"]
