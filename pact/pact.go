package pact

import (
	"encoding/binary"
	"errors"
)

type Type byte

const (
	CONNECT Type = iota + 1
	CONNACK
	PING
	PONG
	PUBLISH
	PUBACK
	SUBSCRIBE
	SUBACK
	UNSUBSCRIBE
	UNSUBACK
	PVTPUBLISH
	PVTPUBACK
	FORWARD
)

var (
	ErrVarintOutOfRange = errors.New("varint out of range")
)

func EncodeVarint(x int) []byte {
	var buf [5]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}

	if n > 4 {
		return nil
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

func DecodeVarint(b []byte) (int, int) {
	u, i := binary.Uvarint(b)
	return int(u), i
}

func (t Type) Encode(body []byte) []byte {
	ebody := append(EncodeVarint(len(body)), body...)
	cbody := append([]byte{byte(t)}, ebody...)
	return cbody
}

// int: seq.
// byte: topicID.
// []byte: remaining content.
func (t Type) PvtDecode(body []byte) (int, byte, []byte) {
	v, n := DecodeVarint(body)
	return v, body[n], body[n+1:]
}

func (t Type) PvtEncode(seq int, topicID byte, body []byte) []byte {
	b := append([]byte{topicID}, body...)
	seqBody := append(EncodeVarint(seq), b...)

	return append([]byte{byte(t)}, EncodeVarint(len(seqBody))...)
}
