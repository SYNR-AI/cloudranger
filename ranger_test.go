package cloudranger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	cr := New()

	tests := []struct {
		name           string
		ip             string
		expectedCloud  string
		expectedRegion string
		found          bool
	}{
		{
			name:           "valid IPv4 address in Amazon Web Services",
			ip:             "3.5.140.101",
			expectedCloud:  "aws",
			expectedRegion: "ap-northeast-2",
			found:          true,
		},
		{
			name:           "valid IPv6 address in Amazon Web Services",
			ip:             "2a05:d077:6081::1",
			expectedCloud:  "aws",
			expectedRegion: "eu-north-1",
			found:          true,
		},
		{
			name:           "valid IPv4 address in Google Cloud Platform",
			ip:             "34.35.1.2",
			expectedCloud:  "gcp",
			expectedRegion: "africa-south1",
			found:          true,
		},
		{
			name:           "valid IPv6 address in Google Cloud Platform",
			ip:             "2600:1900:4010::0000:1",
			expectedCloud:  "gcp",
			expectedRegion: "europe-west1",
			found:          true,
		},
		{
			name:  "non-cloud IP address (localhost)",
			ip:    "127.0.0.1",
			found: false,
		},
		{
			name:  "non-cloud IP address (LAN)",
			ip:    "192.168.1.1",
			found: false,
		},
		{
			name:  "not an IP address",
			ip:    "just a random string",
			found: false,
		},
		{
			name:  "valid IPv4 address in Aliyun",
			ip:    "106.15.46.159",
			found: false,
		},
		{
			name:  "valid IPv4 address in China Telecom",
			ip:    "61.152.51.5",
			found: false,
		},
		{
			name:           "valid IPv4 address in GCP",
			ip:             "34.120.54.55",
			found:          true,
			expectedCloud:  "gcp",
			expectedRegion: "global",
		},
		{
			name:           "valid IPv6 address in GCP",
			ip:             "2600:1900:4080::1111",
			found:          true,
			expectedCloud:  "gcp",
			expectedRegion: "asia-southeast1",
		},
		{
			name:           "valid IPv4 address in Cloudflare",
			ip:             "104.21.40.8",
			found:          true,
			expectedCloud:  "cloudflare",
			expectedRegion: "",
		},
		{
			name:           "valid IPv6 address in Cloudflare",
			ip:             "2405:b500::1111",
			found:          true,
			expectedCloud:  "cloudflare",
			expectedRegion: "",
		},
		{
			name:           "valid IPv4 address in Azure",
			ip:             "40.121.67.30",
			found:          true,
			expectedCloud:  "azure",
			expectedRegion: "eastus",
		},
		{
			name:           "valid IPv6 address in Azure",
			ip:             "2603:1030:40c:1::118",
			found:          true,
			expectedCloud:  "azure",
			expectedRegion: "eastus2",
		},
		{
			name:           "valid IPv4 address in Oracle Cloud (ca-toronto-1)",
			ip:             "192.29.14.122",
			found:          true,
			expectedCloud:  "oracle",
			expectedRegion: "ca-toronto-1",
		},
		{
			name:           "valid IPv4 address in Oracle Cloud (ap-tokyo-1)",
			ip:             "192.29.44.129",
			found:          true,
			expectedCloud:  "oracle",
			expectedRegion: "ap-tokyo-1",
		},
		{
			name:           "valid IPv4 address in AWS",
			ip:             "99.84.188.14",
			found:          true,
			expectedCloud:  "aws",
			expectedRegion: "GLOBAL",
		},
		{
			name:           "valid IPv6 address in AWS",
			ip:             "2600:1fb8:6000::1111",
			found:          true,
			expectedCloud:  "aws",
			expectedRegion: "us-east-2",
		},
		{
			name:           "valid IPv4 address in Linode",
			ip:             "50.116.20.1",
			found:          true,
			expectedCloud:  "linode",
			expectedRegion: "US-TX",
		},
		{
			name:           "valid IPv6 address in Linode",
			ip:             "2a01:7e03::1111",
			found:          true,
			expectedCloud:  "linode",
			expectedRegion: "US-CA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ipinfo, found := cr.GetIP(tt.ip)
			assert.Equal(t, tt.found, found)
			assert.Equal(t, tt.expectedCloud, ipinfo.Cloud())
			assert.Equal(t, tt.expectedRegion, ipinfo.Region())
		})
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = New()
	}
}

func BenchmarkGetIP(b *testing.B) {
	b.StopTimer()
	ranger := New()
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		_, _ = ranger.GetIP("34.35.1.3")
	}
}
