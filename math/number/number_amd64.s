#include "textflag.h"

// func Add(to, value Number) Number
TEXT ·Add(SB), NOSPLIT, $0-48
	MOVQ to_len   + 8(FP), R8
	MOVQ value_len+24(FP), R9

	MOVQ $0, to+32(FP)
	MOVQ $0, to+40(FP)
	RET
