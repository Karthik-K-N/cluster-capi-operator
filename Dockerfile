ARG GOLANG_IMAGE=golang:1.18.5
ARG TARGETPLATFORM=linux/amd64
ARG ARCH=amd64

FROM ${GOLANG_IMAGE} AS builder
ARG ARCH
WORKDIR /go/src/github.com/openshift/cluster-capi-operator
COPY . .
RUN CGO_ENABLED=0 GOARCH=$ARCH go build -o bin/cluster-capi-operator cmd/cluster-capi-operator/main.go

FROM --platform=$TARGETPLATFORM quay.io/centos/centos:stream8 as centos-base
COPY --from=builder /go/src/github.com/openshift/cluster-capi-operator/bin/cluster-capi-operator .
COPY --from=builder /go/src/github.com/openshift/cluster-capi-operator/manifests /manifests

LABEL io.openshift.release.operator true

