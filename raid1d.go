package student

import "github.com/01-edu/z01"

func Raid1d(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 {
			z01.PrintRune(65)
			z01.PrintRune(10)
		}
		if x > 1 {
			z01.PrintRune(65)
			for xl := 0; xl < x-2; xl++ {
				z01.PrintRune(66)
			}
			z01.PrintRune(67)
			z01.PrintRune(10)
		}
		if x == 1 {
			for yl := 0; yl < y-2; yl++ {
				z01.PrintRune(66)
				z01.PrintRune(10)
			}
		}
		if y >= 1 && x != 1 {
			for yl := 0; yl < y-2; yl++ {
				z01.PrintRune(66)
				for xl2 := 0; xl2 < x-2; xl2++ {
					z01.PrintRune(32)
				}
				z01.PrintRune(66)
				z01.PrintRune(10)
			}
		}
		if x == 1 && y != 1 {
			z01.PrintRune(65)
			z01.PrintRune(10)
		}
		if x > 1 && y != 1 {
			z01.PrintRune(65)
			for xl := 0; xl < x-2; xl++ {
				z01.PrintRune(66)
			}
			z01.PrintRune(67)
			z01.PrintRune(10)
		}
	}
}
