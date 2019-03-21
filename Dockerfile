FROM golang:alpine3.8
WORKDIR /project
COPY surprise.go .
RUN apk add --no-cache git
RUN go get github.com/nsf/termbox-go && go get github.com/pdevine/go-asciisprite
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o surprise surprise.go

FROM scratch
COPY --from=0 /project/surprise /surprise
CMD ["/surprise"]
