package main

import (
	"fmt"
	"net/http"
)

func make_svg(
	w http.ResponseWriter,
	r *http.Request,
	followers, left_color, left_t_color, left_o_color, right_color, right_t_color, right_o_color string) {
	w.Header().Add("cache-control", "max-age=0, no-cache, no-store, must-revalidate")
	w.Header().Add("content-type", "image/svg+xml")
	width := (12 + (len(followers) * 5)) + 3
	rect_x := 35 + (12 + len(followers)*3) + 3
	x := 650 + (12 + (len(followers) * 10)) + 3
	switch len(followers) {
	case 4:
		x = 680 + (12 + (len(followers) * 10)) + 3
	case 5:
		x = 700 + (12 + (len(followers) * 10)) + 3
	case 6:
		x = 720 + (12 + (len(followers) * 10)) + 3
	}

	fmt.Fprintf(w, `
		<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="150" height="20" role="img" aria-label="followers: 165k">
		<title>followers: 165k</title>
		<linearGradient id="s" x2="0" y2="100%%">
			<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
			<stop offset="1" stop-opacity=".1"/>
		</linearGradient>
		<clipPath id="r">
			<rect width="960" height="20" rx="3" fill="#fff"/>
		</clipPath>
		<g clip-path="url(#r)">
			<rect width="59" height="20" fill="%s"/>
			<rect x="59" width="%d" height="20" fill="%s"/>
			<rect width="%d" height="20" fill="url(#s)"/>
		</g>
		<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110">
			<text aria-hidden="true" x="305" y="150" fill="%s" fill-opacity=".3" transform="scale(.1)" textLength="490">Followers</text>
			<text x="305" y="140" transform="scale(.1)" fill="%s" textLength="490">Followers</text>
			<text aria-hidden="true" x="%d" y="150" fill="%s" fill-opacity=".3" transform="scale(.1)">%s</text>
			<text x="%d" y="140" transform="scale(.1)" fill="%s">%s</text>
		</g>
		</svg>
	`, left_color,
		width,
		right_color,
		rect_x,
		left_o_color,
		left_t_color,
		x,
		right_o_color,
		followers,
		x,
		right_t_color,
		followers,
	)
}
