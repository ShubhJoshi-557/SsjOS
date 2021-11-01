package main

import (
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2/canvas"
)

func showGalleryApp() {
	// a := app.New()
	
	root_src:="C:\\Users\\shubh\\OneDrive\\Desktop\\Wallpapers"
	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }
	tabs:= container.NewAppTabs()
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extensions:= strings.Split(file.Name(),".")[1];
			if extensions == "png" || extensions == "jpeg" || extensions == "jpg"{
				image:= canvas.NewImageFromFile(root_src+"\\"+file.Name())
				image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(),image))
			}
		}
    }

	tabs.SetTabLocation(container.TabLocationLeading)
	w := myapp.NewWindow("ImageViewer")
	w.Resize(fyne.NewSize(800,600));
	w.SetContent(tabs);
	w.Show();
}

