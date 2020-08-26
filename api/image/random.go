package image

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	randomImageUrl := "https://source.unsplash.com/random"
	http.Redirect(w, r, randomImageUrl, http.StatusSeeOther)
}
