package uring

type (
	CQE struct {
		UserData uint64
		Res      int32
		Flags    uint32
	}
)
