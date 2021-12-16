package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecoder(t *testing.T) {
	t.Run("Test bit at", func(t *testing.T) {
		expected := "110100101111111000101000"
		hexString := "D2FE28"
		for pos := 0; pos < len(hexString); pos++ {
			given := bitAt(hexString, pos)
			assert.Equal(t, expected[pos]-'0', given)
		}
	})

	t.Run("Test decode int", func(t *testing.T) {
		hexString := "D2FE28"
		pos := 0
		assert.Equal(t, 6, decodeInt(hexString, &pos, 3))
		assert.Equal(t, 4, decodeInt(hexString, &pos, 3))
		assert.Equal(t, 6, pos)
	})

	t.Run("Test combine chunks", func(t *testing.T) {
		assert.Equal(t, 2021, combineChunks([]int{7, 14, 5}))
	})

	t.Run("Test literal packet", func(t *testing.T) {
		pos := 0
		packet := decodePacket("D2FE28", &pos)
		assert.Equal(t, uint8(6), packet.version)
		assert.Equal(t, PacketType(LiteralPacket), packet.packetType)
		assert.Equal(t, 2021, packet.value)
		assert.Equal(t, 21, pos)
	})

	t.Run("Test packet operator len type 0", func(t *testing.T) {
		hexString := "38006F45291200"
		packet := decodeTopLevel(hexString)
		assert.Equal(t, uint8(1), packet.version)
		assert.Equal(t, PacketType(LessThanPacket), packet.packetType)
		assert.Equal(t, 2, len(packet.subpackets))
		assert.Equal(t, 10, packet.subpackets[0].value)
		assert.Equal(t, 20, packet.subpackets[1].value)
	})

	t.Run("Test packet operator len type 1", func(t *testing.T) {
		hexString := "EE00D40C823060"
		packet := decodeTopLevel(hexString)
		assert.Equal(t, uint8(7), packet.version)
		assert.Equal(t, PacketType(3), packet.packetType)
		assert.Equal(t, 3, len(packet.subpackets))
		assert.Equal(t, 1, packet.subpackets[0].value)
		assert.Equal(t, 2, packet.subpackets[1].value)
		assert.Equal(t, 3, packet.subpackets[2].value)
	})

	t.Run("Test nested operator packets", func(t *testing.T) {
		hexString := "A0016C880162017C3686B18A3D4780"
		packet := decodeTopLevel(hexString)

		assert.Equal(t, 5, len(packet.subpackets[0].subpackets[0].subpackets))
		assert.Equal(t, 31, packet.versionSum())
	})

	t.Run("Test sum packet", func(t *testing.T) {
		assert.Equal(t, 3, decodeTopLevel("C200B40A82").evaluate())
		assert.Equal(t, 54, decodeTopLevel("04005AC33890").evaluate())
		assert.Equal(t, 7, decodeTopLevel("880086C3E88112").evaluate())
		assert.Equal(t, 9, decodeTopLevel("CE00C43D881120").evaluate())
		assert.Equal(t, 1, decodeTopLevel("D8005AC2A8F0").evaluate())
		assert.Equal(t, 0, decodeTopLevel("F600BC2D8F").evaluate())
		assert.Equal(t, 0, decodeTopLevel("9C005AC2F8F0").evaluate())
		assert.Equal(t, 1, decodeTopLevel("9C0141080250320F1802104A08").evaluate())
	})
}
