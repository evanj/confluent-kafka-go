//go:build !dynamic && musl
// +build !dynamic,musl

// This file was auto-generated by librdkafka_vendor/bundle-import.sh, DO NOT EDIT.

package kafka

// #cgo CFLAGS: -DUSE_VENDORED_LIBRDKAFKA -DLIBRDKAFKA_STATICLIB
// #cgo LDFLAGS: ${SRCDIR}/librdkafka_vendor/librdkafka_musl_linux_arm64.a  -lm -ldl -lpthread -lrt -lpthread -lrt
import "C"

// LibrdkafkaLinkInfo explains how librdkafka was linked to the Go client
const LibrdkafkaLinkInfo = "static musl_linux_arm64 from librdkafka-static-bundle-v2.3.0.tgz"
