package main

import "fmt"

/*
Дана последовательность цифр
9876543210

В ней можно расставить знаки + и - так, чтобы результат арифметического выражения был равен 200
Например,
98+76-5+43-2-10 = 200

Надо написать программу, которая найдёт все варианты расставления
*/

/*
* ToDo
* $ go build test.go
* $ ./test
 */

func main() {
	const result = int(200)

	mass := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	dev1 := result - ((mass[0]+mass[2])*10 + mass[1] + mass[3])
	dev2 := result - ((mass[0]+mass[3])*10 + mass[1] + mass[4])
	dev3 := result - ((mass[0]+mass[4])*10 + mass[1] + mass[5])
	dev4 := result - ((mass[0]+mass[5])*10 + mass[1] + mass[6])
	dev5 := result - ((mass[1]+mass[3])*10 + mass[2] + mass[4])

	zn := [][]int{
		[]int{1, 1, 1, 1, 1},
		[]int{1, -1, 1, 1, 1},
		[]int{-1, 1, 1, 1, 1},
		[]int{-1, -1, 1, 1, 1},
		[]int{1, 1, -1, 1, 1},
		[]int{1, -1, -1, 1, 1},
		[]int{-1, 1, -1, 1, 1},
		[]int{-1, -1, -1, 1, 1},
		[]int{1, 1, 1, -1, 1},
		[]int{1, -1, 1, -1, 1},
		[]int{-1, 1, 1, -1, 1},
		[]int{-1, -1, 1, -1, 1},
		[]int{1, 1, -1, -1, 1},
		[]int{1, -1, -1, -1, 1},
		[]int{-1, 1, -1, -1, 1},
		[]int{-1, -1, -1, -1, 1},
		[]int{1, 1, 1, 1, -1},
		[]int{1, -1, 1, 1, -1},
		[]int{-1, 1, 1, 1, -1},
		[]int{-1, -1, 1, 1, -1},
		[]int{1, 1, 1, -1, -1},
		[]int{1, -1, 1, -1, -1},
		[]int{-1, 1, 1, -1, -1},
		[]int{-1, -1, 1, -1, -1},
		[]int{1, 1, -1, -1, -1},
		[]int{1, -1, -1, -1, -1},
		[]int{-1, 1, -1, -1, -1},
		[]int{-1, -1, -1, -1, -1},
	}

	ten := [][]int{
		[]int{10, 0},
		[]int{1, 0},
	}

	tw := [][]int{
		[]int{21, 0},
		[]int{2, 10},
		[]int{2, 1},
	}

	temp1(dev1, mass, zn, ten, tw, result)
	temp2(dev2, mass, zn, ten, tw, result)
	temp3(dev3, mass, zn, ten, tw, result)
	temp4(dev4, mass, zn, ten, tw, result)
	temp5(dev5, mass, zn, ten, tw, result)
}

func temp1(dev1 int, mass []int, zn [][]int, ten [][]int, tw [][]int, result int) {
	up := [][]int{
		[]int{54, 32, 0},
		[]int{54, 3, 0},
		[]int{5, 43, 0},
		[]int{5, 4, 32},
		[]int{5, 4, 3},
	}
	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(ten); j++ {
			temp := up[0][0]*zn[i][0] + up[0][1]*zn[i][1] + up[0][2]*zn[i][2] + ten[j][0]*zn[i][3] + ten[j][1]*zn[i][4]
			if dev1 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[2], mass[3], "+", up[0][0]*zn[i][0], "+", up[0][1]*zn[i][1], "+", up[0][2]*zn[i][2], "+", ten[j][0]*zn[i][3], "+", ten[j][1]*zn[i][4], "=", result)
			}
		}
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[1][0]*zn[i][0] + up[1][1]*zn[i][1] + up[1][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev1 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[2], mass[3], "+", up[1][0]*zn[i][0], "+", up[1][1]*zn[i][1], "+", up[1][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[2][0]*zn[i][0] + up[2][1]*zn[i][1] + up[2][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev1 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[2], mass[3], "+", up[2][0]*zn[i][0], "+", up[2][1]*zn[i][1], "+", up[2][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}
	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(ten); j++ {
			temp := up[3][0]*zn[i][0] + up[3][1]*zn[i][1] + up[3][2]*zn[i][2] + ten[j][0]*zn[i][3] + ten[j][1]*zn[i][4]
			if dev1 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[2], mass[3], "+", up[3][0]*zn[i][0], "+", up[3][1]*zn[i][1], "+", up[3][2]*zn[i][2], "+", ten[j][0]*zn[i][3], "+", ten[j][1]*zn[i][4], "=", result)
			}
		}
	}
	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[4][0]*zn[i][0] + up[4][1]*zn[i][1] + up[4][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev1 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[2], mass[3], "+", up[4][0]*zn[i][0], "+", up[4][1]*zn[i][1], "+", up[4][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}
}

