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

	USSR = `Союз нерушимый республик свободных
Сплотила навеки Великая Русь.
Да здравствует созданный волей народов
Единый, могучий Советский Союз!

Славься, Отечество наше свободное,
Дружбы народов надёжный оплот!
Партия Ленина — сила народная
Нас к торжеству коммунизма ведёт!

Сквозь грозы сияло нам солнце свободы,
И Ленин великий нам путь озарил:
На правое дело он поднял народы,
На труд и на подвиги нас вдохновил!

Славься, Отечество наше свободное,
Дружбы народов надёжный оплот!
Партия Ленина — сила народная
Нас к торжеству коммунизма ведёт!

В победе бессмертных идей коммунизма
Мы видим грядущее нашей страны,
И Красному знамени славной Отчизны
Мы будем всегда беззаветно верны!

Славься, Отечество наше свободное,
Дружбы народов надёжный оплот!
Партия Ленина — сила народная
Нас к торжеству коммунизма ведёт!`

	USSRPercent = `%D0%A1%D0%BE%D1%8E%D0%B7%20%D0%BD%D0%B5%D1%80%D1%83%D1%88%D0%B8%D0%BC%D1%8B%D0%B9%20%D1%80%D0%B5%D1%81%D0%BF%D1%83%D0%B1%D0%BB%D0%B8%D0%BA%20%D1%81%D0%B2%D0%BE%D0%B1%D0%BE%D0%B4%D0%BD%D1%8B%D1%85%0A%D0%A1%D0%BF%D0%BB%D0%BE%D1%82%D0%B8%D0%BB%D0%B0%20%D0%BD%D0%B0%D0%B2%D0%B5%D0%BA%D0%B8%20%D0%92%D0%B5%D0%BB%D0%B8%D0%BA%D0%B0%D1%8F%20%D0%A0%D1%83%D1%81%D1%8C.%0A%D0%94%D0%B0%20%D0%B7%D0%B4%D1%80%D0%B0%D0%B2%D1%81%D1%82%D0%B2%D1%83%D0%B5%D1%82%20%D1%81%D0%BE%D0%B7%D0%B4%D0%B0%D0%BD%D0%BD%D1%8B%D0%B9%20%D0%B2%D0%BE%D0%BB%D0%B5%D0%B9%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BE%D0%B2%0A%D0%95%D0%B4%D0%B8%D0%BD%D1%8B%D0%B9%2C%20%D0%BC%D0%BE%D0%B3%D1%83%D1%87%D0%B8%D0%B9%20%D0%A1%D0%BE%D0%B2%D0%B5%D1%82%D1%81%D0%BA%D0%B8%D0%B9%20%D0%A1%D0%BE%D1%8E%D0%B7%21%0A%0A%D0%A1%D0%BB%D0%B0%D0%B2%D1%8C%D1%81%D1%8F%2C%20%D0%9E%D1%82%D0%B5%D1%87%D0%B5%D1%81%D1%82%D0%B2%D0%BE%20%D0%BD%D0%B0%D1%88%D0%B5%20%D1%81%D0%B2%D0%BE%D0%B1%D0%BE%D0%B4%D0%BD%D0%BE%D0%B5%2C%0A%D0%94%D1%80%D1%83%D0%B6%D0%B1%D1%8B%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BE%D0%B2%20%D0%BD%D0%B0%D0%B4%D1%91%D0%B6%D0%BD%D1%8B%D0%B9%20%D0%BE%D0%BF%D0%BB%D0%BE%D1%82%21%0A%D0%9F%D0%B0%D1%80%D1%82%D0%B8%D1%8F%20%D0%9B%D0%B5%D0%BD%D0%B8%D0%BD%D0%B0%20%E2%80%94%20%D1%81%D0%B8%D0%BB%D0%B0%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BD%D0%B0%D1%8F%0A%D0%9D%D0%B0%D1%81%20%D0%BA%20%D1%82%D0%BE%D1%80%D0%B6%D0%B5%D1%81%D1%82%D0%B2%D1%83%20%D0%BA%D0%BE%D0%BC%D0%BC%D1%83%D0%BD%D0%B8%D0%B7%D0%BC%D0%B0%20%D0%B2%D0%B5%D0%B4%D1%91%D1%82%21%0A%0A%D0%A1%D0%BA%D0%B2%D0%BE%D0%B7%D1%8C%20%D0%B3%D1%80%D0%BE%D0%B7%D1%8B%20%D1%81%D0%B8%D1%8F%D0%BB%D0%BE%20%D0%BD%D0%B0%D0%BC%20%D1%81%D0%BE%D0%BB%D0%BD%D1%86%D0%B5%20%D1%81%D0%B2%D0%BE%D0%B1%D0%BE%D0%B4%D1%8B%2C%0A%D0%98%20%D0%9B%D0%B5%D0%BD%D0%B8%D0%BD%20%D0%B2%D0%B5%D0%BB%D0%B8%D0%BA%D0%B8%D0%B9%20%D0%BD%D0%B0%D0%BC%20%D0%BF%D1%83%D1%82%D1%8C%20%D0%BE%D0%B7%D0%B0%D1%80%D0%B8%D0%BB:%0A%D0%9D%D0%B0%20%D0%BF%D1%80%D0%B0%D0%B2%D0%BE%D0%B5%20%D0%B4%D0%B5%D0%BB%D0%BE%20%D0%BE%D0%BD%20%D0%BF%D0%BE%D0%B4%D0%BD%D1%8F%D0%BB%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D1%8B%2C%0A%D0%9D%D0%B0%20%D1%82%D1%80%D1%83%D0%B4%20%D0%B8%20%D0%BD%D0%B0%20%D0%BF%D0%BE%D0%B4%D0%B2%D0%B8%D0%B3%D0%B8%20%D0%BD%D0%B0%D1%81%20%D0%B2%D0%B4%D0%BE%D1%85%D0%BD%D0%BE%D0%B2%D0%B8%D0%BB%21%0A%0A%D0%A1%D0%BB%D0%B0%D0%B2%D1%8C%D1%81%D1%8F%2C%20%D0%9E%D1%82%D0%B5%D1%87%D0%B5%D1%81%D1%82%D0%B2%D0%BE%20%D0%BD%D0%B0%D1%88%D0%B5%20%D1%81%D0%B2%D0%BE%D0%B1%D0%BE%D0%B4%D0%BD%D0%BE%D0%B5%2C%0A%D0%94%D1%80%D1%83%D0%B6%D0%B1%D1%8B%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BE%D0%B2%20%D0%BD%D0%B0%D0%B4%D1%91%D0%B6%D0%BD%D1%8B%D0%B9%20%D0%BE%D0%BF%D0%BB%D0%BE%D1%82%21%0A%D0%9F%D0%B0%D1%80%D1%82%D0%B8%D1%8F%20%D0%9B%D0%B5%D0%BD%D0%B8%D0%BD%D0%B0%20%E2%80%94%20%D1%81%D0%B8%D0%BB%D0%B0%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BD%D0%B0%D1%8F%0A%D0%9D%D0%B0%D1%81%20%D0%BA%20%D1%82%D0%BE%D1%80%D0%B6%D0%B5%D1%81%D1%82%D0%B2%D1%83%20%D0%BA%D0%BE%D0%BC%D0%BC%D1%83%D0%BD%D0%B8%D0%B7%D0%BC%D0%B0%20%D0%B2%D0%B5%D0%B4%D1%91%D1%82%21%0A%0A%D0%92%20%D0%BF%D0%BE%D0%B1%D0%B5%D0%B4%D0%B5%20%D0%B1%D0%B5%D1%81%D1%81%D0%BC%D0%B5%D1%80%D1%82%D0%BD%D1%8B%D1%85%20%D0%B8%D0%B4%D0%B5%D0%B9%20%D0%BA%D0%BE%D0%BC%D0%BC%D1%83%D0%BD%D0%B8%D0%B7%D0%BC%D0%B0%0A%D0%9C%D1%8B%20%D0%B2%D0%B8%D0%B4%D0%B8%D0%BC%20%D0%B3%D1%80%D1%8F%D0%B4%D1%83%D1%89%D0%B5%D0%B5%20%D0%BD%D0%B0%D1%88%D0%B5%D0%B9%20%D1%81%D1%82%D1%80%D0%B0%D0%BD%D1%8B%2C%0A%D0%98%20%D0%9A%D1%80%D0%B0%D1%81%D0%BD%D0%BE%D0%BC%D1%83%20%D0%B7%D0%BD%D0%B0%D0%BC%D0%B5%D0%BD%D0%B8%20%D1%81%D0%BB%D0%B0%D0%B2%D0%BD%D0%BE%D0%B9%20%D0%9E%D1%82%D1%87%D0%B8%D0%B7%D0%BD%D1%8B%0A%D0%9C%D1%8B%20%D0%B1%D1%83%D0%B4%D0%B5%D0%BC%20%D0%B2%D1%81%D0%B5%D0%B3%D0%B4%D0%B0%20%D0%B1%D0%B5%D0%B7%D0%B7%D0%B0%D0%B2%D0%B5%D1%82%D0%BD%D0%BE%20%D0%B2%D0%B5%D1%80%D0%BD%D1%8B%21%0A%0A%D0%A1%D0%BB%D0%B0%D0%B2%D1%8C%D1%81%D1%8F%2C%20%D0%9E%D1%82%D0%B5%D1%87%D0%B5%D1%81%D1%82%D0%B2%D0%BE%20%D0%BD%D0%B0%D1%88%D0%B5%20%D1%81%D0%B2%D0%BE%D0%B1%D0%BE%D0%B4%D0%BD%D0%BE%D0%B5%2C%0A%D0%94%D1%80%D1%83%D0%B6%D0%B1%D1%8B%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BE%D0%B2%20%D0%BD%D0%B0%D0%B4%D1%91%D0%B6%D0%BD%D1%8B%D0%B9%20%D0%BE%D0%BF%D0%BB%D0%BE%D1%82%21%0A%D0%9F%D0%B0%D1%80%D1%82%D0%B8%D1%8F%20%D0%9B%D0%B5%D0%BD%D0%B8%D0%BD%D0%B0%20%E2%80%94%20%D1%81%D0%B8%D0%BB%D0%B0%20%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BD%D0%B0%D1%8F%0A%D0%9D%D0%B0%D1%81%20%D0%BA%20%D1%82%D0%BE%D1%80%D0%B6%D0%B5%D1%81%D1%82%D0%B2%D1%83%20%D0%BA%D0%BE%D0%BC%D0%BC%D1%83%D0%BD%D0%B8%D0%B7%D0%BC%D0%B0%20%D0%B2%D0%B5%D0%B4%D1%91%D1%82%21`
)

