// All rights reserved by someone from the Legion. ST.
package main

import (
	"fmt"
	"os"
)

// Клавиши для движения ракеток
const (
	// Движение ракетки 1 вверх
	a    = byte('a')
	astr = string(a)
	A    = byte('A')
	Astr = string(A)

	// Движение ракетки 1 вниз
	z    = byte('z')
	zstr = string(z)
	Z    = byte('Z')
	Zstr = string(Z)

	// Движение ракетки 2 вверх
	k    = byte('k')
	kstr = string(k)
	K    = byte('K')
	Kstr = string(K)

	// Движение ракетки 2 вниз
	m    = byte('m')
	mstr = string(m)
	M    = byte('M')
	Mstr = string(M)

	// Клавиша выхода
	q    = byte('q')
	qstr = string(q)
	Q    = byte('Q')
	Qstr = string(Q)

	// Мячик
	p    = byte('@')
	pstr = string(p)
)

var (
	// Координаты поля
	X = 60
	Y = 20

	// Координаты ракеток
	Pad1x = X / 20
	Pad1y = Y / 2
	Pad2x = X - (X / 20)
	Pad2y = Y / 2

	// Координаты мячика
	BallX = X / 2
	BallY = Y / 2

	// Вектор движения мячика\скорость
	dx = 2
	dy = 1

	// Счетчики
	P1count = 0
	P2count = 0

	// Переменные :)
	c          string
	x, y, flag int
)

func main() {
	Instruction_header()
	for flag != 3 {
		Score_counting()
		Field_rendering()
		Winner_message()
		Racket_movement()
		Ball_movement()
		if c == qstr || c == Qstr { // Exit button
			break
		}
	}
}

// Инструкция/Шапка
func Instruction_header() {
	fmt.Println(` 
	Welcome to step-by-step ping Pong!
	Functional buttons
	a/z - racket 1 - up/down
	k/m - racket 2 - up/down
	q - quit
	any other button - space move
	3 Score to Win!
	Enjoy!
	Please click enter to start`)
	fmt.Scanln()
}

// Отрисовка поля
func Field_rendering() {
	fmt.Println("a-z - racket 1 \t\t q - quit \t k/m - racket 2")
	for y = 0; y < Y; y++ {
		for x = 0; x < X; x++ {
			if x == 0 || x == X-1 {
				fmt.Print("|") // боковые границы
			} else if y == 0 || y == Y-1 {
				fmt.Print("-") // верхния и нижние границы
			} else if x == X/2+5 && y == 3 {
				fmt.Printf("%d", P1count) // Счет Игрок 1
			} else if x == X/2-5 && y == 3 {
				fmt.Printf("%d", P2count) // Счет Игрок 2
			} else if (x == BallX) && (y == BallY) {
				fmt.Print(pstr) // Мячик
			} else if x == X/2 {
				fmt.Print("|") // Сетка по центру
			} else if x == Pad1x && (y == Pad1y || y == Pad1y+1 || y == Pad1y-1) {
				fmt.Print("!") // Ракетка 1
			} else if x == Pad2x && (y == Pad2y || y == Pad2y+1 || y == Pad2y-1) {
				fmt.Print("!") // Ракетка 2
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

// Движение мячика
func Ball_movement() {
	BallX += dx
	BallY += dy
	if BallY == Y-2 || BallY == 1 {
		dy = -(dy)
	}
	if BallX == Pad1x+1 && (BallY == Pad1y || BallY == Pad1y-1) {
		dx = -(dx)
	} else if BallX == Pad1x+1 && (BallY == Pad1y || BallY == Pad1y+1) {
		dx = -(dx)
		dy = -(dy)
	}
	if BallX == Pad2x-1 && (BallY == Pad2y || BallY == Pad2y-1) {
		dx = -(dx)
	} else if BallX == Pad2x-1 && (BallY == Pad2y || BallY == Pad2y+1) {
		dx = -(dx)
		dy = -(dy)
	}
}

// Движение ракеток
func Racket_movement() {
	fmt.Fscan(os.Stdin, &c)
	if (c == astr || c == Astr) && Pad1y != 2 {
		Pad1y = Pad1y - 1
	} else if (c == zstr || c == Zstr) && Pad1y != Y-3 {
		Pad1y = Pad1y + 1
	} else if (c == kstr || c == Kstr) && Pad2y != 2 {
		Pad2y = Pad2y - 1
	} else if (c == mstr || c == Mstr) && Pad2y != Y-3 {
		Pad2y = Pad2y + 1
	}
}

// Подсчет очков
func Score_counting() {
	if BallX <= 0 {
		P1count++
		flag = P1count
		BallX = X / 2
		BallY = Y / 2
		Pad1y = Y / 2
		Pad2y = Y / 2
		dx = -(dx)
	} else if BallX >= X-1 {
		P2count++
		flag = P2count
		BallX = X / 2
		BallY = Y / 2
		Pad1y = Y / 2
		Pad2y = Y / 2
		dx = -(dx)
	}
}

// Сообщение для победителя
func Winner_message() {
	if P2count == 3 {
		fmt.Print("Congratulations! Player 1 Win!\nPlease, press any key for exit ^_^\n")
	} else if P1count == 3 {
		fmt.Print("Congratulations! Player 2 Win!\nPlease, press any key for exit ^_^\n")
	}
}
