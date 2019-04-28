package signal124

import (
	"crypto/rand"
	"io"
	"testing"
)

var encodeSpec = map[string]string{
	"Hell-o World": "tØàù;Θf£&0ì8E(Å",
	"John Doe": "J#c+§ΦèCØå",
	"signal01": "CΩÄ+'6¿£/ù",
}

func TestEncode(t *testing.T) {
	for k, v := range encodeSpec {
		if actual := EncodeString([]byte(k)); actual != v {
			t.Fatalf("Incorrect encoding of `%s`, got `%s`", v, actual)
		} else {
			t.Logf("Correct encoding of `%s`, got `%s`", v, actual)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	s := make([]byte, 1024*1024)
	if _, err := io.ReadFull(rand.Reader, s); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeString(s)
	}
}
