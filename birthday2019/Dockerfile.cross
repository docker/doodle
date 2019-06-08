FROM --platform=$BUILDPLATFORM golang:1.11-alpine AS builder
RUN apk add --no-cache git
RUN go get github.com/pdevine/go-asciisprite
WORKDIR /project
COPY surprise.go .

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o surprise surprise.go

FROM scratch AS release-linux
COPY --from=builder /project/surprise /surprise
ENTRYPOINT ["/surprise"]

FROM mcr.microsoft.com/windows/nanoserver:1809 AS release-windows
COPY --from=builder /project/surprise /surprise.exe
ENTRYPOINT ["\\surprise.exe"]

FROM release-$TARGETOS
