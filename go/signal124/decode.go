package signal124

func decode(srcBytes []byte) []byte {
	src := []rune(string(srcBytes))
	var b, n uint32 = 0, 0
	v := -1
	decoded := []byte{}

	for i := 0; i < len(src); i++ {
		p := dectab[src[i]]
		if p > 123 {
			continue
		}
		if v < 0 {
			v = int(p)
			continue
		}
		v += int(p) * 124
		b |= uint32(v) << n
		if v&8191 > 88 {
			n += 13
		} else {
			n += 14
		}
		for {
			decoded = append(decoded, byte(b))
			b >>= 8
			n -= 8
			if n <= 7 {
				break
			}
		}
		v = -1
	}

	if v > -1 {
		decoded = append(decoded, byte(b|uint32(v)<<n))
	}
	return decoded
}

func Decode(dst, src []byte) int {
	return copy(dst, decode(src))
}

func DecodeString(s string) []byte {
	return decode([]byte(s))
}
