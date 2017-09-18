package nes

import (
	"encoding/binary"
	"fmt"
	"os"
)

// iNES Magic Number is "NES" followed by MS-DOS end-of-file
// Stored in Little endian
const iNESMagicNumber = 0x1a53454e

type iNESFileHeader struct {
	MagicNumber  uint32  // iNES Magic Number (32 bits)
	ProgramROM   byte    // Number of 16 KB PRG-ROM banks (8 bits)
	CharacterROM byte    // Number of 8 KB CHR-ROM banks (8 bits)
	ControlByte1 byte    // ROM Control Byte 1 (8 bits)
	ControlByte2 byte    // ROM Control Byte 2 (8 bits)
	RAM          byte    // Number of 8 KB RAM banks (8 bits)
	_            [7]byte // Unused, should all be 0 (8 bits)
}

// Loader reads an iNES file
func Loader(romPath string) (error) {
	fmt.Println("Loading NES ROM")
	file, err := os.Open(romPath)
	if err != nil {
		fmt.Println("File open error")
		return nil
	}
	defer file.Close()

	romHeader := iNESFileHeader{}
	binary.Read(file, binary.LittleEndian, &romHeader)
	fmt.Println(romHeader)

	// Check valid Magic Number against rom header
	if romHeader.MagicNumber != iNESMagicNumber {
		panic("ROM is invalid: Invalid Magic Number")
	}
	fmt.Println("ROM is valid")

	// Mapper type
	mapperLowerBits := romHeader.ControlByte1 >> 4
	mapperHigherBits := romHeader.ControlByte2 >> 4
	mapper := mapperHigherBits | mapperLowerBits << 1

	fmt.Println(fmt.Sprintf("%b", romHeader))
	fmt.Println(mapper)
	return nil
}
