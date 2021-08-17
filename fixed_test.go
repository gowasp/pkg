package pkg

import (
	"testing"
)

func TestPubDecode(t *testing.T) {
	topic := append([]byte("a/b/c"), []byte("pact")...)
	topic1 := append([]byte{byte(len([]byte("a/b/c")))}, topic...)

	a, b := PubDecode(topic1)
	t.Log(string(a), "-", b)
}
