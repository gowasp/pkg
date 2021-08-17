package pkg

import (
	"fmt"
	"testing"
	"time"
)

func TestPubDecode(t *testing.T) {
	topic := append([]byte("a/b/c"), []byte("pact")...)
	topic1 := append([]byte{byte(len([]byte("a/b/c")))}, topic...)

	a, b := PubDecode(topic1)
	t.Log(string(a), "-", b)
}

func TestPubDecodeSeq(t *testing.T) {
	body := append(EncodeVarint(int(time.Now().Unix())), []byte("pkgtest")...)
	topic := append([]byte("a/b/c"), body...)
	topic1 := append([]byte{byte(len([]byte("a/b/c")))}, topic...)

	a, b, c := PubDecodeSeq(topic1)
	t.Log(fmt.Sprint(a), "-", b, "-", string(c))
}
