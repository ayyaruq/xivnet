package xivnet_test

import (
	"time"

	"github.com/ff14wed/xivnet/v2"
)

var zlibPacket = []byte{
	0x52, 0x52, 0xa0, 0x41, 0xff, 0x5d, 0x46, 0xe2, 0x7f, 0x2a, 0x64, 0x4d, 0x7b,
	0x99, 0xc4, 0x75, 0xe6, 0xf6, 0x93, 0xda, 0x59, 0x01, 0x00, 0x00, 0x94, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x78, 0x9c, 0xb2, 0x60, 0x60, 0x60, 0x10, 0x3d, 0x1b, 0xcd, 0x0e, 0xc2,
	0x3c, 0x0c, 0x0c, 0x0c, 0x22, 0x0c, 0x4e, 0x8c, 0x4a, 0x0c, 0x0c, 0x0c, 0xf6,
	0x0f, 0x3a, 0x23, 0x18, 0xc0, 0xe0, 0xec, 0x43, 0x10, 0xc9, 0x08, 0xc5, 0x20,
	0xa0, 0xc0, 0xc0, 0xc0, 0x80, 0xa9, 0x2f, 0x0c, 0x45, 0x5f, 0xfb, 0x0f, 0x27,
	0xb0, 0x7a, 0x5f, 0x06, 0x26, 0x06, 0x46, 0x86, 0x7c, 0x06, 0x18, 0x30, 0xc0,
	0xd0, 0x27, 0x8d, 0xa2, 0x8f, 0x09, 0x61, 0x0d, 0x18, 0x78, 0x60, 0xa8, 0xf7,
	0xc5, 0x69, 0x0f, 0xdc, 0x71, 0x20, 0xa0, 0x67, 0x64, 0x61, 0x66, 0x60, 0x61,
	0x64, 0xc9, 0xc0, 0xc0, 0x20, 0xc0, 0xc0, 0xc0, 0x00, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x48, 0xbc, 0x1d, 0x4b,
}

var jsonZlibFrameHeader string = `"52 52 a0 41 ff 5d 46 e2 7f 2a 64 4d 7b 99 c4 75"`
var jsonZlibBlock0Header string = `{"SubjectID":123456789,"CurrentID":123456789,"U1":12,"U2":20,"Opcode":322,"Route":34,"Time":"2017-01-26T03:40:47-08:00","U4":0}`
var jsonZlibBlock0Data string = `"00 00 cd e1 00 00 00 00 01 00 00 00 01 00 00 00 00 00 00 00 20 00 00 00"`
var bytesZlibBlock0 = []byte{
	0x38, 0x00, 0x00, 0x00, // Length
	0x15, 0xCD, 0x5B, 0x07, // SubjectID
	0x15, 0xCD, 0x5B, 0x07, // CurrentID
	0x0C, 0x00, 0x00, 0x00, // U1
	0x14, 0x00, 0x42, 0x01, // U2 and Opcode
	0x22, 0x00, 0x00, 0x00, // Route
	0x3f, 0xe0, 0x89, 0x58, // Time
	0x00, 0x00, 0x00, 0x00, // U4
	0x00, 0x00, 0xcd, 0xe1, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
}

var expectedZlibFrame = xivnet.Frame{
	Header: [16]byte{
		0x52, 0x52, 0xa0, 0x41, 0xff, 0x5d, 0x46, 0xe2,
		0x7f, 0x2a, 0x64, 0x4d, 0x7b, 0x99, 0xc4, 0x75,
	},
	Time:        time.Unix(1485430847, 206000000),
	Length:      148,
	NumBlocks:   4,
	Compression: 0x0101,
	Blocks: []*xivnet.Block{
		&xivnet.Block{
			Length: 56,
			Header: xivnet.BlockHeader{
				SubjectID: 123456789,
				CurrentID: 123456789,
				U1:        12,
				U2:        20,
				Opcode:    0x142,
				Route:     34,
				Time:      time.Unix(0x5889e03f, 0),
				U4:        0,
			},
			Data: xivnet.GenericBlockDataFromBytes([]byte{
				0x00, 0x00, 0xcd, 0xe1, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
			}),
		},
		&xivnet.Block{
			Length: 56,
			Header: xivnet.BlockHeader{
				SubjectID: 123456789,
				CurrentID: 123456789,
				U1:        12,
				U2:        20,
				Opcode:    0x156,
				Route:     34,
				Time:      time.Unix(0x5889e03f, 0),
				U4:        0,
			},
			Data: xivnet.GenericBlockDataFromBytes([]byte{
				0x87, 0xf8, 0x42, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x02, 0x00,
				0x01, 0x00, 0x6f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}),
		},
		&xivnet.Block{
			Length: 48,
			Header: xivnet.BlockHeader{
				SubjectID: 123456789,
				CurrentID: 123456789,
				U1:        12,
				U2:        20,
				Opcode:    0x11b,
				Route:     34,
				Time:      time.Unix(0x5889e03f, 0),
				U4:        0,
			},
			Data: xivnet.GenericBlockDataFromBytes([]byte{
				0x02, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			}),
		},
		&xivnet.Block{
			Length: 72,
			Header: xivnet.BlockHeader{
				SubjectID: 123456789,
				CurrentID: 123456789,
				U1:        12,
				U2:        20,
				Opcode:    0x14d,
				Route:     34,
				Time:      time.Unix(0x5889e03f, 0),
				U4:        0,
			},
			Data: xivnet.GenericBlockDataFromBytes([]byte{
				0x87, 0xf8, 0x42, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x02, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x2e, 0x32, 0x38, 0x36, 0x30, 0x38, 0x32, 0x39, 0x00, 0x00, 0x00,
				0x10, 0x00, 0x00, 0x00,
			}),
		},
	},
}

var nonZlibPacket = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x08, 0x00, 0x00, 0x00, 0x15, 0xCD, 0x5B, 0x07, 0x42, 0xe0, 0x89, 0x58,
}

var expectedNonZlibFrame = xivnet.Frame{
	Header: [16]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	},
	Time:        time.Unix(0, 0),
	Length:      64,
	NumBlocks:   1,
	Compression: 0,
	Blocks: []*xivnet.Block{
		&xivnet.Block{
			Length: 24,
			Data: xivnet.GenericBlockDataFromBytes([]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
				0x15, 0xCD, 0x5B, 0x07, 0x42, 0xe0, 0x89, 0x58,
			}),
		},
	},
}

var invalidHeaderPacket = []byte{
	0x52, 0x52, 0x00, 0x00, 0xff, 0x5d, 0x46, 0xe2, 0x7f, 0x2a, 0x64, 0x4d, 0x7b,
	0x99, 0xc4, 0x75, 0xe6, 0xf6, 0x93, 0xda, 0x59, 0x01, 0x00, 0x00, 0x8a, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00,
}

var zeroBlockPacket = []byte{
	0x52, 0x52, 0xa0, 0x41, 0xff, 0x5d, 0x46, 0xe2,
	0x7f, 0x2a, 0x64, 0x4d, 0x7b, 0x99, 0xc4, 0x75,
	0xcf, 0xa1, 0x01, 0x08, 0x61, 0x01, 0x00, 0x00,
	0x30, 0x00, 0x00, 0x00,
	0x00, 0x00,
	0x00, 0x00,
	0x01, 0x01,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00,
	0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01,
}

var invalidBlockPacket = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x08, 0x00, 0x00, 0x00, 0x15, 0xCD, 0x5B, 0x07, 0x42, 0xe0, 0x89, 0x58,
}
