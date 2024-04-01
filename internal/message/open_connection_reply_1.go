package message

import (
	"encoding/binary"
	"io"
)

type OpenConnectionReply1 struct {
	ServerGUID             int64
	Secure                 bool
	ServerPreferredMTUSize uint16
	Cookie                 uint32
}

func (pk *OpenConnectionReply1) UnmarshalBinary(data []byte) error {
	var offset int
	if len(data) < 27 || len(data) < 27+int(data[24])*4 {
		return io.ErrUnexpectedEOF
	}
	// Magic: 16 bytes.
	pk.ServerGUID = int64(binary.BigEndian.Uint64(data[16:]))
	pk.Secure = data[24] != 0
	if pk.Secure {
		offset = 4
		pk.Cookie = binary.BigEndian.Uint32(data[25:29])
	}
	pk.ServerPreferredMTUSize = binary.BigEndian.Uint16(data[25+offset:])
	return nil
}

func (pk *OpenConnectionReply1) MarshalBinary() (data []byte, err error) {
	offset := 0
	if pk.Secure {
		offset = 4
	}
	b := make([]byte, 28+offset)
	b[0] = IDOpenConnectionReply1
	copy(b[1:], unconnectedMessageSequence[:])
	binary.BigEndian.PutUint64(b[17:], uint64(pk.ServerGUID))
	if pk.Secure {
		b[25] = 1
		binary.BigEndian.PutUint32(b[26:], pk.Cookie)
	}
	binary.BigEndian.PutUint16(b[26+offset:], pk.ServerPreferredMTUSize)
	return b, nil
}
