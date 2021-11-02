package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var myapp fyne.App = app.New()
var mywindow fyne.Window = myapp.NewWindow("SsjOS")

func main() {
	// app := app.New()

	// w := app.NewWindow("Hello")

	img := canvas.NewImageFromFile("static\\wallpaper5.jpg")
	img.FillMode = canvas.ImageFillStretch

	icon1 :=canvas.NewImageFromFile("static\\calclogo.png")
	icon1.FillMode = canvas.ImageFillContain
	calc := widget.NewButton("",func ()  {
		showCalc()
	})
	icon2 :=canvas.NewImageFromFile("static\\weathericon.png")
	icon2.FillMode = canvas.ImageFillContain
	weather := widget.NewButton("",func ()  {
		showWeatherApp()
	})
	icon3 :=canvas.NewImageFromFile("static\\gallery1.png")
	icon3.FillMode = canvas.ImageFillContain
	gallery := widget.NewButton("",func ()  {
		showGalleryApp()
	})
	icon4 :=canvas.NewImageFromFile("static\\notepad.png")
	icon4.FillMode = canvas.ImageFillContain
	notepad := widget.NewButton("",func ()  {
		showTextEditor()
	})
	icon5 :=canvas.NewImageFromFile("static\\music2.png")
	icon5.FillMode = canvas.ImageFillContain
	music := widget.NewButton("",func ()  {
		showAudioPlayer()
	})

	icon6 :=canvas.NewImageFromFile("static\\newslogo.png")
	icon6.FillMode = canvas.ImageFillContain
	news := widget.NewButton("",func ()  {
		showNews()
	})

	icon7 :=canvas.NewImageFromFile("static\\shutdown.png")
	icon7.FillMode = canvas.ImageFillContain
	Shutdown := widget.NewButton("", func() {
		mywindow.Close()
	})

	bg_color := canvas.NewRectangle(color.RGBA{R:0,G:0,B:0,A:100})

	apps := container.New(layout.NewVBoxLayout(),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(calc, icon1)),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(weather, icon2)),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(gallery, icon3)),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(notepad, icon4)),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(music, icon5)),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(news, icon6)),
	layout.NewSpacer(),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(Shutdown, icon7)),
	)

	pane:= container.New(
		layout.NewMaxLayout(),
		bg_color,
		apps,
	)
	mywindow.Resize(fyne.NewSize(1280,720))
	mywindow.CenterOnScreen()
	r, _ := fyne.LoadResourceFromPath("static\\oslogo.png")
	mywindow.SetIcon(r)
	mywindow.SetContent(
		container.New(layout.NewMaxLayout(),img,
		container.New(layout.NewBorderLayout(nil, nil,pane, nil),pane)))
	mywindow.ShowAndRun()
}