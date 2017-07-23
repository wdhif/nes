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
	MagicNumber  uint32  // iNES Magic Number
	ProgramROM   byte    // Number of 16 KB PRG-ROM banks
	CharacterROM byte    // Number of 8 KB CHR-ROM banks
	ControlByte1 byte    // ROM Control Byte 1
	ControlByte2 byte    // ROM Control Byte 2
	RAM          byte    // Number of 8 KB RAM banks
	_            [7]byte // Unused, should all be 0
}

func Loader() int {
	fmt.Println("Loading NES ROM")
	// nestest from http://nickmass.com/images/nestest.nes
	file, err := os.Open("nestest.nes")
	if err != nil {
		fmt.Println("File open error")
		return 1
	}
	defer file.Close()
	romHeader := iNESFileHeader{}
	binary.Read(file, binary.LittleEndian, &romHeader)
	fmt.Println(romHeader)
	fmt.Println(romHeader.MagicNumber)
	fmt.Println(iNESMagicNumber)
	// Check valid Magic Number against rom header
	if romHeader.MagicNumber != 0x1a53454e {
		panic("ROM is invalide: Invalid Magic Number")
	}
	fmt.Println("ROM is valid")
	return 0
}
