//go:build linux

package uring

import "unsafe"

type (
	SQEU1 uint64

	SQEU1CmdOpPad1 struct {
		CmdOp uint32
		Pad1  uint32
	}
)

func (u1 *SQEU1) GetOff() uint64 {
	return *(*uint64)(unsafe.Pointer(u1))
}

func (u1 *SQEU1) GetAddr2() uint64 {
	return *(*uint64)(unsafe.Pointer(u1))
}

func (u1 *SQEU1) GetCmdOpPad1() *SQEU1CmdOpPad1 {
	return (*SQEU1CmdOpPad1)(unsafe.Pointer(u1))
}

type (
	SQEU2 uint64

	SQEU2LevelOptname struct {
		Level   uint32
		Optname uint32
	}
)

func (u2 *SQEU2) GetAddr() uint64 {
	return *(*uint64)(unsafe.Pointer(u2))
}

func (u2 *SQEU2) GetSpliceOffIn() uint64 {
	return *(*uint64)(unsafe.Pointer(u2))
}

func (u2 *SQEU2) GetLevelOptname() *SQEU2LevelOptname {
	return (*SQEU2LevelOptname)(unsafe.Pointer(u2))
}

type (
	SQEU3 uint32
)

func (u3 *SQEU3) GetRWFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetFsyncFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetPollEvents() uint16 {
	return *(*uint16)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetPoll32Events() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetSyncRangeFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetMsgFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetTimeoutFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetAcceptFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetCancelFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetOpenFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetStatxFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetFadviseAdvice() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetSpliceFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetRenameFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetUnlinkFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetHardlinkFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetXattrFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetMsgRingFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetUringCmdFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetWaitidFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetFutexFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

func (u3 *SQEU3) GetInstallFDFlags() uint32 {
	return *(*uint32)(unsafe.Pointer(u3))
}

type (
	SQEU4 uint16
)

func (u4 *SQEU4) GetBufIndex() uint16 {
	return *(*uint16)(unsafe.Pointer(u4))
}

func (u4 *SQEU4) GetBufGroup() uint16 {
	return *(*uint16)(unsafe.Pointer(u4))
}

type (
	SQEU5 uint32

	SQEU5AddrLenPad3 struct {
		AddrLen uint16
		Pad3    [1]uint16
	}
)

func (u5 *SQEU5) GetSpliceFDIn() int32 {
	return *(*int32)(unsafe.Pointer(u5))
}

func (u5 *SQEU5) GetFileIndex() uint32 {
	return *(*uint32)(unsafe.Pointer(u5))
}

func (u5 *SQEU5) GetOptlen() uint32 {
	return *(*uint32)(unsafe.Pointer(u5))
}

func (u5 *SQEU5) GetAddrLenPad3() *SQEU5AddrLenPad3 {
	return (*SQEU5AddrLenPad3)(unsafe.Pointer(u5))
}

type (
	SQEU6 [2]uint64

	SQEU6Addr3Pad2 struct {
		Addr3 uint64
		Pad2  [1]uint64
	}
)

func (u6 *SQEU6) GetAddr3Pad2() *SQEU6Addr3Pad2 {
	return (*SQEU6Addr3Pad2)(unsafe.Pointer(u6))
}

func (u6 *SQEU6) GetOptval() uint64 {
	return *(*uint64)(unsafe.Pointer(u6))
}

type (
	IOURingSQE struct {
		Opcode uint8
		Flags  uint8
		IOPrio uint16
		FD     int32
		SQEU1
		SQEU2
		Len uint32
		SQEU3
		UserData uint64
		SQEU4
		Personality uint16
		SQEU5
		SQEU6
	}
)
