package nes

type Cpu struct {
	ProgramCounter uint16   // Program Counter (16 bits)
	StackPointer   byte     // Stack Pointer (8 bits)
	Accumulator    byte     // Accumulator Register (8 bits)
	X              byte     // X Index Register (8 bits)
	Y              byte     // Y Index Register (8 bits)
	StatusRegister struct { // Status Register (8 bits)
		Carry            bool // Bit 0
		Zero             bool // Bit 1
		InterruptDisable bool // Bit 2
		DecimalMode      bool // Bit 3
		BreakCommand     bool // Bit 4
		_                bool // Bit 5 - Unused
		Overflow         bool // Bit 6
		Negative         bool // Bit 7
	}
}

func (cpu *Cpu) Reset() {
	cpu.X = 0
	cpu.Y = 0
	cpu.Accumulator = 0
	cpu.StatusRegister.Carry = false
	cpu.StatusRegister.Zero = false
	cpu.StatusRegister.InterruptDisable = false
	cpu.StatusRegister.DecimalMode = false
	cpu.StatusRegister.BreakCommand = false
	cpu.StatusRegister.BreakCommand = false
	cpu.StatusRegister.BreakCommand = false
	cpu.StackPointer = 0
}
