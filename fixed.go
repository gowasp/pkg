package pkg

import (
	"encoding/binary"
	"errors"
)

type Fixed byte

const (
	FIXED_CONNECT Fixed = iota + 1
	FIXED_CONNACK
	FIXED_PING
	FIXED_PONG
	FIXED_PUBLISH
	FIXED_PUBACK
	FIXED_SUBSCRIBE
	FIXED_SUBACK
	FIXED_UNSUBSCRIBE
	FIXED_UNSUBACK
	FIXED_FORWARD
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

func ConnectEncode(body []byte) []byte {
	vi := EncodeVarint(len(body))
	viBody := append([]byte{byte(len(vi))}, vi...)
	ebody := append(viBody, body...)
	cbody := append([]byte{byte(FIXED_CONNECT)}, ebody...)
	return cbody
}

func PubEncode(topic string, body []byte) []byte {
	t := append([]byte(topic), body...)
	tl := append([]byte{byte(len(topic))}, t...)
	b := append(EncodeVarint(len(tl)), tl...)

	return append([]byte{byte(FIXED_PUBLISH)}, b...)
}
