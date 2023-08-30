// Note: keep in sync with "unicode/utf8"
// Description: This is pseudo-package created because of need in private properties of official go's utf8 package

package UTF8

const (
	RuneError = '\uFFFD'
	RuneSelf  = 0x80
	MaxRune   = '\U0010FFFF'
	UTFMax    = 4
)

const (
	SurrogateMin = 0xD800
	SurrogateMax = 0xDFFF
)

const (
	T1 = 0b00000000
	Tx = 0b10000000
	T2 = 0b11000000
	T3 = 0b11100000
	T4 = 0b11110000
	T5 = 0b11111000

	Maskx = 0b00111111
	Mask2 = 0b00011111
	Mask3 = 0b00001111
	Mask4 = 0b00000111

	Rune1Max uint32 = 1<<7 - 1
	Rune2Max uint32 = 1<<11 - 1
	Rune3Max uint32 = 1<<16 - 1

	locb = 0b10000000
	hicb = 0b10111111

	xx = 0xF1
	as = 0xF0
	s1 = 0x02
	s2 = 0x13
	s3 = 0x03
	s4 = 0x23
	s5 = 0x34
	s6 = 0x04
	s7 = 0x44
)

var first = [256]uint8{
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
	s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
	s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3,
	s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
}

type AcceptRange struct {
	lo uint8
	hi uint8
}

var AcceptRanges = [16]AcceptRange{
	0: {locb, hicb},
	1: {0xA0, hicb},
	2: {locb, 0x9F},
	3: {0x90, hicb},
	4: {locb, 0x8F},
}

func DecodeRune(p []byte) (r rune, size int) {
	n := len(p)
	if n < 1 {
		return RuneError, 0
	}
	p0 := p[0]
	x := first[p0]
	if x >= as {
		mask := rune(x) << 31 >> 31
		return rune(p[0])&^mask | RuneError&mask, 1
	}
	sz := int(x & 7)
	accept := AcceptRanges[x>>4]
	if n < sz {
		return RuneError, 1
	}
	b1 := p[1]
	if b1 < accept.lo || accept.hi < b1 {
		return RuneError, 1
	}
	if sz <= 2 {
		return rune(p0&Mask2)<<6 | rune(b1&Maskx), 2
	}
	b2 := p[2]
	if b2 < locb || hicb < b2 {
		return RuneError, 1
	}
	if sz <= 3 {
		return rune(p0&Mask3)<<12 | rune(b1&Maskx)<<6 | rune(b2&Maskx), 3
	}
	b3 := p[3]
	if b3 < locb || hicb < b3 {
		return RuneError, 1
	}
	return rune(p0&Mask4)<<18 | rune(b1&Maskx)<<12 | rune(b2&Maskx)<<6 | rune(b3&Maskx), 4
}
