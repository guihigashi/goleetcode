package n0006

import (
	"bytes"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]bytes.Buffer, numRows)

	k := 0
	direction := -1

	for _, c := range s {
		rows[k].WriteRune(c)
		if k == 0 || k == numRows-1 {
			direction *= -1
		}
		k += direction
	}

	var sb strings.Builder

	for _, row := range rows {
		sb.Write(row.Bytes())
	}

	return sb.String()
}
