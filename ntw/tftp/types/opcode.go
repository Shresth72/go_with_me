package types

const (
	DatagramSize = 516              // max to avoid fragmentation
	BlockSize    = DatagramSize - 4 // 4byte header
)

/* 
RFC 1350

OpCode - first 2bytes of the header
*/
type OpCode uint16

const (
	OpRRQ OpCode = iota + 1
	_            // no WRQ support
	OpData
	OpAck
	OpErr
)

type ErrCode uint16

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser
)