func temp2(dev2 int, mass []int, zn [][]int, ten [][]int, tw [][]int, result int) {
	up := [][]int{
		[]int{7, 43, 0},
		[]int{7, 4, 32},
		[]int{7, 4, 3},
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[0][0]*zn[i][0] + up[0][1]*zn[i][1] + up[0][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev2 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[3], mass[4], "+", up[0][0]*zn[i][0], "+", up[0][1]*zn[i][1], "+", up[0][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(ten); j++ {
			temp := up[1][0]*zn[i][0] + up[1][1]*zn[i][1] + up[1][2]*zn[i][2] + ten[j][0]*zn[i][3] + ten[j][1]*zn[i][4]
			if dev2 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[3], mass[4], "+", up[1][0]*zn[i][0], "+", up[1][1]*zn[i][1], "+", up[1][2]*zn[i][2], "+", ten[j][0]*zn[i][3], "+", ten[j][1]*zn[i][4], "=", result)
			}
		}
	}
	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[2][0]*zn[i][0] + up[2][1]*zn[i][1] + up[2][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev2 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[3], mass[4], "+", up[2][0]*zn[i][0], "+", up[2][1]*zn[i][1], "+", up[2][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}
}

func temp3(dev3 int, mass []int, zn [][]int, ten [][]int, tw [][]int, result int) {
	up := [][]int{
		[]int{7, 6, 32},
		[]int{7, 6, 3},
	}
	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(ten); j++ {
			temp := up[0][0]*zn[i][0] + up[0][1]*zn[i][1] + up[0][2]*zn[i][2] + ten[j][0]*zn[i][3] + ten[j][1]*zn[i][4]
			if dev3 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[4], mass[5], "+", up[0][0]*zn[i][0], "+", up[0][1]*zn[i][1], "+", up[0][2]*zn[i][2], "+", ten[j][0]*zn[i][3], "+", ten[j][1]*zn[i][4], "=", result)
			}
		}
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[1][0]*zn[i][0] + up[1][1]*zn[i][1] + up[1][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev3 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[4], mass[5], "+", up[1][0]*zn[i][0], "+", up[1][1]*zn[i][1], "+", up[1][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}
}

func temp4(dev4 int, mass []int, zn [][]int, ten [][]int, tw [][]int, result int) {
	up := []int{7, 6, 5}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[0]*zn[i][0] + up[1]*zn[i][1] + up[2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev4 == temp {
				fmt.Println(mass[0], mass[1], "+", mass[5], mass[6], "+", up[0]*zn[i][0], "+", up[1]*zn[i][1], "+", up[2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}
}

func temp5(dev5 int, mass []int, zn [][]int, ten [][]int, tw [][]int, result int) {
	up := [][]int{
		[]int{9, 43, 0},
		[]int{9, 4, 32},
		[]int{9, 4, 3},
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[0][0]*zn[i][0] + up[0][1]*zn[i][1] + up[0][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev5 == temp {
				fmt.Println(mass[1], mass[2], "+", mass[3], mass[4], "+", up[0][0]*zn[i][0], "+", up[0][1]*zn[i][1], "+", up[0][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}

	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(ten); j++ {
			temp := up[1][0]*zn[i][0] + up[1][1]*zn[i][1] + up[1][2]*zn[i][2] + ten[j][0]*zn[i][3] + ten[j][1]*zn[i][4]
			if dev5 == temp {
				fmt.Println(mass[1], mass[2], "+", mass[3], mass[4], "+", up[1][0]*zn[i][0], "+", up[1][1]*zn[i][1], "+", up[1][2]*zn[i][2], "+", ten[j][0]*zn[i][3], "+", ten[j][1]*zn[i][4], "=", result)
			}
		}
	}
	for i := 0; i < len(zn); i++ {
		for j := 0; j < len(tw); j++ {
			temp := up[2][0]*zn[i][0] + up[2][1]*zn[i][1] + up[2][2]*zn[i][2] + tw[j][0]*zn[i][3] + tw[j][1]*zn[i][4]
			if dev5 == temp {
				fmt.Println(mass[1], mass[2], "+", mass[3], mass[4], "+", up[2][0]*zn[i][0], "+", up[2][1]*zn[i][1], "+", up[2][2]*zn[i][2], "+", tw[j][0]*zn[i][3], "+", tw[j][1]*zn[i][4], "=", result)
			}
		}
	}
}
