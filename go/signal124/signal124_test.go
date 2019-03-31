package signal124

import "testing"

var samples = map[string]rune{
	"Hell-o World": "tØàù;Θf£&0ì8E(Å",
	"John Doe": "J#c+§ΦèCØå",
	"signal01": "CΩÄ+'6¿£/ù",
}

func TestEncode(t *testing.T) {
	for k, v := range samples {
		if val := EncodeToString([]byte(k)); val != v {
			t.Errorf("Incorrect encoding of `%s` got `%s`", k, val)
		} else {
			t.Logf("Correct encoding of `%s` got `%s`", k, val)
		}
	}
}

func TestDecode(t *testing.T) {
	for k, v := range samples {
		if val := DecodeToString([]byte(v)); val != k {
			t.Errorf("Incorrect decoding of `%s` got `%s`", v, val)
		} else {
			t.Logf("Correct decoding of `%s` got `%s`", v, val)
		}
	}
}
