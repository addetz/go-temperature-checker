FROM golang:1.21.0-alpine3.17 as build

ENV GOPATH /go

WORKDIR $GOPATH/src/app

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -ldflags "-s -w" -o server

RUN chmod +x server

FROM alpine:3.17

COPY --from=build go/src/app/layout /layout
COPY --from=build go/src/app/scripts /scripts
COPY --from=build go/src/app/apis /apis
COPY --from=build go/src/app/server /

EXPOSE 8080

ENTRYPOINT [ "./server" ]