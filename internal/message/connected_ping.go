package message

import (
	"encoding/binary"
	"io"
)

type ConnectedPing struct {
	ClientTimestamp int64
}

func (pk *ConnectedPing) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return io.ErrUnexpectedEOF
	}
	pk.ClientTimestamp = int64(binary.BigEndian.Uint64(data))
	return nil
}

func (pk *ConnectedPing) MarshalBinary() (data []byte, err error) {
	b := make([]byte, 9)
	b[0] = IDConnectedPing
	binary.BigEndian.PutUint64(b[1:], uint64(pk.ClientTimestamp))
	return b, nil
}
