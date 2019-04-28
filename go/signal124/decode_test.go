package signal124

import (
	"crypto/rand"
	"testing"
)

var decodeSpec = map[string]string{
	"tØàù;Θf£&0ì8E(Å": "Hell-o World",
	"J#c+§ΦèCØå": "John Doe",
	"CΩÄ+'6¿£/ù": "signal01",
}

func TestDecode(t *testing.T) {
	for k, v := range decodeSpec {
		if actual := string(DecodeString(k)); actual != v {
			t.Fatalf("Incorrect decoding of `%s`, got `%s`", v, actual)
		} else {
			t.Logf("Correct decoding of `%s`, got `%s`", v, actual)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	s := make([]byte, 1024*1024)
	if _, err := rand.Read(s); err != nil {
		b.Fatal(err)
	}

	encoded := EncodeString(s)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeString(encoded)
	}
}
