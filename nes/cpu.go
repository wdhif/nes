package nes

import "math/bits"

const (
	RAM_SIZE = 0xffff
	STACK_INDEX = 0x01ff
)

const (
	Carry            byte = (1 << 0)
	Zero             byte = (1 << 1)
	InterruptDisable byte = (1 << 2)
	DecimalMode      byte = (1 << 3)
	BreakCommand     byte = (1 << 4)
	//_              (Bit 5 - Unused)
	Overflow         byte = (1 << 6)
	Negative         byte = (1 << 7)
)

type CPU struct {
	PC     uint16   // Program Counter (16 bits)
	A      byte     // Accumulator Register (8 bits)
	X      byte     // X Index Register (8 bits)
	Y      byte     // Y Index Register (8 bits)
	S      byte     // Stack Pointer (8 bits)
	P      byte     // Processor Status (8 bits)
	Memory [RAM_SIZE]byte
}

func NewCpu() *CPU {
	cpu := CPU{}
	cpu.Reset()
	return &cpu
}


func (cpu *CPU) Reset() {
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.S = 0
	cpu.P = 0
}

// Private functions
func (cpu *CPU) fromAddress(address uint16) byte {
	return cpu.Memory[address]
}

func (cpu *CPU) push(value byte) {
	index := STACK_INDEX - uint16(cpu.S)
	cpu.Memory[index] = value
	cpu.S += 1
}

func (cpu *CPU) push16(value uint16) {
	upper_byte := byte(value >> 4)
	lower_byte := byte(value & 0x0f)
	cpu.push(upper_byte)
	cpu.push(lower_byte)
}

func (cpu *CPU) pop() byte {
	cpu.S -= 1
	index := STACK_INDEX - uint16(cpu.S)
	return cpu.Memory[index]
}

func (cpu *CPU) pop16() uint16 {
	lower_byte := uint16(cpu.pop())
	upper_byte := uint16(cpu.pop())
	return (upper_byte << 4) | lower_byte
}

func (cpu *CPU) setFlag(flag byte, on bool) {
	if (on) {
		// Toggle "on"
		cpu.P |= flag
	} else {
		// Toggle "off"
		cpu.P |= flag
		cpu.P ^= bits.Reverse8(flag)
	}
}

func (cpu *CPU) cmp(lhs byte, rhs byte) {
	// Reset carry and zero flags
	cpu.setFlag(Carry, lhs >= rhs)
	cpu.setFlag(Zero, lhs == rhs)
}

func (cpu *CPU) branchIf(address uint16, condition bool) {
	if (condition) {
		cpu.Jmp(address)
	}
}
// End of private functions

// Load into registers
func (cpu *CPU) Lda(value byte) {
	cpu.A = value
}
func (cpu *CPU) Ldx(value byte) {
	cpu.X = value
}
func (cpu *CPU) Ldy(value byte) {
	cpu.Y = value
}
// End

// Math functions
func (cpu *CPU) Inc(address uint16) {
	cpu.Memory[address] += 1
}
func (cpu *CPU) Dec(address uint16) {
	cpu.Memory[address] -= 1
}
func (cpu *CPU) Inx() {
	cpu.X += 1
}
func (cpu *CPU) Iny() {
	cpu.Y += 1
}
func (cpu *CPU) Dex() {
	cpu.X -= 1
}
func (cpu *CPU) Dey() {
	cpu.Y -= 1
}
func (cpu *CPU) Adc(value byte) {
	cpu.setFlag(Carry, uint16(cpu.A) + uint16(value) > 0xff)
	cpu.A += value
}
func (cpu *CPU) Sbc(value byte) {
	cpu.setFlag(Carry, value > cpu.A)
	cpu.A -= value
}
// End

// Bitwise functions
func (cpu *CPU) And(address uint16) {
	cpu.A &= cpu.Memory[address]
}
func (cpu *CPU) Ora(address uint16) {
	cpu.A |= cpu.Memory[address]
}
func (cpu *CPU) Eor(address uint16) {
	cpu.A ^= cpu.Memory[address]
}
func (cpu *CPU) Bit(address uint16) {
	cpu.A &= cpu.Memory[address]
	cpu.setFlag(Negative, (cpu.Memory[address] & 128) == 1)
	cpu.setFlag(Overflow, (cpu.Memory[address] & 64) == 1)
	cpu.setFlag(Zero, cpu.A == 0)
}
func (cpu *CPU) Lsr() {
	cpu.setFlag(Carry, false)
	cpu.P |= cpu.A & 1
	cpu.A = cpu.A >> 1
}
func (cpu *CPU) Asl() {
	cpu.setFlag(Overflow, (cpu.A & 128) == 1)
	cpu.A = cpu.A << 1
}
func (cpu *CPU) Rol() {
	carryFlag := cpu.P & Carry
	oldBit7 := cpu.A & 128
	cpu.A = cpu.A << 1
	if (carryFlag != 0) {
		cpu.A |= 1
	}
	cpu.setFlag(Carry, oldBit7 != 0)
}
func (cpu *CPU) Ror() {
	carryFlag := cpu.P & Carry
	oldBit0 := cpu.A & 1
	cpu.A = cpu.A >> 1
	if (carryFlag != 0) {
		cpu.A |= 128
	}
	cpu.setFlag(Carry, oldBit0 != 0)
}
// End

