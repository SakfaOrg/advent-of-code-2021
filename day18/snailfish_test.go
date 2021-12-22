package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnailfishNumbers(t *testing.T) {
	t.Run("Test parse pair", func(t *testing.T) {
		parsed := parse("[12,23]")
		assertLinks(parsed)

		assert.Equal(t, 12, parsed.leftChild.value)
		assert.Nil(t, parsed.leftChild.leftSibling)
		assert.Equal(t, 23, parsed.leftChild.rightSibling.value)

		assert.Equal(t, 23, parsed.rightChild.value)
		assert.Equal(t, 12, parsed.rightChild.leftSibling.value)
		assert.Nil(t, parsed.rightChild.rightSibling)
	})

	t.Run("Test parse nested", func(t *testing.T) {
		parsed := parse("[12,[23,34]]")

		assert.Equal(t, 12, parsed.leftChild.value)
		assert.Equal(t, 23, parsed.rightChild.leftChild.value)
		assert.Equal(t, 34, parsed.rightChild.rightChild.value)
	})

	t.Run("Test track siblings", func(t *testing.T) {
		parsed := parse("[12,[23,34]]")

		assert.Equal(t, 23, parsed.leftChild.rightSibling.value)
		assert.Nil(t, parsed.leftChild.leftSibling)

		assert.Equal(t, 12, parsed.rightChild.leftChild.leftSibling.value)
		assert.Equal(t, 34, parsed.rightChild.leftChild.rightSibling.value)

		assert.Equal(t, 23, parsed.rightChild.rightChild.leftSibling.value)
		assert.Nil(t, parsed.rightChild.rightChild.rightSibling)
	})

	t.Run("Test track parents", func(t *testing.T) {
		parsed := parse("[12,[23,34]]")

		assert.Equal(t, 12, parsed.rightChild.rightChild.parent.parent.leftChild.value)
	})

	t.Run("Test magnitude", func(t *testing.T) {
		assert.Equal(t, 143, parse("[[1,2],[[3,4],5]]").magnitude())
	})

	t.Run("Test add", func(t *testing.T) {
		a := parse("[1,2]")
		b := parse("[[3,4],5]")
		sum := add(a, b)
		assertLinks(sum)
		assert.Equal(t, "[[1,2],[[3,4],5]]", sum.String())
	})

	t.Run("Test add needing reduce", func(t *testing.T) {
		a := parse("[[[[4,3],4],4],[7,[[8,4],9]]]")
		b := parse("[1,1]")
		sum := add(a, b)
		assertLinks(sum)
		assert.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", sum.String())
	})

	t.Run("Test explode once example 1", func(t *testing.T) {
		number := parse("[[[[[9,8],1],2],3],4]")
		reduce(number)
		assertLinks(number)
		assert.Equal(t, "[[[[0,9],2],3],4]", number.String())
	})

	t.Run("Test explode once example 2", func(t *testing.T) {
		number := parse("[7,[6,[5,[4,[3,2]]]]]")
		reduce(number)
		assertLinks(number)
		assert.Equal(t, "[7,[6,[5,[7,0]]]]", number.String())
	})

	t.Run("Test explode once example 3", func(t *testing.T) {
		number := parse("[[6,[5,[4,[3,2]]]],1]")
		reduce(number)
		assertLinks(number)
		assert.Equal(t, "[[6,[5,[7,0]]],3]", number.String())
	})

	t.Run("Test explode once example 4", func(t *testing.T) {
		number := parse("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
		explodeWalk(0, number)
		assertLinks(number)
		assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", number.String())
	})

	t.Run("Test reduce example 4 loop", func(t *testing.T) {
		number := parse("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
		reduce(number)
		assertLinks(number)
		assert.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", number.String())
	})

	t.Run("Test split example 1", func(t *testing.T) {
		number := parse("[[[[0,7],4],[15,[0,13]]],[1,1]]")
		splitWalk(number)
		assert.Equal(t, "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", number.String())
		splitWalk(number)
		assertLinks(number)
		assert.Equal(t, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", number.String())
	})

	t.Run("Test reduce sample input", func(t *testing.T) {
		number := parse("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
		reduce(number)
		assertLinks(number)
		assert.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", number.String())
	})

	t.Run("Test add many", func(t *testing.T) {
		lines := []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}
		result := addLines(lines)
		assertLinks(result)
		assert.Equal(t, "[[[[5,0],[7,4]],[5,5]],[6,6]]", result.String())
	})

	t.Run("Test weird", func(t *testing.T) {
		a := parse("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
		b := parse("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]")
		result := add(a, b)
		assertLinks(result)
		assert.Equal(t, "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", result.String())
	})

	t.Run("Test demo input", func(t *testing.T) {
		lines := []string{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			"[7,[5,[[3,8],[1,4]]]]",
			"[[2,[2,2]],[8,[8,1]]]",
			"[2,9]",
			"[1,[[[9,3],9],[[9,0],[0,7]]]]",
			"[[[5,[7,4]],7],1]",
			"[[[[4,2],2],6],[8,7]]",
		}
		result := addLines(lines)
		assertLinks(result)
		assert.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", result.String())
		assert.Equal(t, 3488, result.magnitude())
	})

	demoLines := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}

	t.Run("Test demo input 2", func(t *testing.T) {
		result := addLines(demoLines)
		assertLinks(result)
		assert.Equal(t, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", result.String())
		assert.Equal(t, 4140, result.magnitude())
	})

	t.Run("Test part2", func(t *testing.T) {
		result := Part2(demoLines)
		assert.Equal(t,
			"Maximum magnitude 3993 when adding [[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]] + [[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			result)
	})
}

/**
 * asserts all child<->parent and left sibling <-> right sibling relations match
 */
func assertLinks(a *SnailfishNumber) {
	if a.isPair() {
		if a.leftChild.parent != a {
			panic("a.leftChild.parent is not a!")
		}
		if a.rightChild.parent != a {
			panic("a.rightChild.parent is not a!")
		}
		assertLinks(a.leftChild)
		assertLinks(a.rightChild)
	} else {
		if a.leftSibling != nil {
			if a != a.leftSibling.rightSibling {
				panic("a is not right sibling of its left sibling!")
			}
		}
		if a.rightSibling != nil {
			if a != a.rightSibling.leftSibling {
				panic("a is not left sibling of its right sibling!")
			}
		}
	}
}
