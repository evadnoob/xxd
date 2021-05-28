package xxd

import (
	"testing"
)

func TestDump(t *testing.T) {
	data := []byte{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 32, 33, 126,
	}
	tests := []struct {
		offset   int
		expected []string
	}{
		{
			offset: 0,
			expected: []string{
				"00000000: 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 20  ............... ",
				"00000010: 21 7e                                            !~",
			},
		},
		{
			offset: 3,
			expected: []string{
				"00000000:          01 02 03 04 05 06 07 08 09 0a 0b 0c 0d     .............",
				"00000010: 0e 0f 20 21 7e                                   .. !~",
			},
		},
		{
			offset: 15,
			expected: []string{
				"00000000:                                              01                 .",
				"00000010: 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 20 21  .............. !",
				"00000020: 7e                                               ~",
			},
		},
		{
			offset: 46,
			expected: []string{
				"00000020:                                           01 02                ..",
				"00000030: 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 20 21 7e  ............. !~",
			},
		},
	}
	for i, ans := range tests {
		dump := Dump(ans.offset, data)
		if a, b := len(dump), len(ans.expected); a != b {
			t.Errorf("test=%d wrong line count: got=%d, want=%d", i, a, b)
		} else {
			for j := 0; j < len(dump); j++ {
				if a, b := dump[j], ans.expected[j]; a != b {
					t.Errorf("test=%d,line=%d got=%q want=%q", i, j, a, b)
				}
			}
		}
	}
}

func ExamplePrint() {
	Print(0x4f, []byte{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 32, 33, 126,
	})
	// Output: 00000040:                                              01                 .
	// 00000050: 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 20 21  .............. !
	// 00000060: 7e                                               ~
}
