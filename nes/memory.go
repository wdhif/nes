package nes

type Memory struct {
	RAM []byte
}

func NewMemory() *Memory{
	memory := Memory{}
	return &memory
}
