package day16

import (
	"fmt"
	"math"
)

type PacketType uint8
const (
	SumPacket PacketType = 0
	ProductPacket  = 1
	MinimumPacket  = 2
	MaximumPacket  = 3
	LiteralPacket = 4
	GreaterThanPacket  = 5
	LessThanPacket  = 6
	EqualToPacket  = 7
)

type Packet struct {
	packetType PacketType
	version uint8
	value int
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

var hexToBinary = map[uint8]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func bitAt(hex string, pos int) uint8 {
	hexDigit := hex[pos / 4]
	bin := hexToBinary[hexDigit]
	bit := bin[pos % 4] - '0'
	return bit
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
		InitState ParserState = 0
		OperatorState ParserState = 1
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
					version: version,
					packetType: packetType,
					subpackets: subpackets,
				}
			} else if length < 0 {
				panic("Too many bits read.");
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
