package main

import (
	"bufio"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myapp fyne.App = app.New();
var myWindow fyne.Window = myapp.NewWindow("Ssj Os")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget

var img fyne.CanvasObject;
var DeskBtn fyne.Widget
var panelContent *fyne.Container


// btn := widget.NewButtonWithIcon("Browse", fyne.NewStaticResource("icon", b), func() {
	// do something
// })

func main(){
	// myapp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("static\\wallpaper5.jpg")
	// iconFile, _ := fyne.LoadResourceFromPath("static\\oslogo.jpg")
	i1, _ := os.Open("static\\weather3.png")
	r1 := bufio.NewReader(i1)
	b1, _ := ioutil.ReadAll(r1)
	btn1 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b1), func() {
		showWeatherApp()
	})
	i2, _ := os.Open("static\\calclogo.png")
	r2 := bufio.NewReader(i2)
	b2, _ := ioutil.ReadAll(r2)
	btn2 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b2), func() {
		showCalc()
	})
	i3, _ := os.Open("static\\gallery.png")
	r3 := bufio.NewReader(i3)
	b3, _ := ioutil.ReadAll(r3)
	btn3 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b3), func() {
		showGalleryApp()
	})
	i4, _ := os.Open("static\\notepadlogo.png")
	r4 := bufio.NewReader(i4)
	b4, _ := ioutil.ReadAll(r4)
	btn4 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b4), func() {
		showTextEditor()
	})
	i5, _ := os.Open("static\\music1.png")
	r5 := bufio.NewReader(i5)
	b5, _ := ioutil.ReadAll(r5)
	btn5 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b5), func() {
		showAudioPlayer()
	})
	i6, _ := os.Open("static\\newslogo.png")
	r6 := bufio.NewReader(i6)
	b6, _ := ioutil.ReadAll(r6)
	btn6 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b6), func() {
		showNews()
	})
	DeskBtn = widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(nil,nil , panelContent,  nil,img))
	})

	panelContent = container.NewVBox(container.NewGridWithRows(7, DeskBtn, btn1, btn2, btn3, btn4, btn5, btn6))

	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen()
	r, _ := fyne.LoadResourceFromPath("static\\oslogo.png")
	myWindow.SetIcon(r)
	myWindow.SetContent(
		container.NewBorder(nil,nil , panelContent,  nil,img),
	)
	myWindow.ShowAndRun()
}