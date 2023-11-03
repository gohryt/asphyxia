#include "textflag.h"

// func reset(buffer *Buffer)
TEXT ·reset(SB), NOSPLIT, $0-8
	MOVQ buffer+0(FP), AX
    MOVQ $0, 8(AX)
	RET

// func asString(buffer *Buffer) string
TEXT ·asString(SB), NOSPLIT, $0-24
	MOVQ buffer+0(FP), AX
	MOVQ 0(AX), BX
	MOVQ 8(AX), CX
	MOVQ BX, ret+8(FP)
	MOVQ CX, ret+16(FP)
	RET
