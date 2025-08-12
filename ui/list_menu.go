package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ListMenu() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")
	myWindow.Resize(fyne.NewSize(350, 680))

	leftContent := container.NewMax()

	labels := []string{"Mon", "Tues", "Wed", "Thurs", "Fri", "Sat", "Sun"}

	list := widget.NewList(
		func() int {
			return len(labels)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(labels[i])
		},
	)

	//list.OnSelected = func(id widget.ListItemID) {
	//	fmt.Println("selected")
	//}

	// Put the list where you want it — here I put it on the left
	//content := container.NewBorder(
	//	container.NewMax(widget.NewSeparator()), nil, list, nil,
	//)

	leftWidgets := container.NewHSplit(list, widget.NewEntry())
	leftContent.Add(leftWidgets)

	myWindow.SetContent(leftContent)
	myWindow.ShowAndRun()
}
