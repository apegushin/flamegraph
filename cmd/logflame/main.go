package main

import (
	"fmt"
	"net/http"
)

func SvgTextBox(x, y, rx, ry, width, fontSize, textPadding int, text string) (res string) {
	// groups SVG rectangle and tries to fit some text into it
	height := y + 2*textPadding + fontSize
	res += "<g>"
	res += fmt.Sprintf("<rect x=\"%d\" y=\"%d\" rx=\"%d\" ry=\"%d\" width=\"%d\" height=\"%d\" stroke=\"%s\" fill=\"%s\" stroke-width=\"%d\" />",
		x, y, rx, ry, width, y+2*textPadding+fontSize, "black", "transparent", 1)
	text_x, text_y := x+4, height-textPadding
	res += fmt.Sprintf("<text x=\"%d\" y=\"%d\" font-size=\"%d\">%s</text>", text_x, text_y, fontSize, text)
	res += "/<g>"
	return res

	// groupopen := `<g>`
	// rect1 := `<rect x="10" y="10" width="50" height="30" stroke="black" fill="transparent" stroke-width="2"/>`
	// rect1name := `<text x="10" y="28" font-size="10">Func Name</text>`
	// groupclose := `</g>`
	// rect2 := `<rect x="70" y="10" width="50" height="30" stroke="black" fill="transparent" stroke-width="2"/>`
	// svgclose := `</svg>`
}

func greet(w http.ResponseWriter, r *http.Request) {
	header := `<html><body>`
	footer := `</body> </html>`
	title := `<h1>My first SVG</h1>`
	svgopen := `<svg width="1000" height="1000">`
	svgclose := `</svg>`
	page := header + title + svgopen + SvgTextBox(2, 2, 4, 4, 100, 14, 8, "there you go!") + svgclose + footer
	fmt.Fprintf(w, page)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
