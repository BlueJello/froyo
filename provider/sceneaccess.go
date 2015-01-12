package provider

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type SceneAccess struct {
	cookiejar *cookiejar.Jar
	client    *http.Client
}

func NewSceneAccessTracker() (*SceneAccess, error) {
	cj, err := cookiejar.New(nil)

	if err != nil {
		return nil, err
	}
	sa := &SceneAccess{
		client: &http.Client{
			Jar: cj,
		},
	}
	return sa, nil
}
func (sa *SceneAccess) LoginRequired() bool {
	return true
}
func (sa *SceneAccess) Login(username, password string) error {
	_, err := sa.client.PostForm("https://sceneaccess.eu/login",
		url.Values{"username": {username}, "password": {password}})

	if err != nil {
		return err
	}

	return nil
}

func (sa *SceneAccess) Search(title string) ([]*SearchResult, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   "sceneaccess.eu",
		Path:   "browse",
		RawQuery: url.Values{
			"search": {title},
			"method": {"2"},
			"sort":   {"6"},
			"type":   {"desc"},
		}.Encode(),
	}

	resp, err := sa.client.Get(url.String())
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	results := []*SearchResult{}
	doc.Find("#torrents-table tbody .tt_row").Each(func(i int, s *goquery.Selection) {
		results = append(results, &SearchResult{
			Name: s.Find(".ttr_name a").Text(),
			Size: 1,
		})

	})
	return results, nil
}
