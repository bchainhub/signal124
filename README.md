# Signal124
## Binary data to GSM 7-bit transfer

> Signal124 is method focusing to transfer binary data in obsolete communication with the most effective solution.
> Ideal for streaming short binary strings to the network.
> It is effective because 124 GSM accepted chars was used instead of other Base distributions.

The current algorithm has been written for the most usable way for obsolete communication. For minifying data is recommended to use compressor before encoding.


## Alphabet

Signal124 needs 124 characters to represent the encoded binary data in ASCII.
> From the 128 [GSM 03.38 characters](https://en.wikipedia.org/wiki/GSM_03.38), the following three ones have been omitted:

| Char | Explanation | Verdict |
| --- | --- | --- |
| LF | a Line Feed control | excluded |
| ESC | an Escape control | excluded |

> String with special meaning:

| Char | Explanation | Verdict |
| --- | --- | --- |
| SP | a Space character; indicate the use of the encoding | excluded |
| CR | a Carriage Return control; used as a delimiter between messages | excluded |


### The translation table is composed of following characters:

| № | Symbol | № | Symbol | № | Symbol | № | Symbol | № | Symbol | № | Symbol | № | Symbol | № | Symbol |
| ---: | :--- | ---: | :--- | ---: | :--- | ---: | :--- | ---: | :--- | ---: | :--- | ---: | :--- | ---: | :--- |
| 0 | @ | 14 | Δ | ✕ | SP | 44 | 0 | 60 | ¡ | 76 | P | 92 | ¿ | 108 | p |
| 1 | £ | 15 | _ | 29 | ! | 45 | 1 | 61 | A | 77 | Q | 93 | a | 109 | q |
| 2 | $ | 16 | Φ | 30 | " | 46 | 2 | 62 | B | 78 | R | 94 | b | 110 | r |
| 3 | ¥ | 17 | Γ | 31 | # | 47 | 3 | 63 | C | 79 | S | 95 | c | 111 | s |
| 4 | è | 18 | Λ | 32 | ¤ | 48 | 4 | 64 | D | 80 | T | 96 | d | 112 | t |
| 5 | é | 19 | Ω | 33 | % | 49 | 5 | 65 | E | 81 | U | 97 | e | 113 | u |
| 6 | ù | 20 | Π | 34 | & | 50 | 6 | 66 | F | 82 | V | 98 | f | 114 | v |
| 7 | ì | 21 | Ψ | 35 | ' | 51 | 7 | 67 | G | 83 | W | 99 | g | 115 | w |
| 8 | ò | 22 | Σ | 36 | ( | 52 | 8 | 68 | H | 84 | X | 100 | h | 116 | x |
| 9 | Ç | 23 | Θ | 37 | ) | 53 | 9 | 69 | I | 85 | Y | 101 | i | 117 | y |
| ✕ | LF | 24 | Ξ | 38 | * | 54 | : | 70 | J | 86 | Z | 102 | j | 118 | z |
| 10 | Ø | ✕ | ESC | 39 | + | 55 | ; | 71 | K | 87 | Ä | 103 | k | 119 | ä |
| 11 | ø | 25 | Æ | 40 | , | 56 | < | 72 | L | 88 | Ö | 104 | l | 120 | ö |
| ✕ | CR | 26 | æ | 41 | - | 57 | = | 73 | M | 89 | Ñ | 105 | m | 121 | ñ |
| 12 | Å | 27 | ß | 42 | . | 58 | > | 74 | N | 90 | Ü | 106 | n | 122 | ü |
| 13 | å | 28 | É | 43 | / | 59 | ? | 75 | O | 91 | § | 107 | o | 123 | à |


## Example of transaction

> Encoded:

<kbd>&nbsp;tØàù;Θf£&0ì8E(Å<br>
J#c+§ΦèCØå<br>
CΩÄ+'6¿£/ù</kbd>

- a Space character indicates the encoding is used (optional)
- a Carriage return is a delimiter for next data

> Decoded:

space - encoding signal124

<kbd>Hell-o World<br>
John Doe<br>
signal01</kbd>

## Distributions

* [Go](go) WIP
* [PHP](php)

## Project

This is a beta version, which is not meant for a production environment.
The Mathematical model can be changed to get better results.
Please, join the project to release the first stable version.
You can raise [issue](issues) or send [pull](pulls) request.

## Compression 🗜️

Consider using lossless compression before transfer and stream your data. Probably you want to achieve a compromise between fast encoding and data size.

Examples of lossless compression libraries:

* [Brotli](https://github.com/google/brotli) - allows a denser packing than gzip and deflate
* [LZ4](https://github.com/lz4/lz4) - compression and decompression speed
* [Snappy](https://github.com/google/snappy) - aims for very high speeds and reasonable compression

## Developers

#### Signal124
[Rastislav](https://github.com/raisty) @ CRYPTO ■ HUB™

## License

⚖️ [GPL-3.0](LICENSE)

