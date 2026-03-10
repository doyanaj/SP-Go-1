package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var start, name, surname, patronymic string
	var sum1, sum2, sum3 float64

	fmt.Scan(&start, &name, &surname, &patronymic,
		&sum1, &sum2, &sum3)

	tstart, _ := time.Parse("02.01.2006", start)
	dsign := tstart.AddDate(0, 0, 15)
	sum := sum1 + sum2 + sum3
	rub, kop := math.Modf(sum)
	kop *= 100
	kop = math.Round(kop)

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n"+
		"Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n"+
		"Общая сумма выплат составит %.0f руб. %.0f коп.\n\n"+
		"С уважением,\nГл. бух. Иванов А.Е.\n",
		surname, name, patronymic, dsign.Format("02.01.2006"), rub, kop)
}
