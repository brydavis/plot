package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func HistogramHandler(res http.ResponseWriter, req *http.Request) {
	// max, _ := strconv.Atoi(req.FormValue("r"))

	width, _ := strconv.Atoi(req.FormValue("width"))
	height, _ := strconv.Atoi(req.FormValue("height"))
	tiles, _ := strconv.Atoi(req.FormValue("tiles"))

	div := fmt.Sprintf(`
		<div class="chart" style="height: %dpx; width: %dpx">`,
		height,
		width,
	)

	for i := 0; i < tiles; i++ {
		h := rand.Intn(height)
		div += fmt.Sprintf(
			`<div class="tile" style="height: %dpx; margin-top: %dpx"></div>`,
			h,
			height-h,
		)

	}

	div += fmt.Sprintf(`</div>
		<style> 
			.chart {
				/* position: absolute; */
				/* bottom: 0; */
				/* left: 0; */
			}

			.tile {
				border: 1px solid white; width: %d%%; float: left;
				background-color: grey;
				margin-left: 1px;
			}
			.tile:hover { background-color: black; }


		</style>`,

		(100/tiles)-1,
	)

	res.Write([]byte(fmt.Sprintf(markup, div)))
}
