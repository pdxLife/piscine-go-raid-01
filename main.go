package main

import "fmt"
import "os"

func main() {
	v_str := os.Args[1]+os.Args[2]+os.Args[3]+os.Args[4]+os.Args[5]+os.Args[6]+os.Args[7]+os.Args[8]+os.Args[9]
	v_rune := []rune(v_str)
	v_int := make([]int, 81)
	for a := 0; a < 81; a++ {
		if v_rune[a] == '.' {
			v_rune[a] = '0'
		}
		if v_rune[a] >= '0' && v_rune[a] <= '9' {
			v_int[a] = int(v_rune[a]-48)
		}
	}

	var table = [9][9]int{
		{v_int[0], v_int[1], v_int[2], v_int[3], v_int[4], v_int[5], v_int[6], v_int[7], v_int[8]},
		{v_int[9], v_int[10], v_int[11], v_int[12], v_int[13], v_int[14], v_int[15], v_int[16], v_int[17]},
		{v_int[18], v_int[19], v_int[20], v_int[21], v_int[22], v_int[23], v_int[24], v_int[25], v_int[26]},
		{v_int[27], v_int[28], v_int[29], v_int[30], v_int[31], v_int[32], v_int[33], v_int[34], v_int[35]},
		{v_int[36], v_int[37], v_int[38], v_int[39], v_int[40], v_int[41], v_int[42], v_int[43], v_int[44]},
		{v_int[45], v_int[46], v_int[47], v_int[48], v_int[49], v_int[50], v_int[51], v_int[52], v_int[53]},
		{v_int[54], v_int[55], v_int[56], v_int[57], v_int[58], v_int[59], v_int[60], v_int[61], v_int[62]},
		{v_int[63], v_int[64], v_int[65], v_int[66], v_int[67], v_int[68], v_int[69], v_int[70], v_int[71]},
		{v_int[72], v_int[73], v_int[74], v_int[75], v_int[76], v_int[77], v_int[78], v_int[79], v_int[80]}}
	
	IsValid(&table, 0)
	PrintTable(table)
}

func PrintTable(table [9][9]int) {
	var x, y int
	for x = 0; x < 9; x++ {
		for y = 0; y < 9; y++ {
			fmt.Print(table[x][y])
			if y != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func lineCheck(k int, table [9][9]int, x int) bool {
	var y int
	for y = 0; y < 9; y++ {
		if table[x][y] == k {
			return false
		}
	}
	return true
}

func rowCheck(k int, table [9][9]int, y int) bool {
	var x int
	for x = 0; x < 9; x++ {
		if table[x][y] == k {
			return false
		}
	}
	return true
}

func blockCheck(k int, table [9][9]int, x int, y int) bool {
	var firstX, firstY int
	firstX = x - (x % 3)
	firstY = y - (y % 3)
	for x = firstX; x < firstX+3; x++ {
		for y = firstY; y < firstY+3; y++ {
			if table[x][y] == k {
				return false
			}
		}
	}
	return true
}

func IsValid(table *[9][9]int, position int) bool {
	if position == 9*9 {
		return true
	}

	var x, y, k int

	x = position / 9
	y = position % 9

	if table[x][y] != 0 {
		return IsValid(table, position+1)
	}

	for k = 1; k <= 9; k++ {
		if lineCheck(k, *table, x) && rowCheck(k, *table, y) && blockCheck(k, *table, x, y) {
			table[x][y] = k

			if IsValid(table, position+1) {
				return true
			}
		}
	}
	table[x][y] = 0
	return false
}
