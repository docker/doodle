FROM --platform=$BUILDPLATFORM golang:1.11-alpine AS builder
RUN apk add --no-cache git
RUN go get github.com/pdevine/go-asciisprite
WORKDIR /project
COPY cheers.go .

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o cheers cheers.go

FROM scratch AS release-linux
COPY --from=builder /project/cheers /cheers
ENTRYPOINT ["/cheers"]

FROM mcr.microsoft.com/windows/nanoserver:1809 AS release-windows
COPY --from=builder /project/cheers /cheers.exe
ENTRYPOINT ["\\cheers.exe"]

FROM release-$TARGETOS
