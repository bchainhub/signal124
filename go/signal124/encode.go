package signal124

func encode(src []byte) []byte {
	var b, n uint32 = 0, 0
	encoded := []rune{}

	for i := 0; i < len(src); i++ {
		b |= uint32(src[i]) << n
		n += 8

		if n <= 13 {
			continue
		}

		v := b & 8191
		if v > 88 {
			b >>= 13
			n -= 13
		} else {
			v = b & 16383
			b >>= 14
			n -= 14
		}

		encoded = append(encoded, enctab[v%124], enctab[v/124])
	}

	if n != 0 {
		encoded = append(encoded, enctab[b%124])
		if n > 7 || b > 123 {
			encoded = append(encoded, enctab[b/124])
		}
	}
	return []byte(string(encoded))
}

func Encode(dst, src []byte) int {
	return copy(dst, encode(src))
}

func EncodeString(src []byte) string {
	return string(encode(src))
}
