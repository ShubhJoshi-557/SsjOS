package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
	"net/http"
)

func showWeatherApp() {
	// a := app.New()
	w := myapp.NewWindow("Weather App")
	w.Resize(fyne.NewSize(400,500))
	r, _ := fyne.LoadResourceFromPath("static\\weather3.png")
	w.SetIcon(r)
	
	
	
	image:= canvas.NewImageFromFile("static\\weather3.png")
	image.FillMode = canvas.ImageFillOriginal
	combo := widget.NewSelect([]string{"Mumbai", "Delhi", "Ahmedabad", "Chennai", "Bangalore", "Kolkata",
		"Surat",
		"Pune",
		"Thane",
		"Hyderabad",
		"Jaipur",
		"Lucknow",
		"Kanpur",
		"Nagpur",
		"Indore",
		"Bhopal",
		"Amritsar",
		"Ranchi",
		"Gurgaon",
		"Noida",}, func(value string) {
		log.Println("Select set to", value)
		res, err:=http.Get("http://api.openweathermap.org/data/2.5/weather?q="+value+"&APPID=b2d327a4ff36fdcee7cd18945283e86f")
		if err!=nil{
			fmt.Println(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err!=nil{
			fmt.Println(err)
		}

		weather, err := UnmarshalWeather(body)
		if err!=nil{
			fmt.Println(err)
		}

		label1:= canvas.NewText("Weather Details:", color.White) 
		label1.TextStyle = fyne.TextStyle{Bold: true}
		label2:= canvas.NewText(fmt.Sprintf("Country: \t%s" , weather.Sys.Country), color.White)
		label3:= canvas.NewText(fmt.Sprintf("City: \t%s" , weather.Name), color.White) 
		label4:= canvas.NewText(fmt.Sprintf("Wind Speed: \t%.2f" , weather.Wind.Speed), color.White) 
		label5:= canvas.NewText(fmt.Sprintf("Temperature: \t%.2f" , weather.Main.Temp), color.White) 
		label6:= canvas.NewText(fmt.Sprintf("Humidity: \t%.2d" , weather.Main.Humidity), color.White) 

		weatherContainer:=container.NewVBox(
				image,
				label1,
				// combo,
				label2,
				label3,
				label4,
				label5,
				label6,
		)
		w.SetContent(container.NewBorder(nil,nil,nil,nil,weatherContainer),)
	})
	weatherContainer:=container.NewVBox(
			image,
			combo,
	)
	w.SetContent(container.NewBorder(nil,nil,nil,nil,weatherContainer),)
	
	w.Show()
}


// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`     
	Weather    []WeatherElement `json:"weather"`   
	Base       string           `json:"base"`      
	Main       Main             `json:"main"`      
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`      
	Clouds     Clouds           `json:"clouds"`    
	Dt         int64            `json:"dt"`        
	Sys        Sys              `json:"sys"`       
	Timezone   int64            `json:"timezone"`  
	ID         int64            `json:"id"`        
	Name       string           `json:"name"`      
	Cod        int64            `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type WeatherElement struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}
