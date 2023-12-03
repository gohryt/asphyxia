//go:build amd64

package JIT

type (
	RegisterID uint32
)

const (
	AX RegisterID = iota + 0
	CX
	DX
	BX
	SP
	BP
	SI
	DI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
)

const (
	Void = 0

	BaseStart = 32
	BaseEnd   = 44

	IntStart = 32
	IntEnd   = 41

	// Abstract signed integer type that has a native size.
	IntPtr = 32
	// Abstract unsigned integer type that has a native size.
	UIntPtr = 33

	// 8-bit signed integer type.
	Int8 = 34
	// 8-bit unsigned integer type.
	UInt8 = 35
	// 16-bit signed integer type.
	Int16 = 36
	// 16-bit unsigned integer type.
	UInt16 = 37
	// 32-bit signed integer type.
	Int32 = 38
	// 32-bit unsigned integer type.
	UInt32 = 39
	// 64-bit signed integer type.
	Int64 = 40
	// 64-bit unsigned integer type.
	UInt64 = 41

	FloatStart = 42
	FloatEnd   = 44

	// 32-bit floating point type.
	Float32 = 42
	// 64-bit floating point type.
	Float64 = 43
	// 80-bit floating point type.
	Float80 = 44

	MaskStart = 45
	MaskEnd   = 48

	// 8-bit opmask register (K).
	Mask8 = 45
	// 16-bit opmask register (K).
	Mask16 = 46
	// 32-bit opmask register (K).
	Mask32 = 47
	// 64-bit opmask register (K).
	Mask64 = 48

	MMXStart = 49
	MMXEnd   = 50

	// 64-bit MMX register only used for 32 bits.
	MMX32 = 49
	// 64-bit MMX register.
	MMX64 = 50

	Vec32Start = 51
	Vec32End   = 60

	Int8x4    = 51
	UInt8x4   = 52
	Int16x2   = 53
	UInt16x2  = 54
	Int32x1   = 55
	UInt32x1  = 56
	Float32x1 = 59

	Vec64Start = 61
	Vec64End   = 70

	Int8x8    = 61
	UInt8x8   = 62
	Int16x4   = 63
	UInt16x4  = 64
	Int32x2   = 65
	UInt32x2  = 66
	Int64x1   = 67
	UInt64x1  = 68
	Float32x2 = 69
	Float64x1 = 70

	Vec128Start = 71
	Vec128End   = 80

	Int8x16   = 71
	UInt8x16  = 72
	Int16x8   = 73
	UInt16x8  = 74
	Int32x4   = 75
	UInt32x4  = 76
	Int64x2   = 77
	UInt64x2  = 78
	Float32x4 = 79
	Float64x2 = 80

	Vec256Start = 81
	Vec256End   = 90

	Int8x32   = 81
	UInt8x32  = 82
	Int16x16  = 83
	UInt16x16 = 84
	Int32x8   = 85
	UInt32x8  = 86
	Int64x4   = 87
	UInt64x4  = 88
	Float32x8 = 89
	Float64x4 = 90

	Vec512Start = 91
	Vec512End   = 100

	Int8x64    = 91
	UInt8x64   = 92
	Int16x32   = 93
	UInt16x32  = 94
	Int32x16   = 95
	UInt32x16  = 96
	Int64x8    = 97
	UInt64x8   = 98
	Float32x16 = 99
	Float64x8  = 100
)

func (builder *builder) MOV(register RegisterID, value uint32) {
}

func (builder *builder) RET() {
}
