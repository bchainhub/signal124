package signal124

var enctab = []byte("@£$¥èéùìòÇØøÅåΔ_ΦΓΛΩΠΨΣΘΞÆæßÉ!\"#¤%&',()*+,-./0123456789:;<=>?¡ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÑÜ§¿abcdefghijklmnopqrstuvwxyzäöñüà")
var dectab = map[rune]byte{
	'@': 0,		'£': 1,		'$': 2,		'¥': 3,		'è': 4,		'é': 5,		'ù': 6,		'ì': 7,		'ò': 8,		'Ç': 9,		'Ø': 10,	'ø': 11,	'Å': 12,	'å': 13,
	'Δ': 14,	'_': 15,	'Φ': 16,	'Γ': 17,	'Λ': 18,	'Ω': 19,	'Π': 20,	'Ψ': 21,	'Σ': 22,	'Θ': 23,	'Ξ': 24,	'Æ': 25,	'æ': 26,	'ß': 27,	'É': 28,
	'!': 29,	'"': 30,	'#': 31,	'¤': 32,	'%': 33,	'&': 34,	'\'': 35,	'(': 36,	')': 37,	'*': 38,	'+': 39,	',': 40,	'-': 41,	'.': 42,	'/': 43,
	'0': 44,	'1': 45,	'2': 46,	'3': 47,	'4': 48,	'5': 49,	'6': 50,	'7': 51,	'8': 52,	'9': 53,	':': 54,	';': 55,	'<': 56,	'=': 57,	'>': 58,	'?': 59,
	'¡': 60,	'A': 61,	'B': 62,	'C': 63,	'D': 64,	'E': 65,	'F': 66,	'G': 67,	'H': 68,	'I': 69,	'J': 70,	'K': 71,	'L': 72,	'M': 73,	'N': 74,	'O': 75,
	'P': 76,	'Q': 77,	'R': 78,	'S': 79,	'T': 80,	'U': 81,	'V': 82,	'W': 83,	'X': 84,	'Y': 85,	'Z': 86,	'Ä': 87,	'Ö': 88,	'Ñ': 89,	'Ü': 90,	'§': 91,
	'¿': 92,	'a': 93,	'b': 94,	'c': 95,	'd': 96,	'e': 97,	'f': 98,	'g': 99,	'h': 100,	'i': 101,	'j': 102,	'k': 103,	'l': 104,	'm': 105,	'n': 106,	'o': 107,
	'p': 108,	'q': 109,	'r': 110,	's': 111,	't': 112,	'u': 113,	'v': 114,	'w': 115,	'x': 116,	'y': 117,	'z': 118,	'ä': 119,	'ö': 120,	'ñ': 121,	'ü': 122,	'à': 123,
}

func Encode(d []byte) []byte {
	var n, b uint
	var o []byte

	for i := 0; i < len(d); i++ {
		b |= uint(d[i]) << n
		n += 8
		if n > 13 {
			v := b & 8191
			if v > 88 {
				b >>= 13
				n -= 13
			} else {
				v = b & 16383
				b >>= 14
				n -= 14
			}
			o = append(o, enctab[v%124], enctab[v/124])
		}
	}
	if n > 0 {
		o = append(o, enctab[b%124])
		if n > 7 || b > 123 {
			o = append(o, enctab[b/124])
		}
	}
	return o
}

func Decode(d []rune) []byte {
	var b, n uint
	var o []byte
	v := -1

	for i := 0; i < len(d); i++ {
		c, ok := dectab[d[i]]
		if !ok {
			continue
		}
		if v < 0 {
			v = int(c)
		} else {
			v += int(c) * 124
			b |= uint(v) << n
			if v&8191 > 88 {
				n += 13
			} else {
				n += 14
			}
			o = append(o, byte(b&255))
			b >>= 8
			n -= 8
			for n > 7 {
				o = append(o, byte(b&255))
				b >>= 8
				n -= 8
			}
			v = -1
		}
	}
	if v+1 > 0 {
		o = append(o, byte((b|uint(v)<<n)&255))
	}
	return o
}

func EncodeToString(d []byte) string {
	return string(Encode(d))
}

func DecodeToString(d []rune) string {
	return string(Decode(d))
}
