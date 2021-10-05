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

func (f Fixed) Encode(body []byte) []byte {
	evBody := append(EncodeVarint(len(body)), body...)
	varintLen := len(evBody)
	ebody := append([]byte{byte(varintLen)}, evBody...)
	cbody := append([]byte{byte(f)}, ebody...)
	return cbody
}

func PubDecode(body []byte) ([]byte, []byte) {
	return body[1 : 1+body[0]], body[1+body[0]:]
}

func PubEncode(topic string, body []byte) []byte {
	t := append([]byte(topic), body...)
	tl := append([]byte{byte(len(topic))}, t...)
	b := append(EncodeVarint(len(tl)), tl...)

	return append([]byte{byte(FIXED_PUBLISH)}, b...)
}

func PubEncodeSeq(seq int, topic string, body []byte) []byte {
	t := append([]byte(topic), body...)
	tl := append([]byte{byte(len(topic))}, t...)

	idbody := append(EncodeVarint(seq), tl...)
	return FIXED_PUBLISH.Encode(idbody)
}

func PubDecodeSeq(body []byte) (int, string, []byte, error) {
	v, n := DecodeVarint(body)

	begin := n + 1
	end := n + 1 + int(body[n])
	if end <= begin {
		return 0, "", nil, errors.New("error data")
	}

	if n+1+int(body[n]) >= len(body) {
		return 0, "", nil, errors.New("error data")

	}

	return v, string(body[n+1 : n+1+int(body[n])]), body[n+1+int(body[n]):], nil
}
