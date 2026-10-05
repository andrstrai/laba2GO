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
	side1 float64
	side2 float64
	side3 float64
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

func float_equal(a, b float64) bool {
	const eps = 1e-9
	diff := math.Abs(a - b)
	scale := math.Max(math.Abs(a), math.Abs(b))
	return diff <= eps*math.Max(1, scale)
}

func type_by_angle(tr triangle) {
	var answer string
	maxi := max(tr.side1, tr.side2, tr.side3)
	mini := min(tr.side1, tr.side2, tr.side3)
	mid := tr.side1 + tr.side2 + tr.side3 - mini - maxi

	maxi_sq := maxi * maxi
	summ_sq := mini*mini + mid*mid
	if float_equal(maxi_sq, summ_sq) {
		answer += fmt.Sprintf("Треугольник прямоугольный (прямой угол между катетами %.15f и %.15f)", mini, mid)
	} else if maxi_sq < summ_sq {
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
	if type_by_side(tr) == "равносторонний" {
		fmt.Println("Угол 1 - 60, угол 2 - 60, угол 3 - 60")
		return
	}
	maxi := max(tr.side1, tr.side2, tr.side3)
	mini := min(tr.side1, tr.side2, tr.side3)
	mid := tr.side1 + tr.side2 + tr.side3 - mini - maxi
	if float_equal(maxi*maxi, mini*mini+mid*mid) {
		alpha := math.Atan2(mini, mid) * 180 / math.Pi
		beta := 90 - alpha
		fmt.Printf("Угол 1 - 90, угол 2 - %.15f, угол 3 - %.15f\n", alpha, beta)
		return
	}
	a, b, c := tr.side1, tr.side2, tr.side3
	A := math.Acos((b*b+c*c-a*a)/(2*b*c)) * 180 / math.Pi
	B := math.Acos((a*a+c*c-b*b)/(2*a*c)) * 180 / math.Pi
	C := 180 - A - B
	fmt.Printf("Угол 1 - %.15f, угол 2 - %.15f, угол 3 - %.15f\n", A, B, C)
}

func type_by_side(tr triangle) string {
	if float_equal(tr.side1, tr.side2) && float_equal(tr.side2, tr.side3) {
		return "равносторонний"
	} else if float_equal(tr.side1, tr.side2) || float_equal(tr.side1, tr.side3) || float_equal(tr.side2, tr.side3) {
		return "равнобедренный"
	} else {
		return "разносторонний"
	}
}

// основная функция
func main() {
	all_triangles := []triangle{}
	reader := bufio.NewReader(os.Stdin)
A:
	for {
		fmt.Println("Для начала работы пожалуйста, нажмите Enter, для завершения введите 0.....")
		inputStart, _ := reader.ReadString('\n')
		inputStart = strings.TrimSpace(inputStart)
		if inputStart == "0" {
			return
		} else {
			clear_screen()

			fmt.Println("Укажите 3 стороны треугольника через пробел: ")

			inputSides, _ := reader.ReadString('\n')
			inputSides = strings.TrimSpace(inputSides)

			parts := strings.Fields(inputSides)
			if len(parts) != 3 {
				fmt.Println("Ошибка: нужно ввести ровно 3 значения через пробел.")
				continue
			}

			s1, err1 := strconv.ParseFloat(parts[0], 64)
			s2, err2 := strconv.ParseFloat(parts[1], 64)
			s3, err3 := strconv.ParseFloat(parts[2], 64)

			if err1 != nil || err2 != nil || err3 != nil || s1 <= 0 || s2 <= 0 || s3 <= 0 {
				fmt.Println("Некорректный ввод! Пожалуйста, введите положительные числа.")
				continue
			}
			all_triangles = append(all_triangles, triangle{s1, s2, s3})

			for {
				fmt.Println("Меню команд:")
				fmt.Println("1 - Проверка существования треугольника")
				fmt.Println("2 - Расчёт периметра и площади треугольника")
				fmt.Println("3 - Определение типа треугольника по сторонам")
				fmt.Println("4 - Определение типа треугольника по углам")
				fmt.Println("5 - Определение углов треугольника")
				fmt.Println("0 - Выйти из меню команд")
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
					continue A
				case 1:
					//функция Проверки существования треугольника
				case 2:
					//функция Расчёта периметра и площади треугольника
				case 3:
					fmt.Println("Треугольник", type_by_side(all_triangles[len(all_triangles)-1]))
				case 4:
					type_by_angle(all_triangles[len(all_triangles)-1])
				case 5:
					angles(all_triangles[len(all_triangles)-1])
				default:
					fmt.Println("Нет такого пункта меню!")
				}
			}
		}

	}
}
