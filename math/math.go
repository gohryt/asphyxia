package math

type (
	Integer interface {
		int8 | int16 | int32 | int64
	}

	Unsigned interface {
		uint8 | uint16 | uint32 | uint64
	}

	Float interface {
		float32 | float64
	}
)
