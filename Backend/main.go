package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type PageData struct {
	Time            string
	Estimated_price string
}

func main() {

	handleReq()
}

func pageHome(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	fmt.Println(r.Method)
	updateTime := time.Now().Format("15:04:05")
	data.Time = updateTime
	//fmt.Fprint(w)

	weight := r.FormValue("weight")
	weight_inte, _ := strconv.Atoi(weight)

	if weight_inte > 100 {
		data.Estimated_price = "40000"
	} else if weight_inte <= 100 && weight_inte > 0 {
		data.Estimated_price = "20000"
	} else if weight_inte < 0 {
		data.Estimated_price = "Число не может быть меньше или равно нулю"
	}

	fmt.Println(weight)

	tmpl, _ := template.ParseFiles("hehe.html")
	tmpl.Execute(w, data)

}

func submitFor(w http.ResponseWriter, r *http.Request) {
	// title := r.FormValue("name")
	// urla := r.FormValue("url")
	// size_width := r.FormValue("size_width")
	// size_depth := r.FormValue("size_depth")
	// size_height := r.FormValue("size_height")
	// sleep_Width := r.FormValue("sleep_Width")
	// sleep_Depth := r.FormValue("sleep_Depth")
	// mechanism := r.FormValue("mechanism")
	// linen_drawer := r.FormValue("linen_drawer")
	// filler := r.FormValue("filler")
	// frame_Material := r.FormValue("frame_Material")
	// textile := r.FormValue("textile")
	// armrests := r.FormValue("armrests")
	// decorative_Pillows := r.FormValue("decorative_Pillows")
	// lifetime := r.FormValue("lifetime")
	// warranty := r.FormValue("warranty")
	// configuration := r.FormValue("configuration")
	weight := r.FormValue("weight")
	// load := r.FormValue("load")
	fmt.Println(weight)
}
func handleReq() {
	http.HandleFunc("/", pageHome)
	http.HandleFunc("/submitForm", submitFor)
	http.ListenAndServe(":8080", nil)
}
