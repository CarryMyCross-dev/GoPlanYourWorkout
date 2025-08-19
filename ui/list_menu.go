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

	labels := [7]string{"Mon", "Tues", "Wed", "Thurs", "Fri", "Sat", "Sun"}

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
	
	rightContent := widget.NewLabel("")

	list.OnSelected = func(id widget.ListItemID) {
		rightContent.SetText(labels[id])
		rightContent.Refresh()
	}

	leftWidgets := container.NewHSplit(list, rightContent)
	leftContent.Add(leftWidgets)

	myWindow.SetContent(leftContent)
	myWindow.ShowAndRun()
}
