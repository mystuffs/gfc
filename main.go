package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const prefix = "#"

type JSONOutput struct {
	Followers uint64 `json:"followers"`
	Message   string `json:"message"`
}

func makeReq(w http.ResponseWriter, r http.Request, user string, left_color, left_t_color, left_o_color, right_color, right_t_color, right_o_color string) string {
	filtered_username := strings.ReplaceAll(user, "/", "")
	body, _ := http.Get("https://api.github.com/users/" + filtered_username)
	var out JSONOutput
	util, _ := io.ReadAll(body.Body)

	json.Unmarshal([]byte(string(util)), &out)
	if len(out.Message) != 0 {
		http.Error(w, "GFC has hit GitHub API ratelimit, if you're using a vps, please try again with a VPN connection", http.StatusNotAcceptable)
		return ""
	}
	fmt.Println(out.Followers)
	make_svg(w, &r, NearestThousandFormat(float64(out.Followers)),
		prefix+left_color, prefix+left_t_color, prefix+left_o_color, prefix+right_color, prefix+right_t_color, prefix+right_o_color)

	return strconv.FormatUint(out.Followers, 10)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		left_color := r.URL.Query().Get("left_color")
		right_color := r.URL.Query().Get("right_color")
		left_t_color := r.URL.Query().Get("left_txt_color")
		left_o_color := r.URL.Query().Get("left_op_color")
		right_t_color := r.URL.Query().Get("right_txt_color")
		right_o_color := r.URL.Query().Get("right_op_color")

		if len(username) == 0 {
			http.Error(w, "Please provide a username", http.StatusBadRequest)
			return
		}
		if len(left_color) == 0 {
			left_color = "555"
		}
		if len(right_color) == 0 {
			right_color = "bf51d3"
		}
		if len(left_o_color) == 0 && len(right_o_color) == 0 {
			left_o_color = "010101"
			right_o_color = left_o_color
		}
		if len(left_t_color) == 0 && len(right_t_color) == 0 {
			left_t_color = "fff"
			right_t_color = left_t_color
		}
		makeReq(w, *r, username, left_color, left_t_color, left_o_color, right_color, right_t_color, right_o_color)
	})

	http.ListenAndServe("localhost:1337", nil)
}
