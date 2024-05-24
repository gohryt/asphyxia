package percent_test

import (
	"net/url"
	"testing"

	"github.com/gohryt/asphyxia/bytes"
	"github.com/gohryt/asphyxia/encoding/percent"
)

const (
	LoremIpsum        = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
	LoremIpsumPercent = `Lorem%20ipsum%20dolor%20sit%20amet%2C%20consectetur%20adipiscing%20elit%2C%20sed%20do%20eiusmod%20tempor%20incididunt%20ut%20labore%20et%20dolore%20magna%20aliqua.%20Ut%20enim%20ad%20minim%20veniam%2C%20quis%20nostrud%20exercitation%20ullamco%20laboris%20nisi%20ut%20aliquip%20ex%20ea%20commodo%20consequat.%20Duis%20aute%20irure%20dolor%20in%20reprehenderit%20in%20voluptate%20velit%20esse%20cillum%20dolore%20eu%20fugiat%20nulla%20pariatur.%20Excepteur%20sint%20occaecat%20cupidatat%20non%20proident%2C%20sunt%20in%20culpa%20qui%20officia%20deserunt%20mollit%20anim%20id%20est%20laborum.`
)

const TestFailed = `Test failed
expected: %s
get:      %s`

func TestEncodeStd(t *testing.T) {
	result := url.PathEscape(LoremIpsum)

	if result != LoremIpsumPercent {
		t.Fatalf(TestFailed, LoremIpsumPercent, result)
	}
}

func TestEncode(t *testing.T) {
	result := percent.Encode(bytes.Buffer(LoremIpsum))

	if string(result) != LoremIpsumPercent {
		t.Fatalf(TestFailed, LoremIpsumPercent, string(result))
	}
}

func TestDecodeStd(t *testing.T) {
	result, err := url.PathUnescape(LoremIpsumPercent)
	if err != nil {
		t.Fatal(err)
	}

	if result != LoremIpsum {
		t.Fatalf(TestFailed, LoremIpsum, result)
	}
}

func TestDecode(t *testing.T) {
	result := percent.Decode(bytes.Buffer(LoremIpsumPercent))

	if string(result) != LoremIpsum {
		t.Fatalf(TestFailed, LoremIpsum, string(result))
	}
}

func BenchmarkEncodeStd(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = url.QueryEscape(LoremIpsum)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = percent.Encode(bytes.Buffer(LoremIpsum))
	}
}

func BenchmarkDecodeStd(b *testing.B) {
	err := error(nil)

	for i := 0; i < b.N; i += 1 {
		_, err = url.QueryUnescape(LoremIpsumPercent)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = percent.Decode(bytes.Buffer(LoremIpsumPercent))
	}
}
