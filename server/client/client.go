package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"grpc_pr3.1/pkg/proto/proto"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"google.golang.org/grpc"
)

func connect_clietn(x float64, a float64, b float64, c float64) string {

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())

	if err != nil {
		log.Fatal("11")
	}

	client := proto.NewAdderClient(conn)

	responce, err := client.Add(context.Background(), &proto.AddRequest{X: x, A: a, B: b, C: c})

	if err != nil {
		log.Fatal("Помилка відповіді!")
	}

	fmt.Println(responce.GetResult())

	result := responce.GetResult()
	result_str := strconv.FormatFloat(result, 'f', -1, 64)

	return result_str
}

func main() {

	a := app.New()

	w := a.NewWindow("Maths")
	w.Resize(fyne.NewSize(600, 500))

	label_info := widget.NewLabel("Введіть значення x,a,b,c")

	label_x := widget.NewLabel("Введіть значення x:")
	entry_x := widget.NewEntry()

	label_a := widget.NewLabel("Введіть значення a:")
	entry_a := widget.NewEntry()

	label_b := widget.NewLabel("Введіть значення b:")
	entry_b := widget.NewEntry()

	label_c := widget.NewLabel("Введіть значення c:")
	entry_c := widget.NewEntry()

	label_footer := widget.NewLabel("Відповідь:")
	label_result := widget.NewLabel("")

	button := widget.NewButton("Обрахувати", func() {

		x, errx := strconv.ParseFloat(entry_x.Text, 32)
		a, erra := strconv.ParseFloat(entry_a.Text, 32)
		b, errb := strconv.ParseFloat(entry_b.Text, 32)
		c, errc := strconv.ParseFloat(entry_c.Text, 32)

		if errx != nil || erra != nil || errb != nil || errc != nil {
			log.Fatal("Помилка введених даних!")
		}

		answer_serv := connect_clietn(x, a, b, c)

		label_result.SetText(answer_serv)
	})

	w.SetContent(container.NewVBox(label_info, label_x, entry_x, label_a, entry_a, label_b, entry_b, label_c, entry_c, button, label_footer, label_result))

	w.ShowAndRun()
}
