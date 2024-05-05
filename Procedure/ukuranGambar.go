package main

import "fmt"

func main() {
	var p, l, skala, w, h int
	var simbol string

	fmt.Scan(&p, &l)
	fmt.Scan(&simbol, &skala)

	if simbol == "+" && skala >= 1 && skala <= 100 {
		zoomIn(p, l, skala, &w, &h)
	} else if simbol == "-" && skala >= 1 && skala <= 100 {
		zoomOut(p, l, skala, &w, &h)
	}

	fmt.Println(w, h)

}

func zoomIn(p int, l int, skala int, w *int, h *int) {
	*w = p * skala
	*h = l * skala
}

func zoomOut(p int, l int, skala int, w *int, h *int) {
	*w = p / skala
	*h = l / skala
}
