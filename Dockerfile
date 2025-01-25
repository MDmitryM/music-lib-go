FROM golang:latest AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

RUN go mod download; CGO_ENABLED=0 go build -o /music-lib-go ./cmd/main.go

FROM alpine:3.17

COPY --from=build /music-lib-go /music-lib-go
COPY ./configs/ /configs/

CMD ["./music-lib-go"]