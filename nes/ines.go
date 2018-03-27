package nes

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

// iNES Magic Number is "NES" followed by MS-DOS end-of-file
// Hexadecimal, stored in Little endian
const iNESMagicNumber = 0x1a53454e

type iNESFileHeader struct {
	MagicNumber             uint32  // iNES Magic Number (32 bits / bytes 0-3)
	ProgramROMBanksNumber   byte    // Number of 16 KB PRG-ROM banks (8 bits / byte 4)
	CharacterROMBanksNumber byte    // Number of 8 KB CHR-ROM banks (8 bits / byte 5)
	ControlByte1            byte    // ROM Control Byte 1 (8 bits / byte 6)
	ControlByte2            byte    // ROM Control Byte 2 (8 bits / byte 7)
	RAM                     byte    // Number of 8 KB RAM banks (8 bits / byte 8)
	_                       [7]byte // Unused, should all be 0 (8 bits / bytes 9-15)
}

// Loader reads an iNES file and return a ROM
func Loader(romPath string) (string, error) {

	DEBUG := false

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

	// Mirroring type
	Mirror1 := int(romHeader.ControlByte1) & 1
	Mirror2 := int(romHeader.ControlByte1>>3) & 1
	Mirror := Mirror1 | Mirror2<<1

	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "ProgramROM Size:\t", romHeader.ProgramROMBanksNumber*16, "\tKb")
	fmt.Fprintln(w, "CharacterROM Size:\t", romHeader.CharacterROMBanksNumber*8, "\tKb")
	if Mirror == 0 {
		fmt.Fprintln(w, "Mirror Type:\t", "Horizontal")
	} else if Mirror == 1 {
		fmt.Fprintln(w, "Mirror Type:\t", "Vertical")
	}
	fmt.Fprintln(w, "Mapper Number:\t", mapper)
	w.Flush()

	if DEBUG == true {
		fmt.Println()

		fmt.Fprintln(w, "romHeader:\t", romHeader)
		fmt.Fprintln(w, "MagicNumber:\t", romHeader.MagicNumber)
		fmt.Fprintln(w, "ProgramROM:\t", romHeader.ProgramROMBanksNumber, "\tx16k")
		fmt.Fprintln(w, "CharacterROM:\t", romHeader.CharacterROMBanksNumber, "x8k")
		fmt.Fprintln(w, "ControlByte1:\t", romHeader.ControlByte1)
		fmt.Fprintln(w, "ControlByte2:\t", romHeader.ControlByte2)
		fmt.Fprintln(w, "mapperLowerBits:\t", mapperLowerBits)
		fmt.Fprintln(w, "mapperHigherBits:\t", mapperHigherBits)
		fmt.Fprintln(w, "mapperLowerBits:\t", Mirror1)
		fmt.Fprintln(w, "mapperHigherBits:\t", Mirror2)
		fmt.Fprintln(w, "Mirror:\t", Mirror)
		fmt.Fprintln(w, "RAM:\t", romHeader.RAM)
		fmt.Fprintln(w, "Mapper:\t", mapper)
		w.Flush()

		fmt.Println()

		fmt.Fprintln(w, "romHeader in binary:\t", fmt.Sprintf("%08b", romHeader))
		fmt.Fprintln(w, "MagicNumber in binary:\t", fmt.Sprintf("%08b", romHeader.MagicNumber))
		fmt.Fprintln(w, "ProgramROM in binary:\t", fmt.Sprintf("%08b", romHeader.ProgramROMBanksNumber))
		fmt.Fprintln(w, "CharacterROM in binary:\t", fmt.Sprintf("%08b", romHeader.CharacterROMBanksNumber))
		fmt.Fprintln(w, "ControlByte1 in binary:\t", fmt.Sprintf("%08b", romHeader.ControlByte1))
		fmt.Fprintln(w, "ControlByte2 in binary:\t", fmt.Sprintf("%08b", romHeader.ControlByte2))
		fmt.Fprintln(w, "mapperLowerBits in binary:\t", fmt.Sprintf("%08b", mapperLowerBits))
		fmt.Fprintln(w, "mapperHigherBits in binary:\t", fmt.Sprintf("%08b", mapperHigherBits))
		fmt.Fprintln(w, "Mirror1 in binary:\t", fmt.Sprintf("%08b", Mirror1))
		fmt.Fprintln(w, "Mirror2 in binary:\t", fmt.Sprintf("%08b", Mirror2))
		fmt.Fprintln(w, "Mirror in binary:\t", fmt.Sprintf("%08b", Mirror))
		fmt.Fprintln(w, "RAM in binary:\t", fmt.Sprintf("%08b", romHeader.RAM))
		fmt.Fprintln(w, "Mapper in binary:\t", fmt.Sprintf("%08b", mapper))
		w.Flush()
	}

	return "nil", nil
}
