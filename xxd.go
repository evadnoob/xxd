// Package xxd generates an ASCII dump of binary data in the format of
// the `xxd -g1` command line tool.
package xxd

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Dump converts a binary blob based at address into an array
// of separate lines in the format of `xxd -g`.
func Dump(address int, data []byte) (lines []string) {
	offset := address & 15
	base := address - offset
	index := 0
	for n := len(data); n > 0; offset, base = 0, base+16 {
		parts := []string{fmt.Sprintf("%08x:", base)}
		count := 16 - offset
		if count > n {
			count = n
		}
		ch := make([]byte, 17)
		ch[0] = ' '
		for i := 1; i <= offset; i++ {
			parts = append(parts, "  ")
			ch[i] = byte(' ')
		}
		for i := 0; i < count; i++ {
			c := data[index+i]
			parts = append(parts, fmt.Sprintf("%02x", c))
			if c < 32 || c >= 127 {
				c = byte('.')
			}
			ch[1+i+offset] = c
		}
		for i := offset + count; i < 16; i++ {
			parts = append(parts, "  ")
		}
		parts = append(parts, string(ch[:1+offset+count]))
		lines = append(lines, strings.Join(parts, " "))
		index += count
		n -= count
	}
	return
}

// Print dumps the xxd.Dump() output directly to stdout.
func Print(writer io.Writer, offset int, data []byte) {
	for _, line := range Dump(offset, data) {
		_, _ = writer.Write([]byte(line + "\n"))
	}
}

// Log dumps the xxd.Dump() output directly with log.Print().
func Log(offset int, data []byte) {
	for _, line := range Dump(offset, data) {
		log.Print(line)
	}
}
