package favicon

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var errFaviconNotFound = errors.New("favicon not found")

func getFavicon(targetUrl string) (string, error) {
	res, err := http.Get(targetUrl)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		err := fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	favicon, ok := doc.Find("link[rel='icon']").Attr("href")
	if !ok || favicon == "" {
		return "", errFaviconNotFound
	}
	return favicon, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	targetUrls, ok := r.URL.Query()["url"]
	if !ok || len(targetUrls) == 0 {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	targetUrl := targetUrls[0]

	faviconUrl, err := getFavicon(targetUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, faviconUrl, http.StatusSeeOther)
}
