package calculator

//小文字と大文字の違いが外部パッケージで使用できるか否か

var offset float64 = 1
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	return a + b + offset
}
