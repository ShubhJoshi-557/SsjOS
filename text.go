// package main

// import (
// 	"io/ioutil"
// 	"strconv"

// 	"fyne.io/fyne/v2"
// 	// "fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/storage"
// 	"fyne.io/fyne/v2/widget"
// ) 
// var count int = 1
// func showTextEditor() {
// 	w:= myapp.NewWindow("SSJ Editor")
// 	w.Resize(fyne.NewSize(600,600))
	
// 	content:=container.NewVBox(
// 		container.NewVBox(
// 			widget.NewLabel("SSJ Text Editor"),
// 		),
// 	)
// 	content.Add(widget.NewButton("Add New File", func(){
// 		content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
// 		count++
// 	}))

// 	input:= widget.NewMultiLineEntry()
// 	input.SetPlaceHolder("Enter Text...")
// 	input.Resize(fyne.NewSize(400,400))
// 	saveBtn:= widget.NewButton("Save Text File", func() {
// 		saveFileDialog:=dialog.NewFileSave(
// 			func(uc fyne.URIWriteCloser, _ error){
// 				textData:= []byte(input.Text)
// 				uc.Write(textData)
// 			},w)

// 		saveFileDialog.SetFileName("New File " + strconv.Itoa(count-1)+".txt")
// 		saveFileDialog.Show()
// 	})

// 	openBtn:= widget.NewButton("Open File", func() {
// 		openFileDialog:= dialog.NewFileOpen(
// 			func(r fyne.URIReadCloser, _ error){
// 				ReadData ,_:=ioutil.ReadAll(r)

// 				output:=fyne.NewStaticResource("New File", ReadData)

// 				viewData:=widget.NewMultiLineEntry()
// 				viewData.SetText(string(output.StaticContent))
// 				w:= fyne.CurrentApp().NewWindow(
// 					string(output.StaticName))
// 					w.SetContent(container.NewScroll(viewData))
// 					w.Resize(fyne.NewSize(400,400))
// 					w.Show()
// 			},w)

// 			openFileDialog.SetFilter(
// 				storage.NewExtensionFileFilter([] string{".txt"}))
// 			openFileDialog.Show()
// 	})

// 	textContainer:=container.NewVBox(
// 			content,
// 			input,
// 			container.NewHBox(
// 				saveBtn,
// 				openBtn,
// 			),
// 		)
		
	
// 	w.SetContent(container.NewBorder(nil,nil,nil,nil,textContainer))
// 	w.Show()
// }

package main
import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "fyne.io/fyne/v2"
    // "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)
var count int = 1
var filepath string
func showTextEditor() {
    // a := app.New()
    w := myapp.NewWindow("Ssj NotePad")
	w.Resize(fyne.NewSize(500, 500))
    input := widget.NewMultiLineEntry()
    input.SetPlaceHolder("Enter text...")
    input.Move(fyne.NewPos(0, 0))
    input.Resize(fyne.NewSize(500, 500))
    new1 := fyne.NewMenuItem("New", func() {
        filepath = ""
        w.SetTitle("Ssj NotePad")
        input.Text = ""
        input.Refresh()
    })
    save1 := fyne.NewMenuItem("Save", func() {
        if filepath != "" {
            f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
            if err != nil {
                //handle error
            }
            defer f.Close()
            //f.Write([]byte(input.Text))
            f.WriteString(input.Text)
        } else {
            saveFileDialog := dialog.NewFileSave(
                func(r fyne.URIWriteCloser, _ error) {
                    textData := []byte(input.Text)
                    r.Write(textData)
                    filepath = r.URI().Path()
                    w.SetTitle(filepath)
                }, w)
            saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
            saveFileDialog.Show()
        }
    })
    saveAs1 := fyne.NewMenuItem("Save as..", func() {
        saveFileDialog := dialog.NewFileSave(
            func(r fyne.URIWriteCloser, _ error) {
                textData := []byte(input.Text)
                r.Write(textData)
                filepath = r.URI().Path()
                w.SetTitle(filepath)
            }, w)
        saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
        saveFileDialog.Show()
    })
    open1 := fyne.NewMenuItem("Open", func() {
        openfileDialog := dialog.NewFileOpen(
            func(r fyne.URIReadCloser, _ error) {
                data, _ := ioutil.ReadAll(r)
                result := fyne.NewStaticResource("name", data)
                input.SetText(string(result.StaticContent))
                fmt.Println(result.StaticName + r.URI().Path())
                filepath = r.URI().Path()
                w.SetTitle(filepath)
            }, w)
        openfileDialog.SetFilter(
            storage.NewExtensionFileFilter([]string{".txt"}))
        openfileDialog.Show()
    })
    menuItem := fyne.NewMenu("File", new1, save1, saveAs1, open1)
    menux1 := fyne.NewMainMenu(menuItem)
    w.SetMainMenu(menux1)
    textContainer:=container.NewWithoutLayout(input,)
    r, _ := fyne.LoadResourceFromPath("static\\notepadlogo.png")
	w.SetIcon(r)
    w.SetContent(container.NewBorder(nil,nil,nil,nil,textContainer))
    w.Show()
}
