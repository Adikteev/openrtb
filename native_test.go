package openrtb_test

import (
	"github.com/tinylib/msgp/msgp"
	"reflect"
	"testing"

	. "github.com/bsm/openrtb/v3"
)

func TestNative(t *testing.T) {
	var subject *Native
	if err := fixture("native", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Native{
		Request: msgp.Raw(`"PAYLOAD"`),
		Version: "2",
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
