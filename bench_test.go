package openrtb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bytedance/sonic"
)

func BenchmarkBidRequest_Unmarshal(b *testing.B) {
	data, err := os.ReadFile(filepath.Join("testdata", "breq.video.json"))
	if err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var req *BidRequest
		if err := sonic.Unmarshal(data, &req); err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkBidRequest_Marshal(b *testing.B) {
	data, err := os.ReadFile(filepath.Join("testdata", "breq.video.json"))
	if err != nil {
		b.Fatal(err.Error())
	}

	var req *BidRequest
	if err := sonic.Unmarshal(data, &req); err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sonic.Marshal(req); err != nil {
			b.Fatal(err.Error())
		}
	}
}
