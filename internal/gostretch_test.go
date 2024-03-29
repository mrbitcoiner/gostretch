package internal

import (
	"bytes"
	"testing"
)

func TestStretch(t *testing.T) {
	payload := "123"
	diffV := [...]uintptr{
		LOW, MID, HIGH, EXTREME,
	}
	expectedV := [...][]byte{
		{
			0x20,0xce,0x66,0x44,0x94,0x35,0x87,0x2d,
			0x8f,0x60,0xb7,0x4a,0x94,0xfc,0x8d,0x65,
			0xc4,0x78,0x2b,0x0e,0xe4,0x48,0x72,0x69,
			0x47,0x83,0xb8,0x35,0xd0,0x63,0x70,0xa9,
			0xdb,0x58,0x9d,0xe3,0xdc,0x65,0x31,0x85,
			0x80,0x93,0xa9,0xfd,0x59,0x81,0x52,0x6f,
			0x36,0xd2,0xb9,0x6f,0x07,0x91,0x3c,0xa1,
			0x71,0x92,0xab,0x27,0x5a,0x20,0x91,0x97,
		},
		{
			0x2c,0x52,0xbc,0xda,0x18,0xe6,0xee,0x22,
			0x41,0x1a,0xe1,0x81,0x93,0xa6,0xa4,0x6e,
			0xbd,0x2e,0x9d,0xc2,0x2d,0x64,0xdf,0x6b,
			0x1e,0xd0,0xec,0xd2,0xa0,0xcf,0x91,0x1b,
			0xcd,0x61,0x45,0xe8,0x0a,0xea,0x8d,0x02,
			0x7f,0xc1,0xf6,0xad,0xf4,0xca,0x81,0xe2,
			0xee,0x7d,0xf2,0x4d,0x7a,0x9f,0x37,0x58,
			0x8e,0x71,0xb3,0xad,0xe0,0xbc,0x20,0x76,
		},
		{
			0xf5,0xff,0x7c,0x2c,0xc4,0x39,0x7d,0x5c,
			0x39,0x12,0x50,0xff,0x89,0xdd,0xe1,0x15,
			0xe7,0xe4,0x88,0x6b,0x15,0x5e,0x6d,0xff,
			0x53,0x05,0xc3,0xae,0xd5,0x65,0x60,0x1d,
			0x6d,0x04,0x54,0x98,0xcf,0xb2,0x84,0xce,
			0xf1,0x9d,0x23,0x29,0xdc,0xc0,0xa4,0xa9,
			0xdf,0x5d,0x39,0x3e,0x12,0xea,0x64,0x6c,
			0x16,0x03,0x7c,0x76,0x9d,0xbf,0xb2,0x0e,
		},
		{
			0xa7,0x78,0x38,0xdd,0xc4,0x48,0xfc,0xef,
			0x15,0x13,0x02,0xc4,0xbd,0xa1,0x39,0x1f,
			0xdd,0x47,0xf3,0xb2,0xb8,0x89,0x9d,0xbd,
			0x8c,0x80,0x78,0x38,0x5c,0x9f,0x98,0x4f,
			0x80,0x0d,0xac,0x1a,0xf0,0x9a,0x38,0x1a,
			0x9b,0x9a,0x1e,0xd0,0x73,0x62,0x58,0x39,
			0xa2,0xe7,0x13,0x5a,0x12,0x0c,0xc6,0xac,
			0xe9,0x50,0x20,0xd1,0x43,0xba,0x7a,0x05,
		},
	}
	buf := make([]byte, 0, 4096 * 1024 * 1024)
	for i, v := range expectedV {
		buf = Stretch(payload, diffV[i], buf)
		if bytes.Compare(v, buf) != 0 {
			t.Fatalf("expected: %2x\ngot: %2x\nat index %d", v, buf, i)
		}
	}
}
