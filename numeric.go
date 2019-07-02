package numeric

import (
	"fmt"
	"strconv"
	"strings"
)

func NumericFormatString(n float64) string {
	ns := strconv.FormatFloat(n, 'f', -1, 64)

	afterPoint := ""
	dotIndex := strings.Index(ns, ".")
	if dotIndex > -1 {
		afterPoint = ns[dotIndex:]
		ns = ns[:dotIndex]
	}

	negativeSign := ""
	if string(ns[0]) == "-" {
		negativeSign = "-"
		ns = ns[1:]
	}

	nsLength := len([]rune(ns)) - 1
	result := ""
	flag := nsLength

	for i := nsLength - 3; i >= 0; i -= 3 {
		flag = i
		result = fmt.Sprintf(",%s%s", ns[i+1:i+4], result)
	}
	result = fmt.Sprintf("%s%s%s%s", negativeSign, ns[:flag+1], result, afterPoint)

	return result
}
