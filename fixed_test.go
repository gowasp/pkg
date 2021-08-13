package pkg

import (
	"testing"
	"time"
)

func TestPubDecode(t *testing.T) {
	topic := append([]byte("a/b/c"), []byte("pact")...)
	topic1 := append([]byte{byte(len([]byte("a/b/c")))}, topic...)
	seq := append(EncodeVarint(int(time.Now().Unix())), topic1...)

	a, b, c := PubDecode(seq)
	t.Log(a, "-", b, "-", string(c))
}
