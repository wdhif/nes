package nes

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

// iNES Magic Number is "NES" followed by MS-DOS end-of-file.
// Hexadecimal, stored in Little endian.
const iNESMagicNumber = 0x1a53454e

type iNESFileHeader struct {
	MagicNumber             uint32  // iNES Magic Number (32 bits / bytes 0-3)
	ProgramROMBanksNumber   byte    // Number of 16 KB PRG-ROM banks (8 bits / byte 4)
	CharacterROMBanksNumber byte    // Number of 8 KB CHR-ROM banks (8 bits / byte 5)
	ControlByte1            byte    // ROM Control Byte 1 (Flag 6) (8 bits / byte 6)
	ControlByte2            byte    // ROM Control Byte 2 (Flag 7) (8 bits / byte 7)
	RAM                     byte    // Number of 8 KB RAM banks (8 bits / byte 8)
	_                       [7]byte // Unused, should all be 0 (8 bits / bytes 9-15)
}

// Loader reads an iNES file and return a ROM.
func Loader(romPath string) (*Rom, error) {

	fmt.Print("Loading NES ROM... ")
	file, err := os.Open(romPath)
	if err != nil {
		return nil, errors.New("file open error")
	}
	defer file.Close()

	// Read rom header.
	romHeader := iNESFileHeader{}
	binary.Read(file, binary.LittleEndian, &romHeader)

	// Check valid Magic Number against rom header
	if romHeader.MagicNumber != iNESMagicNumber {
		return nil, errors.New("file is not a valid ROM: Invalid Magic Number.")
	}
	fmt.Println("ROM is valid")

	// Check iNES format (NES 2.0)
	if romHeader.ControlByte2&0x0c == 0x08 {
		return nil, errors.New("NES 2.0 format is not implemented.")
	}

	// ProgramRom
	ProgramRom := make([]byte, int(romHeader.ProgramROMBanksNumber)*16384)
	if _, err := io.ReadFull(file, ProgramRom); err != nil {
		return nil, err
	}

	// CharacterRom
	CharacterRom := make([]byte, int(romHeader.CharacterROMBanksNumber)*8192)
	if _, err := io.ReadFull(file, CharacterRom); err != nil {
		return nil, err
	}

	// Mapper
	mapperLowerBits := romHeader.ControlByte1 >> 4
	mapperHigherBits := romHeader.ControlByte2 >> 4
	Mapper := mapperHigherBits | mapperLowerBits<<1

	// Mirroring
	Mirror1 := int(romHeader.ControlByte1) & 1
	Mirror2 := int(romHeader.ControlByte1>>3) & 1
	Mirror := Mirror1 | Mirror2<<1

	// Battery
	Battery := romHeader.ControlByte1&2 == 2

	rom := Rom{
		ProgramRom:   ProgramRom,
		CharacterROM: CharacterRom,
		Mapper:       int(Mapper),
		Mirror:       Mirror,
		Battery:      Battery,
	}

	return &rom, nil
}
