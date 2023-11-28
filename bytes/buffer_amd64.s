#include "textflag.h"

// func Reset(buffer *Buffer)
TEXT ·Reset(SB), NOSPLIT, $0-8
	MOVQ buffer(FP), AX
	MOVQ $0, 8(AX)
	RET

// func String(buffer *Buffer) string
TEXT ·String(SB), NOSPLIT, $0-24
	MOVQ buffer(FP), AX
	MOVQ (AX), BX
	MOVQ 8(AX), CX
	MOVQ BX, ret_base+8(FP)
	MOVQ CX, ret_len+16(FP)
	RET
