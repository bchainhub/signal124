<?php
header("Content-Type: text/plain");
include "../signal124.php";

$samples = Array(
	"Hell-o World" => "tØàù;Θf£&0ì8E(Å",
	"John Doe" => "J#c+§ΦèCØå",
	"signal01" => "CΩÄ+'6¿£/ù"
);

function TestEncode() {
	global $samples;
	foreach ($samples as $dec => $enc) {
		$encoded = signal124Encode($dec);
		if($encoded != $enc) {
			printf("Incorrect encoding of `%s` got `%s`\n", $enc, $encoded);
		} else {
			printf("Correct encoding of `%s` got `%s`\n", $enc, $encoded);
		}
	}
}

function TestDecode() {
	global $samples;
	foreach ($samples as $dec => $enc) {
		$decoded = signal124Decode($enc);
		if($decoded != $dec) {
			printf("Incorrect decoding of `%s` got `%s`\n", $dec, $decoded);
		} else {
			printf("Correct decoding of `%s` got `%s`\n", $dec, $decoded);
		}
	}
}

TestEncode();
print("\n");
TestDecode();
?>