package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Divan struct {
	Name,
	Price,
	Size_width,
	Size_depth,
	Size_height,
	Sleep_width,
	Sleep_depth,
	Mechanism,
	Linen_drawer,
	Filler,
	Frame_material,
	Textile,
	Life_time,
	Armrests,
	Decorative_pillows,
	Guarantee,
	Configuration,
	Weight,
	Load,
	url string
}

func main() {
	c := colly.NewCollector()
	c2 := colly.NewCollector()

	divans := []Divan{}
	divan := Divan{}

	c.OnHTML("a.product__img", func(h *colly.HTMLElement) {
		c2.Visit(h.Request.AbsoluteURL(h.Attr("href")))
		divan.url = h.Request.AbsoluteURL(h.Attr("href"))
		divans = append(divans, divan)
	})
	// c.OnHTML("div.catalog-list-section-pagination", func(h *colly.HTMLElement) {
	// 	h.ChildAttr("div.btn", "onclick"), func(int, *HTMLElement)
	// })
	c2.OnHTML("div.item-info__col > div.item-info__specs > ul > li", func(h *colly.HTMLElement) {
		selection := h.DOM
		childNodes := selection.Children().Nodes
		Char := strings.TrimSpace(selection.FindNodes(childNodes[0]).Text())
		Value_char := strings.TrimSpace(selection.FindNodes(childNodes[1]).Text())
		// 	//fmt.Println(Char, Value_char)
		// 	// currentTime := time.Now()

		// 	// fmt.Println("Current Time in String: ", currentTime.String())
		switch Char {
		case "Габариты (ШхГхВ)":
			size_raw := strings.Split(Value_char, " ")[0]
			divan.Size_width = strings.Split(size_raw, "x")[0]
			divan.Size_depth = strings.Split(size_raw, "x")[1]
			divan.Size_height = strings.Split(size_raw, "x")[2]
		case "Спальное место (ШхГ)":
			size_raw := strings.Split(Value_char, " ")[0]
			divan.Sleep_width = strings.Split(size_raw, "x")[0]
			divan.Sleep_depth = strings.Split(size_raw, "x")[1]
		case "Механизм трансформации":
			divan.Mechanism = Value_char
		case "Наполнитель":
			divan.Filler = Value_char
		case "Бельевой ящик":
			divan.Linen_drawer = Value_char
		case "Материал каркаса":
			divan.Frame_material = Value_char
		case "Ткань":
			divan.Textile = Value_char
		case "Гарантия":
			guar := strings.Split(Value_char, " ")[0]
			divan.Guarantee = guar
			fmt.Println(guar)
		case "Срок службы":
			srok := strings.Split(Value_char, " ")[0]
			divan.Life_time = srok
			fmt.Println(srok)
		case "Подлокотники":
			divan.Armrests = Value_char
		case "Декоративные подушки":
			divan.Decorative_pillows = Value_char
		case "Конфигурация":
			divan.Configuration = Value_char
		case "Вес":
			wes := strings.Split(Value_char, " ")[0]
			divan.Weight = wes
			fmt.Println(wes)
		case "Нагрузка":
			lod := strings.Split(Value_char, " ")[0]
			divan.Load = lod
			fmt.Println(lod)
		}
		//divans = append(divans, divan)
		//fmt.Println(divans)

	})

	c2.OnHTML("div.item-header__inner > div.description-container > h1", func(h *colly.HTMLElement) {
		name := h.Text
		noSpaces := strings.TrimSpace(name)
		divan.Name = noSpaces

	})
	c2.OnHTML("div.item-header__info.price-type__1 > div.item-header__prices > p:nth-child(3) > span:nth-child(1)", func(h *colly.HTMLElement) {
		price := h.Text
		noSpaces := strings.ReplaceAll(price, " ", "")
		SPACE := strings.TrimSpace(noSpaces)
		divan.Price = SPACE

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.Visit("https://mnogomebeli.com/divany/")

	// fmt.Println(divans)

	file, err := os.Create("Parser_data.csv")
	if err != nil {
		log.Fatalln("Failed to create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Название", "Ссылка", "Габариты(ширина)", "Габариты(глубина)", "Габариты(высота)", "Спальная ширина", "Спальная глубина", "Механизм", "Бельевой ящик", "Наполнитель", "Материал каркаса", "Ткань", "Подлокотники", "Декоративные подушки", "Срок службы(годы)", "Гарантия(месяцы)", "Конфигурация", "Вес(кг)", "Нагрузка(кг)", "Цена(рубли)"})

	for _, value := range divans {
		writer.Write([]string{value.Name, value.url, value.Size_width, value.Size_depth, value.Size_height, value.Sleep_width, value.Sleep_depth, value.Mechanism, value.Linen_drawer, value.Filler, value.Frame_material, value.Textile, value.Armrests, value.Decorative_pillows, value.Life_time, value.Guarantee, value.Configuration, value.Weight, value.Load, value.Price})
	}
}
