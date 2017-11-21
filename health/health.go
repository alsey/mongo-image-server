package health

import(
	"os"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
)

var envs map[string]string

func init() {
	envs = make(map[string]string)
	for _, item := range os.Environ() {
		envs[strings.Split(item, "=")[0]] = strings.Split(item, "=")[1]
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string {
		"status" : "UP",
	})
}

func Env(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(envs)
}

func Favicon(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "image/png")
  fmt.Fprint(w, `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAcxJREFUeNp80t8r3XEcx/Hv8SM/IhdL6UzRUGTHj3NWqBEJR201NwvHBeUPsNoFN1smkhtKWhQXxG4ou9glSrK2aaE1P7O7JS5YGykcnu96fXV241OPvp/v9/P5vj/vz+f98QQCgZDjOL2IwzZmMYm/zv8tBk1ohg9RGIz2er1zdB4hCb+QgHb9tKlnEEN4iBuUIBl+i3qtSUd4jjNkoxvlOMQTvMU3zd1AAS4tjVV9tMge9ffRahPwBq8ifnb03dpPC/BBL6koVt+LNSyiE8so05htN6D+jJ2B7btCA3/wFSs41TaWEIt0ZZuhDBbw3kMVHB3KR63WhjzkKOC9LUrPL6hGnbKxrPwqbbxWTZOMiLO6C2CtBo22LzzDjipklRrQHTnAGPLdIG6AQvTjXOU6xm9c4R+mkKLvE5jHCzeARytYunMq4Wu80+FZa9B5BLWNLPTYpbP0KlHrlkXPRa1iZ/NAF6weWwjjQtt4aRmE9NNJxGX5jlyUYlyLfNbYns7CWosFeKoXN7J7qfqQiRF0KW13Xlh9n20hUS+W6jR28Rif0KExq86wVk/SuLUEu4l2q4pUATvQdVViKaLEP3TR7A5UqbQ2f/RWgAEA6xlnphy+IP4AAAAASUVORK5CYII=`)
}