// P register functions
func (cpu *CPU) Clc() {
	cpu.setFlag(Carry, false)
}
func (cpu *CPU) Cld() {
	cpu.setFlag(DecimalMode, false)
}
func (cpu *CPU) Cli() {
	cpu.setFlag(InterruptDisable, false)
}
func (cpu *CPU) Clv() {
	cpu.setFlag(Overflow, false)
}
func (cpu *CPU) Sec() {
	cpu.setFlag(Carry, true)
}
func (cpu *CPU) Sed() {
	cpu.setFlag(DecimalMode, true)
}
func (cpu *CPU) Sei() {
	cpu.setFlag(InterruptDisable, true)
}
// End

// Registers transfers
func (cpu *CPU) Tax() {
	cpu.X = cpu.A
}
func (cpu *CPU) Tay() {
	cpu.Y = cpu.A
}
func (cpu *CPU) Txa() {
	cpu.A = cpu.X
}
func (cpu *CPU) Tya() {
	cpu.A = cpu.Y
}
func (cpu *CPU) Txs() {
	cpu.S = cpu.X
}
func (cpu *CPU) Tsx() {
	cpu.X = cpu.S
}
// End

// Store registers
func (cpu *CPU) Sta(address uint16) {
	cpu.Memory[address] = cpu.A
}
func (cpu *CPU) Stx(address uint16) {
	cpu.Memory[address] = cpu.X
}
func (cpu *CPU) Sty(address uint16) {
	cpu.Memory[address] = cpu.Y
}
// End

// Stack
func (cpu *CPU) Pha() {
	cpu.push(cpu.A)
}
func (cpu *CPU) Pla() {
	cpu.A = cpu.pop()
}
func (cpu *CPU) Php() {
	cpu.push(cpu.P)
}
func (cpu *CPU) Plp() {
	cpu.P = cpu.pop()
}
// End

// Comparisons
func (cpu *CPU) Cmp(address uint16) {
	value := cpu.Memory[address]
	cpu.cmp(cpu.A, value)
}
func (cpu *CPU) Cpx(address uint16) {
	value := cpu.Memory[address]
	cpu.cmp(cpu.X, value)
}
func (cpu *CPU) Cpy(address uint16) {
	value := cpu.Memory[address]
	cpu.cmp(cpu.Y, value)
}
// End

// Jumps
func (cpu *CPU) Jmp(address uint16) {
	cpu.PC = address
}
func (cpu *CPU) Jsr(address uint16) {
	cpu.push16(cpu.PC - 1)
	cpu.Jmp(address)
}
func (cpu *CPU) Beq(address uint16) {
	cpu.branchIf(address, (cpu.P & Zero) != 0)
}
func (cpu *CPU) Bne(address uint16) {
	cpu.branchIf(address, (cpu.P & Zero) == 0)
}
func (cpu *CPU) Bmi(address uint16) {
	cpu.branchIf(address, (cpu.P & Negative) != 0)
}
func (cpu *CPU) Bpl(address uint16) {
	cpu.branchIf(address, (cpu.P & Negative) == 0)
}
func (cpu *CPU) Bcc(address uint16) {
	cpu.branchIf(address, (cpu.P & Carry) == 0)
}
func (cpu *CPU) Bcs(address uint16) {
	cpu.branchIf(address, (cpu.P & Carry) != 0)
}
func (cpu *CPU) Bvc(address uint16) {
	cpu.branchIf(address, (cpu.P & Overflow) == 0)
}
func (cpu *CPU) Bvs(address uint16) {
	cpu.branchIf(address, (cpu.P & Overflow) != 0)
}
func (cpu *CPU) Rts() {
	cpu.PC = cpu.pop16()
}
func (cpu *CPU) Rti() {
	cpu.P = cpu.pop()
	cpu.PC = cpu.pop16()
}
func (cpu *CPU) Brk() {
	cpu.push16(cpu.PC)
	cpu.push(cpu.P)
	cpu.setFlag(BreakCommand, true)
	// Todo: Load IRQ interrupt vector at $FFFE/F
}
// End
