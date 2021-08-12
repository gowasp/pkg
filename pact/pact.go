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

func PubDecode(body []byte) (int, string, []byte) {
	v, n := DecodeVarint(body)
	return v, string(body[n+1 : n+1+int(body[n])]), body[n+1+int(body[n]):]
}

func PubEncode(seq int, topic string, body []byte) []byte {
	t := append([]byte(topic), body...)
	tl := len([]byte(topic))
	t1 := append([]byte{byte(tl)}, t...)
	s := append(EncodeVarint(seq), t1...)

	b := append(EncodeVarint(len(s)), s...)
	return append([]byte{byte(PUBLISH)}, b...)
}

// int: seq.
// byte: topicID.
// []byte: remaining content.
func PvtPubDecode(body []byte) (int, byte, []byte) {
	v, n := DecodeVarint(body)
	return v, body[n], body[n+1:]
}

func PvtPubAckEncode(seq int) []byte {
	seqBody := EncodeVarint(seq)

	v := append(EncodeVarint(len(seqBody)), seqBody...)

	return append([]byte{byte(PVTPUBACK)}, v...)
}
