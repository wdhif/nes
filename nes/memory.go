package nes

type Cpu struct {
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

func (cpu *Cpu) NewCpu(memory Memory)

func (cpu *Cpu) GetFlags() byte {
	var flags byte
	flags |= cpu.StatusRegister.Carry << 0
	flags |= cpu.StatusRegister.Zero << 1
	flags |= cpu.StatusRegister.InterruptDisable << 2
	flags |= cpu.StatusRegister.DecimalMode << 3
	flags |= cpu.StatusRegister.BreakCommand << 4
	flags |= cpu.StatusRegister.Overflow << 6
	flags |= cpu.StatusRegister.Negative << 7
	return flags
}

func (cpu *Cpu) SetFlags(flags byte) {
	cpu.StatusRegister.Carry = (flags >> 0) & 1
}

func (cpu *Cpu) Reset() {
	cpu.X = 0
	cpu.Y = 0
	cpu.Accumulator = 0
	cpu.StatusRegister.Carry = 0
	cpu.StatusRegister.Zero = false
	cpu.StatusRegister.InterruptDisable = false
	cpu.StatusRegister.DecimalMode = false
	cpu.StatusRegister.BreakCommand = false
	cpu.StatusRegister.BreakCommand = false
	cpu.StatusRegister.BreakCommand = false
	cpu.StackPointer = 0
}
