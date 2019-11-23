package nes

import "log"

type NES struct {
	Cpu			*CPU
	Memory		*Memory
	Rom			*Rom
}

func NewNES(path string) (*NES, error) {
	rom, err := Loader(path)
	if err != nil {
		log.Fatal(err)
	}

	nes := NES{NewCpu(), NewMemory(), rom}
	return &nes, nil
}
