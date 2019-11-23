package nes

type CPU struct {
	ProgramCounter uint16   // Program Counter (16 bits)
	StackPointer   byte     // Stack Pointer (8 bits)
	Accumulator    byte     // Accumulator Register (8 bits)
	X              byte     // X Index Register (8 bits)
	Y              byte     // Y Index Register (8 bits)
	StatusRegister struct { // Status Register (8 bits)
		Carry            byte // Bit 0
		Zero             byte // Bit 1
		InterruptDisable byte // Bit 2
		DecimalMode      byte // Bit 3
		BreakCommand     byte // Bit 4
		_                byte // Bit 5 - Unused
		Overflow         byte // Bit 6
		Negative         byte // Bit 7
	}
}

func NewCpu() *CPU{
	cpu := CPU{}
	cpu.Reset()
	return &cpu
}

func (cpu *CPU) Reset() {
	cpu.X = 0
	cpu.Y = 0
	cpu.Accumulator = 0
	cpu.StatusRegister.Carry = 0
	cpu.StatusRegister.Zero = 0
	cpu.StatusRegister.InterruptDisable = 0
	cpu.StatusRegister.DecimalMode = 0
	cpu.StatusRegister.BreakCommand = 0
	cpu.StatusRegister.BreakCommand = 0
	cpu.StatusRegister.BreakCommand = 0
	cpu.StackPointer = 0
}
