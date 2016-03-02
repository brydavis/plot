package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func GridHandler(res http.ResponseWriter, req *http.Request) {
	r, _ := strconv.Atoi(req.FormValue("r"))
	r = base(float64(r), 2)
	width, _ := strconv.Atoi(req.FormValue("width"))
	height, _ := strconv.Atoi(req.FormValue("height"))
	tiles, _ := strconv.Atoi(req.FormValue("tiles"))

	div := fmt.Sprintf(`
		<div style="height: %dpx; width: %dpx">`,
		height,
		width,
	)

	for i := 0; i < tiles; i++ {
		div += fmt.Sprintf(
			`<div class="tile"></div>`,
		)

	}

	div += fmt.Sprintf(`</div>
		<style> 
			.tile {
				border: 1px solid white; height: %dpx; width: %dpx; float: left;
				background-color: grey;
			}
			.tile:hover { background-color: black; }


		</style>`,

		r,
		r,
	)

	res.Write([]byte(fmt.Sprintf(markup, div)))
}
