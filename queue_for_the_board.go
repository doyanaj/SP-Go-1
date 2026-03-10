package main

import "fmt"

func main() {
	var name string
	var num int

	queue := [5]string{"-", "-", "-", "-", "-"}

	for {
		fmt.Scan(&name)
		if name == "очередь" || name == "конец" {
			for i := 0; i < 5; i++ {
				fmt.Printf("%d. %s\n", i+1, queue[i])
			}
			if name == "конец" {
				break
			}
		} else if name == "количество" {
			free := 0
			occ := 0

			for i := 0; i < 5; i++ {
				if queue[i] == "-" {
					free++
				} else {
					occ++
				}
			}

			fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", free, occ)
		} else {
			fmt.Scan(&num)

			if num < 1 || num > 5 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
				continue
			}

			full := true
			for i := 0; i < 5; i++ {
				if queue[i] == "-" {
					full = false
					break
				}
			}

			if full {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
				continue
			}

			if queue[num-1] != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
				continue
			}

			queue[num-1] = name
		}
	}
}
