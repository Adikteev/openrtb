package openrtb_test

import (
	"testing"

	. "github.com/bsm/openrtb/v3"
	"github.com/bytedance/sonic"
)

func TestNumberOrString(t *testing.T) {
	var v NumberOrString

	if err := sonic.Unmarshal([]byte(`33`), &v); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if exp := NumberOrString(33); exp != v {
		t.Errorf("expected %#v, got %#v", exp, v)
	}

	if err := sonic.Unmarshal([]byte(`"34"`), &v); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if exp := NumberOrString(34); exp != v {
		t.Errorf("expected %#v, got %#v", exp, v)
	}
}

func TestStringOrNumber(t *testing.T) {
	var v StringOrNumber

	if err := sonic.Unmarshal([]byte(`33`), &v); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if exp := StringOrNumber("33"); exp != v {
		t.Errorf("expected %#v, got %#v", exp, v)
	}

	if err := sonic.Unmarshal([]byte(`"34"`), &v); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if exp := StringOrNumber("34"); exp != v {
		t.Errorf("expected %#v, got %#v", exp, v)
	}

	if err := sonic.Unmarshal([]byte(`""`), &v); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if exp := StringOrNumber(""); exp != v {
		t.Errorf("expected %#v, got %#v", exp, v)
	}
}