const TestFailed = `Test failed
expected: %s
get:      %s`

func TestEncodeStd(t *testing.T) {
	result := url.PathEscape(LoremIpsum)

	if result != LoremIpsumPercent {
		t.Fatalf(TestFailed, LoremIpsumPercent, result)
	}

	result = url.PathEscape(USSR)

	if result != USSRPercent {
		t.Fatalf(TestFailed, USSRPercent, result)
	}
}

func TestEncode(t *testing.T) {
	result := percent.Encode(bytes.Buffer(LoremIpsum))

	if string(result) != LoremIpsumPercent {
		t.Fatalf(TestFailed, LoremIpsumPercent, string(result))
	}

	result = percent.Encode(bytes.Buffer(USSR))

	if string(result) != USSRPercent {
		t.Fatalf(TestFailed, USSRPercent, string(result))
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

	result, err = url.PathUnescape(USSRPercent)
	if err != nil {
		t.Fatal(err)
	}

	if result != USSR {
		t.Fatalf(TestFailed, USSR, result)
	}
}

func TestDecode(t *testing.T) {
	result := percent.Decode(bytes.Buffer(LoremIpsumPercent))

	if string(result) != LoremIpsum {
		t.Fatalf(TestFailed, LoremIpsum, string(result))
	}

	result = percent.Decode(bytes.Buffer(USSRPercent))

	if string(result) != USSR {
		t.Fatalf(TestFailed, USSR, string(result))
	}
}

func BenchmarkEncodeStd(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = url.PathEscape(LoremIpsum)
		_ = url.PathEscape(USSR)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = percent.Encode(bytes.Buffer(LoremIpsum))
		_ = percent.Encode(bytes.Buffer(USSR))
	}
}

func BenchmarkDecodeStd(b *testing.B) {
	err := error(nil)

	for i := 0; i < b.N; i += 1 {
		_, err = url.PathUnescape(LoremIpsumPercent)
		if err != nil {
			b.Fatal(err)
		}

		_, err = url.PathUnescape(USSRPercent)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		_ = percent.Decode(bytes.Buffer(LoremIpsumPercent))
		_ = percent.Decode(bytes.Buffer(USSR))
	}
}
