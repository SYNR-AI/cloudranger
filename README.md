# cloudranger

`cloudranger` is a Go library designed to identify cloud provider information from IP addresses, with current support for AWS and GCP.

It functions without any external runtime dependencies, as IP range data is stored internally. Meant for high throughput, low-latency environments, `cloudranger` also focuses on rapid startup, loading in under 4ms. You can verify this on your system by running `make bench` and checking the `BenchmarkNew` results.

New releases are automatically created in response to updates in cloud providers' IP range information. This process, facilitated through GitHub Actions, is executed weekly to ensure the library remains up-to-date.

The inspiration for `cloudranger` came from a similar library found at https://github.com/kubernetes/registry.k8s.io, used by the Kubernetes OCI registry for redirecting requests to the appropriate cloud provider. We developed `cloudranger` to provide a standalone library adaptable for various projects, offering greater control for our specific use cases and minimizing the impact of upstream changes. Unlike the original project, which uses its own trie implementation, `cloudranger` depends on github.com/infobloxopen/go-trees. While both implementations have not been directly benchmarked against each other, their performance is expected to be comparable.

## Supported Providers

- [x] AWS
- [x] GCP
- [x] Azure
- [x] Linode
- [x] Cloudflare
- [x] Fastly
- [x] DigitalOcean
- [x] Oracle Cloud
- [ ] OVH
- [ ] Vultr
- [ ] Aliyun
- [ ] Tencent Cloud
- [ ] UCloud
- [ ] Huawei Cloud

## Usage

```sh
go get github.com/planetscale/cloudranger
```

```go
package main

import (
	"fmt"

	"github.com/planetscale/cloudranger"
)

func main() {
	ranger := cloudranger.New()
	ipinfo, found := ranger.GetIP("3.5.140.101")
	if found {
		fmt.Printf("cloud: %s, region: %s\n", ipinfo.Cloud(), ipinfo.Region())
	}
}
```

A small cli is included in [cmd/cloudranger](cmd/cloudranger). It is EXPERIMENTAL and its behavior, flags, and output is likely to change.

```sh
$ go run cmd/cloudranger/main.go 3.5.140.101

{"cloud":"AWS","region":"ap-northeast-2"}
```

## Testing and Benchmarks

```sh
make lint
make test
make bench
```

```
goos: darwin
goarch: arm64
pkg: github.com/SYNR-AI/cloudranger
cpu: Apple M1 Pro (proc 10:8:2)
mem: 32 GB

BenchmarkNew-10               42          28786575 ns/op         9518694 B/op     233498 allocs/op
BenchmarkGetIP-10        8359640               139.6 ns/op            64 B/op          2 allocs/op
```

## IP Range Database Updates

IP range data is sourced from:

- AWS: https://ip-ranges.amazonaws.com/ip-ranges.json
- GCP: https://www.gstatic.com/ipranges/cloud.json

A GitHub Actions workflow is run weekly to update the IP range data if changed by the supported cloud providers. A new version is created and tagged if changes are detected. Use dependabot or renovate to automate updates to the latest version.
