<?php
$encTab = array(
	'@',	'£',	'$',	'¥',	'è',	'é',	'ù',	'ì',	'ò',	'Ç',	'Ø',	'ø',	'Å',	'å',
	'Δ',	'_',	'Φ',	'Γ',	'Λ',	'Ω',	'Π',	'Ψ',	'Σ',	'Θ',	'Ξ',	'Æ',	'æ',	'ß',	'É',
	'!',	'"',	'#',	'¤',	'%',	'&',	'\'',	'(',	')',	'*',	'+',	',',	'-',	'.',	'/',
	'0',	'1',	'2',	'3',	'4',	'5',	'6',	'7',	'8',	'9',	':',	';',	'<',	'=',	'>',	'?',
	'¡',	'A',	'B',	'C',	'D',	'E',	'F',	'G',	'H',	'I',	'J',	'K',	'L',	'M',	'N',	'O',
	'P',	'Q',	'R',	'S',	'T',	'U',	'V',	'W',	'X',	'Y',	'Z',	'Ä',	'Ö',	'Ñ',	'Ü',	'§',
	'¿',	'a',	'b',	'c',	'd',	'e',	'f',	'g',	'h',	'i',	'j',	'k',	'l',	'm',	'n',	'o',
	'p',	'q',	'r',	's',	't',	'u',	'v',	'w',	'x',	'y',	'z',	'ä',	'ö',	'ñ',	'ü',	'à'
);
$decTab = array_flip($encTab);

function signal124Encode($str)
{
	global $encTab;
	$n = 0;
	$byte = $out = false;
	$length = strlen($str);
	for ($ite = 0; $ite < $length; ++$ite) {
		$byte |= mb_ord($str{$ite}) << $n;
		$n += 8;
		if ($n > 13) {
			$v = $byte & 8191;
			if ($v > 88) {
				$byte >>= 13;
				$n -= 13;
			} else {
				$v = $byte & 16383;
				$byte >>= 14;
				$n -= 14;
			}
			$out .= $encTab[$v % 124] . $encTab[$v / 124];
		}
	}
	if ($n) {
		$out .= $encTab[$byte % 124];
		if ($n > 7 || $byte > 123)
			$out .= $encTab[$byte / 124];
	}
	return $out;
}

function signal124Decode($str)
{
	global $decTab;
	$n = 0;
	$byte = $out = false;
	$length = mb_strlen($str, 'utf8');
	$v = -1;
	for ($ite = 0; $ite < $length; ++$ite) {
		$char = $decTab[mb_substr($str, $ite, 1)];
		if (!isset($char))
			continue;
		if ($v < 0)
			$v = $char;
		else {
			$v += $char * 124;
			$byte |= $v << $n;
			$n += ($v & 8191) > 88 ? 13 : 14;
			do {
				$out .= mb_chr($byte & 255);
				$byte >>= 8;
				$n -= 8;
			} while ($n > 7);
			$v = -1;
		}
	}
	if ($v + 1)
		$out .= mb_chr(($byte | $v << $n) & 255);
	return $out;
}
?>
