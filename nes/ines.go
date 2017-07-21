package nes

import (
	"fmt"
)

// iNES Magic Number is "NES" followed by MS-DOS end-of-file
// Stored in Little endian
const iNESMagicNumber = 0x1a53454e

type iNESFileHeader struct {
	MagicNumber  uint32  // iNES Magic Number
	ProgramROM   byte    // Number of 16 KB PRG-ROM banks
	CharacterROM byte    // Number of 8 KB CHR-ROM banks
	ControlByte1 byte    // ROM Control Byte 1
	ControlByte2 byte    // ROM Control Byte 2
	RAM          byte    // Number of 8 KB RAM banks
	_            [7]byte // Unused, should all be 0
}

func Loader() {
	fmt.Println("Loading NES ROM")
}
