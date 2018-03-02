package nes

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// iNES Magic Number is "NES" followed by MS-DOS end-of-file
// Hexadecimal, stored in Little endian
const iNESMagicNumber = 0x1a53454e

type iNESFileHeader struct {
	MagicNumber  uint32  // iNES Magic Number (32 bits / bytes 0-3)
	ProgramROM   byte    // Number of 16 KB PRG-ROM banks (8 bits / byte 4)
	CharacterROM byte    // Number of 8 KB CHR-ROM banks (8 bits / byte 5)
	ControlByte1 byte    // ROM Control Byte 1 (8 bits / byte 6)
	ControlByte2 byte    // ROM Control Byte 2 (8 bits / byte 7)
	RAM          byte    // Number of 8 KB RAM banks (8 bits / byte 8)
	_            [7]byte // Unused, should all be 0 (8 bits / bytes 9-15)
}

// Loader reads an iNES file and return a ROM
func Loader(romPath string) (string, error) {
	fmt.Println("Loading NES ROM")
	file, err := os.Open(romPath)
	if err != nil {
		return "nil", errors.New("file open error")
	}
	defer file.Close()

	romHeader := iNESFileHeader{}
	binary.Read(file, binary.LittleEndian, &romHeader)

	// Check valid Magic Number against rom header
	if romHeader.MagicNumber != iNESMagicNumber {
		return "nil", errors.New("file is not a valid ROM: Invalid Magic Number")
	}
	fmt.Println("ROM is valid")

	// Mapper type ()
	mapperLowerBits := romHeader.ControlByte1 >> 4
	mapperHigherBits := romHeader.ControlByte2 >> 4
	mapper := mapperHigherBits | mapperLowerBits<<1

	fmt.Println()

	fmt.Println("romHeader:", romHeader)
	fmt.Println("MagicNumber:", romHeader.MagicNumber)
	fmt.Print("ProgramROM: ", romHeader.ProgramROM, "x16k\n")
	fmt.Print("CharacterROM: ", romHeader.CharacterROM, "x8k\n")
	fmt.Println("ControlByte1:", romHeader.ControlByte1)
	fmt.Println("ControlByte2:", romHeader.ControlByte2)
	fmt.Println("RAM:", romHeader.RAM)
	fmt.Println("Mapper:", mapper)

	fmt.Println()

	fmt.Println("romHeader in binary:", fmt.Sprintf("%b", romHeader))
	fmt.Println("MagicNumber in binary:", fmt.Sprintf("%b", romHeader.MagicNumber))
	fmt.Println("ProgramROM in binary:", fmt.Sprintf("%b", romHeader.ProgramROM))
	fmt.Println("CharacterROM in binary:", fmt.Sprintf("%b", romHeader.CharacterROM))
	fmt.Println("ControlByte1 in binary:", fmt.Sprintf("%b", romHeader.ControlByte1))
	fmt.Println("ControlByte2 in binary:", fmt.Sprintf("%b", romHeader.ControlByte2))
	fmt.Println("RAM in binary:", fmt.Sprintf("%b", romHeader.RAM))
	fmt.Println("Mapper:", mapper)
	return "nil", nil
}
