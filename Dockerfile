# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY cni/ cni/
COPY controllers/ controllers/
COPY node/ node/
COPY polycube/ polycube/
COPY types/ types/
COPY utils/ utils/

WORKDIR /workspace/cni
# Build polycube-cni-plugin
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o plugins/polykube-cni-plugin
WORKDIR /workspace
# Build polykube
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o polykube main.go

# Use distroless as minimal base image to package the manager binary and the CNI plugin
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static
FROM alpine:3.14
RUN apk add --no-cache bash iproute2

RUN mkdir -p /host/var/log /host/opt/cni/bin /host/etc/cni/net.d/

WORKDIR /cni-plugins

# Copy CNI plugins in /cni-plugins
COPY --from=builder /workspace/cni/plugins/* ./

WORKDIR /
# Copy CNI installation and uninstallation scripts and manager binary in the root folder
COPY --from=builder /workspace/cni/*.sh workspace/polykube ./

CMD ["/polykube"]
