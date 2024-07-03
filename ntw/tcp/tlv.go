package tcp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	BinaryType uint8 = iota + 1
	StringType

	MaxPayloadSize = 10 << 20 // 10Mb
)

var ErrMaxPayloadSize = errors.New("maximum payload size exceeded")

// Type-Length-Value Encoding (Dynamic Buffer Size)
// Impl -> 5byte Header: 1byte type | 4byte length
type Payload interface {
	fmt.Stringer
	io.ReaderFrom
	io.WriterTo
	Bytes() []byte
}

// Binary implements Payload
type Binary []byte

func (m Binary) Bytes() []byte  { return m }
func (m Binary) String() string { return string(m) }

func (m Binary) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, BinaryType) // 1byte type
	if err != nil {
		return 0, err
	}

	var n int64 = 1
	err = binary.Write(w, binary.BigEndian, uint32(len(m))) //4byte size
	if err != nil {
		return n, err
	}
	n += 4

	o, err := w.Write(m)

	return n + int64(o), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {
	var size uint32
  err := binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return 0, err
	}

	if size > MaxPayloadSize {
		return 4, ErrMaxPayloadSize
	}

	*m = make([]byte, size)
	p, err := r.Read(*m) // payload

	return int64(p) + 4, err
}

// String implements Payload
type String string

func (m String) Bytes() []byte  { return []byte(m) }
func (m String) String() string { return string(m) }

func (m String) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, StringType)
	if err != nil {
		return 0, err
	}

	var n int64 = 1
	err = binary.Write(w, binary.BigEndian, uint32(len(m)))
	if err != nil {
		return n, err
	}
	n += 4

	o, err := w.Write([]byte(m))
	return n + int64(o), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	var size uint32
  err := binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, size)
	p, err := r.Read(buf)
	if err != nil {
		return 4, err
	}

	*m = String(buf)

	return int64(p) + 4, nil
}
