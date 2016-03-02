package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func ImageHandler(res http.ResponseWriter, req *http.Request) {
	r, _ := strconv.Atoi(req.FormValue("r"))
	width, _ := strconv.Atoi(req.FormValue("width"))
	height, _ := strconv.Atoi(req.FormValue("height"))
	points, _ := strconv.Atoi(req.FormValue("points"))

	// rect := image.Rect(0, 0, width, height)
	// i := image.NewRGBA(rect)

	// for iy := 0; iy <= rect.Max.Y; iy += 1 {
	// 	for ix := 0; ix <= rect.Max.X; ix += 1 {
	// 		i.Set(ix, iy, img.At(int(float64(ix)/dx), int(float64(iy)/dy)))
	// 	}
	// }

	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < points; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)

		svg += fmt.Sprintf(
			`<circle cx="%d" cy="%d" r="%d" stroke="black" stroke-width="1" fill="azure" />`,
			x,
			y,
			r,
		)

		svg += fmt.Sprintf(
			`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="0.7"></line>`,
			x-(r*3),
			y,
			x+(r*3),
			y,
		)
		svg += fmt.Sprintf(
			`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="0.7"></line>`,
			x,
			y-(r*3),
			x,
			y+(r*3),
		)
		// }
	}

	for x := 0; x < width; x += base(float64(width/10), 5) {
		for y := 0; y < height; y += 10 {
			svg += fmt.Sprintf(
				`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="0.5px"></line>`,
				x,
				y,
				x,
				y+1, // y+5,
			)
		}
	}

	for y := 0; y < height; y += base(float64(height/10), 5) {
		for x := 0; x < width; x += 10 {
			svg += fmt.Sprintf(
				`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="0.5px"></line>`,
				x,
				y,
				x, // x+5,
				y+1,
			)
		}
	}

	svg += fmt.Sprintf(`</svg>`)

	res.Write([]byte(fmt.Sprintf(markup, svg)))
}
