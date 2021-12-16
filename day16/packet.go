package day16

import (
	"fmt"
	"math"
)

type PacketType uint8

const (
	SumPacket PacketType = iota
	ProductPacket
	MinimumPacket
	MaximumPacket
	LiteralPacket
	GreaterThanPacket
	LessThanPacket
	EqualToPacket
)

type Packet struct {
	packetType PacketType
	version    uint8
	value      int
	subpackets []Packet
}

func (p Packet) evaluate() int {
	switch p.packetType {
	case LiteralPacket:
		return p.value
	case SumPacket:
		sum := 0
		for _, subpacket := range p.subpackets {
			sum += subpacket.evaluate()
		}
		return sum
	case ProductPacket:
		product := 1
		for _, subpacket := range p.subpackets {
			product *= subpacket.evaluate()
		}
		return product
	case MinimumPacket:
		minimum := math.MaxInt
		for _, subpacket := range p.subpackets {
			val := subpacket.evaluate()
			if val < minimum {
				minimum = val
			}
		}
		return minimum
	case MaximumPacket:
		maximum := -1
		for _, subpacket := range p.subpackets {
			val := subpacket.evaluate()
			if val > maximum {
				maximum = val
			}
		}
		return maximum
	case GreaterThanPacket:
		if p.subpackets[0].evaluate() > p.subpackets[1].evaluate() {
			return 1
		} else {
			return 0
		}
	case LessThanPacket:
		if p.subpackets[0].evaluate() < p.subpackets[1].evaluate() {
			return 1
		} else {
			return 0
		}
	case EqualToPacket:
		if p.subpackets[0].evaluate() == p.subpackets[1].evaluate() {
			return 1
		} else {
			return 0
		}
	default:
		panic("Not implemented")
	}
}

func (p Packet) versionSum() int {
	result := int(p.version)
	for _, subpacket := range p.subpackets {
		result += subpacket.versionSum()
	}
	return result
}

// yep, this is fast
func bitAt(hex string, pos int) uint8 {
	hexDigit := hex[pos/4]
	bitPos := pos % 4
	switch true {
	case hexDigit == '1' && bitPos == 3:
		return 1
	case hexDigit == '2' && bitPos == 2:
		return 1
	case hexDigit == '3' && (bitPos == 2 || bitPos == 3):
		return 1
	case hexDigit == '4' && bitPos == 1:
		return 1
	case hexDigit == '5' && (bitPos == 1 || bitPos == 3):
		return 1
	case hexDigit == '6' && (bitPos == 1 || bitPos == 2):
		return 1
	case hexDigit == '7' && bitPos != 0:
		return 1
	case hexDigit == '8' && bitPos == 0:
		return 1
	case hexDigit == '9' && (bitPos == 0 || bitPos == 3):
		return 1
	case hexDigit == 'A' && (bitPos == 0 || bitPos == 2):
		return 1
	case hexDigit == 'B' && bitPos != 1:
		return 1
	case hexDigit == 'C' && (bitPos == 0 || bitPos == 1):
		return 1
	case hexDigit == 'D' && bitPos != 2:
		return 1
	case hexDigit == 'E' && bitPos != 3:
		return 1
	case hexDigit == 'F':
		return 1
	default:
		return 0
	}
}

func decodeInt(hex string, pos *int, bits int) int {
	power := 1
	result := 0
	for p := *pos + bits - 1; p >= *pos; p-- {
		result += int(bitAt(hex, p)) * power
		power *= 2
	}
	*pos += bits
	return result
}

func combineChunks(chunks []int) int {
	power := 1
	result := 0
	for i := len(chunks) - 1; i >= 0; i-- {
		result += chunks[i] * power
		power *= 16
	}
	return result
}

func decodeTopLevel(hex string) Packet {
	pos := 0
	return decodePacket(hex, &pos)
}

func decodePacket(hex string, pos *int) Packet {
	type ParserState uint8
	const (
		InitState ParserState = iota
		OperatorState
	)

	var version uint8
	var packetType PacketType
	var value int
	var lengthType int
	var length int
	var subpackets []Packet

	state := InitState
	for {
		switch state {
		case InitState:
			version = uint8(decodeInt(hex, pos, 3))
			packetType = PacketType(decodeInt(hex, pos, 3))
			if packetType == 4 {
				var chunks []int
				hasNext := true
				for hasNext {
					hasNext = decodeInt(hex, pos, 1) == 1
					chunks = append(chunks, decodeInt(hex, pos, 4))
				}
				value = combineChunks(chunks)
				return Packet{
					version:    version,
					packetType: packetType,
					value:      value,
				}
			} else {
				lengthType = decodeInt(hex, pos, 1)
				if lengthType == 0 {
					length = decodeInt(hex, pos, 15)
				} else {
					length = decodeInt(hex, pos, 11)
				}
				state = OperatorState
			}
		case OperatorState:
			posBefore := *pos
			subpacket := decodePacket(hex, pos)
			subpackets = append(subpackets, subpacket)

			if lengthType == 0 {
				consumedBits := *pos - posBefore
				length -= consumedBits
			} else {
				length -= 1
			}

			if length == 0 {
				return Packet{
					version:    version,
					packetType: packetType,
					subpackets: subpackets,
				}
			} else if length < 0 {
				panic("Too many bits read.")
			}
		}
	}
}

func Part1(lines []string) string {
	packet := decodeTopLevel(lines[0])
	return fmt.Sprintf("Version sum: %d", packet.versionSum())
}

func Part2(lines []string) string {
	packet := decodeTopLevel(lines[0])
	return fmt.Sprintf("Packet value: %d", packet.evaluate())
}
