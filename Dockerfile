# build stage
FROM golang:1.16-alpine as build

ADD . /go/src/github.com/dihr/app
WORKDIR /go/src/github.com/dihr/app
RUN go mod download \
    && go mod tidy
RUN CGO_ENABLED=0 go build -a -installsuffix main.go -o main

# final stage
FROM alpine
WORKDIR /app
ENV BOT_ID="951642055220805633"
ENV BOT_TOKEN="OTUxNjQyMDU1MjIwODA1NjMz.Yiqbug.x4bV5QtPuelrgYtOB8pTiOETIsY"
COPY --from=build /go/src/github.com/dihr/app/main /app/main
ENTRYPOINT ./main
