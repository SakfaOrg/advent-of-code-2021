package day25

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCucumbers(t *testing.T) {
	demoLines := strings.Split("...>...\n.......\n......>\nv.....>\n......>\n.......\n..vvv..", "\n")

	demoLinesBig := strings.Split("v...>>.vv>\n.vv>>.vv..\n>>.>v>...v\n>>v>>.>.v.\nv>v.vv.v..\n>.>>..v...\n.vv..>.>v.\nv.v..>>v.v\n....v..v.>", "\n")

	t.Run("Test parse", func(t *testing.T) {
		region := parseRegion(demoLines)
		assert.Equal(t, EAST, region[0][3])
		assert.Equal(t, SOUTH, region[6][2])
		assert.Equal(t, EMPTY, region[6][1])
	})

	t.Run("Test step easts", func(t *testing.T) {
		region := parseRegion([]string{"...>>>>>..."})
		region.step()
		region.step()
		assert.Equal(t, "...>>>.>.>.", region.String())
	})

	t.Run("Test wrap east", func(t *testing.T) {
		region := parseRegion([]string{">..>"})
		region.step()
		assert.Equal(t, ".>.>", region.String())
		region.step()
		assert.Equal(t, ">.>.", region.String())
	})

	t.Run("Test wrap south", func(t *testing.T) {
		region := parseRegion([]string{"v", ".", "v"})
		region.step()
		assert.Equal(t, ".\nv\nv", region.String())
		region.step()
		assert.Equal(t, "v\nv\n.", region.String())
	})

	t.Run("Test step easts then souths", func(t *testing.T) {
		region := parseRegion([]string{"..........", ".>v....v..", ".......>..", ".........."})
		region.step()
		assert.Equal(t, "..........\n.>........\n..v....v>.\n..........", region.String())
	})

	t.Run("Test part 1 demo", func(t *testing.T) {
		assert.Equal(t, "Cucumbers stopped moving after 58 steps", Part1(demoLinesBig))
	})
}