FROM --platform=$BUILDPLATFORM golang:1.11-alpine AS builder
RUN apk add --no-cache git
RUN go get github.com/pdevine/go-asciisprite
WORKDIR /project
COPY summer.go .

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o summer summer.go

FROM scratch AS release-linux
COPY --from=builder /project/summer /summer
ENTRYPOINT ["/summer"]

FROM mcr.microsoft.com/windows/nanoserver:1809 AS release-windows
COPY --from=builder /project/summer /summer.exe
ENTRYPOINT ["\\summer.exe"]

FROM release-$TARGETOS
