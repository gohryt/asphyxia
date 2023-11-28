#include "textflag.h"

// func BigEndianCopyUint16(to []byte, from ...uint16) int
TEXT ·BigEndianCopyUint16(SB), NOSPLIT, $0-56
	MOVQ to+0   (FP), R8
	MOVQ from_base+24(FP), R9
	MOVQ to+8   (FP), R10
	MOVQ from_len+32(FP), R11

	MOVQ $0, R12
	MOVQ $0, R13

	SHRQ $1, R10

	CMPQ R10, R11
	JLT  cycle

	MOVQ R11, R10

cycle:
	CMPQ R12, R10
	JGE  return

	ADDQ $1, R12
	ADDQ $2, R13

	JMP cycle

return:
	MOVQ R13, ret+48(FP)
	RET
