package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type triangle struct {
	side1 uint
	side2 uint
	side3 uint
}

// функция для очистки экрана (работает и на Windows, и на macOS/Linux)
func clear_screen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func type_by_angle(tr triangle) {
	var answer string
	maxi := max(tr.side1, tr.side2, tr.side3)
	mini := min(tr.side1, tr.side2, tr.side3)
	mid := tr.side1 + tr.side2 + tr.side3 - mini - maxi
	if maxi*maxi == mini*mini+mid*mid {
		answer += fmt.Sprintf("Треугольник прямоугольный (прямой угол между катетами %d и %d)", mini, mid)
	} else if maxi*maxi < mini*mini+mid*mid {
		answer += "Треугольник остроугольный"
	} else {
		answer += "Треугольник тупоугольный"
	}
	if type_by_side(tr) == "равносторонний" {
		answer += " со всеми равными углами по 60 градусов"
	} else if type_by_side(tr) == "равнобедренный" {
		answer += " с двумя равными углами"
	} else {
		answer += " со всеми разными углами"
	}
	fmt.Println(answer)
}

func angles(tr triangle) {
	a, b, c := float64(tr.side1), float64(tr.side2), float64(tr.side3)
	A := math.Acos((b*b+c*c-a*a)/(2*b*c)) * 180 / math.Pi
	B := math.Acos((a*a+c*c-b*b)/(2*a*c)) * 180 / math.Pi
	C := 180 - A - B
	fmt.Printf("Угол 1 - %f, угол 2 - %f, угол 3 - %f\n", A, B, C)
}

func type_by_side(tr triangle) string {
	if tr.side1 == tr.side2 && tr.side2 == tr.side3 {
		return "равносторонний"
	} else if tr.side1 == tr.side2 || tr.side1 == tr.side3 || tr.side2 == tr.side3 {
		return "равнобедренный"
	} else {
		return "разносторонний"
	}
}

// основная функция
func main() {
	all_triangles := []triangle{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Для начала работы пожалуйста, нажмите Enter...")
		reader.ReadString('\n')
		clear_screen()

		fmt.Println("Укажите 3 стороны треугольника через пробел: ")

		inputSides, _ := reader.ReadString('\n')
		inputSides = strings.TrimSpace(inputSides)

		parts := strings.Fields(inputSides)
		if len(parts) != 3 {
			fmt.Println("Ошибка: нужно ввести ровно 3 значения через пробел.")
			continue
		}

		s1, err1 := strconv.ParseUint(parts[0], 10, 64)
		s2, err2 := strconv.ParseUint(parts[1], 10, 64)
		s3, err3 := strconv.ParseUint(parts[2], 10, 64)

		all_triangles = append(all_triangles, triangle{uint(s1), uint(s2), uint(s3)})
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Некорректный ввод! Пожалуйста, введите числа.")
			continue
		}

		fmt.Println("Меню команд:")
		fmt.Println("1 - Проверка существования треугольника")
		fmt.Println("2 - Расчёт периметра и площади треугольника")
		fmt.Println("3 - Определение типа треугольника по сторонам")
		fmt.Println("4 - Определение типа треугольника по углам")
		fmt.Println("5 - Определение углов треугольника")
		fmt.Println("0 - Завершить работу")
		fmt.Println("Выберите пункт из меню управления: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		a, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Некорректный ввод! Пожалуйста, введите число.")
			continue
		}
		switch a {
		case 0:
			return
		case 1:
			//функция Проверки существования треугольника
		case 2:
			//функция Расчёта периметра и площади треугольника
		case 3:
			fmt.Println("Треугольник", type_by_side(all_triangles[0]))
		case 4:
			type_by_angle(all_triangles[0])
		case 5:
			angles(all_triangles[0])
		default:
			fmt.Println("Нет такого пункта меню!")
		}
	}
}
