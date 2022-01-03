package day24

import "strconv"

/**
 * copied an input file and converted it with bunch of regexps
 */

type Block func (s *State)()

func (s *State) runBlock(digitIdx int) {
	switch digitIdx {
	case 1: s.block1()
	case 2: s.block2()
	case 3: s.block3()
	case 4: s.block4()
	case 5: s.block5()
	case 6: s.block6()
	case 7: s.block7()
	case 8: s.block8()
	case 9: s.block9()
	case 10: s.block10()
	case 11: s.block11()
	case 12: s.block12()
	case 13: s.block13()
	case 14: s.block14()
	default: panic("unknown digit idx " + strconv.Itoa(digitIdx))
	}
}

func (s *State) block1() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 14
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 16
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block2() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 11
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 3
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block3() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 12
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 2
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block4() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 11
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 7
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block5() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -10
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 13
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block6() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 15
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 6
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block7() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -14
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 10
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block8() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 10
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 11
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block9() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -4
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 6
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block10() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -3
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 5
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block11() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 1
	s.x = s.x + 13
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 11
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block12() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -3
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 4
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block13() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -9
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 4
	s.y = s.y * s.x
	s.z = s.z + s.y
}

func (s *State) block14() {
	s.x = s.x * 0
	s.x = s.x + s.z
	s.x = s.x % 26
	s.z = s.z / 26
	s.x = s.x + -12
	if s.x == s.w {
		s.x = 1
	} else {
		s.x = 0
	}
	if s.x == 0 {
		s.x = 1
	} else {
		s.x = 0
	}
	s.y = s.y * 0
	s.y = s.y + 25
	s.y = s.y * s.x
	s.y = s.y + 1
	s.z = s.z * s.y
	s.y = s.y * 0
	s.y = s.y + s.w
	s.y = s.y + 6
	s.y = s.y * s.x
	s.z = s.z + s.y
}
