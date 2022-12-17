FROM ubuntu:22.04

RUN apt-get update && apt-get install avahi-utils -y

COPY --from=golang:1.19-bullseye /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=1 go build -o /mdns-docker

CMD [ "/mdns-docker" ]