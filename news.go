package main
import (
    "encoding/json"
    "fmt"
    "image/color"
    "io/ioutil"
    "net/http"
    "fyne.io/fyne/v2"
    // "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)
var num_article int = 1
var news News
func showNews() {
    // a := app.NewWithID("")
    w := myapp.NewWindow("SSJ News App")
    w.Resize(fyne.NewSize(300, 400))
    URL := fmt.Sprintf("https://gnews.io/api/v4/top-headlines?token=f4242aaddbb109f07654431e207e2412&lang=en&max=100")
    //API
    res, _ := http.Get(URL)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    news, _ = UnmarshalNews(body)
    fmt.Println(news)
    // num_article = int(news.TotalArticles)
    label2 := widget.NewLabel(fmt.Sprintf("No of articles:%d",
        news.TotalArticles))
    //show title
    label3 := widget.NewLabel(fmt.Sprintf("%s", news.Articles[1].Title))
    label3.TextStyle = fyne.TextStyle{Bold: true}
    label3.Wrapping = fyne.TextWrapBreak
    // show articles
    entry1 := widget.NewLabel(fmt.Sprintf("%s",news.Articles[1].Description))
    //entry1.MultiLine = true
    entry1.Wrapping = fyne.TextWrapBreak
    //entry1.Resize(fyne.NewSize(300, 400))
    //entry1.Resize(fyne.NewSize(300, 300))
    // label4 := widget.NewLabel(fmt.Sprintf("Article: %s", news.Articles[0].Description))
    btn := widget.NewButton("Next", func() {
        num_article += 1
        label3.Text = news.Articles[num_article].Title
        entry1.Text = news.Articles[num_article].Description
        label3.Refresh()
        entry1.Refresh()
    })
    label1 := canvas.NewText("-- News App --", color.White)
    label1.Alignment = fyne.TextAlignCenter
    label1.TextStyle = fyne.TextStyle{Bold: true}
    img := canvas.NewImageFromFile("static\\news1.png")
    // img.Resize(fyne.NewSize(150, 150))
    img.FillMode = canvas.ImageFillOriginal
    e := container.NewVBox(label1, label3, entry1, btn)
	r, _ := fyne.LoadResourceFromPath("static\\newslogo.png")
	w.SetIcon(r)
    // e.Resize(fyne.NewSize(300, 300))
    w.SetContent(container.NewBorder(img, label2, nil, nil, e))
    w.CenterOnScreen()
    w.Show()
}
// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    news, err := UnmarshalNews(bytes)
//    bytes, err = news.Marshal()
func UnmarshalNews(data []byte) (News, error) {
    var r News
    err := json.Unmarshal(data, &r)
    return r, err
}
func (r *News) Marshal() ([]byte, error) {
    return json.Marshal(r)
}
type News struct {
    TotalArticles int64     `json:"totalArticles"`
    Articles      []Article `json:"articles"`
}
type Article struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Content     string `json:"content"`
    URL         string `json:"url"`
    Image       string `json:"image"`
    PublishedAt string `json:"publishedAt"`
    Source      Source `json:"source"`
}
type Source struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}
