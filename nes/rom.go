package nes

type Rom struct {
	ProgramRom   []byte
	CharacterROM []byte
	Mapper       int
	Mirror       int
	Battery      bool
}
