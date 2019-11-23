package nes

type Memory interface {
	Read(address uint16) byte
	Read16(address uint16) uint16
	Write(address uint16, value byte)
}

func NewMemory(nes *Nes)