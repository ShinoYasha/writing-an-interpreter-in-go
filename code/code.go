package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte
type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name         string
	OprandWidths []int
}

var definitions = map[Opcode]*Definition{
	// oprandwidth为参数所需的位数
	OpConstant: {"OpConstant", []int{2}},
}

func LookUp(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

func Make(op Opcode, oprands ...int) []byte {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return []byte{}
	}
	// 操作符长度
	instructionLen := 1
	// 操作符长度+操作数位数为实际所需长度
	for _, w := range def.OprandWidths {
		instructionLen += w
	}
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range oprands {
		width := def.OprandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
