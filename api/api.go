package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const homeUrl = "https://bandcamp.com/api/discover/3/get_web"

type queryArgs struct {
	g             string
	s             string
	p             int
	gn            int
	f             string
	w             int
	lo            bool
	lo_action_url string
}

type FeaturedTrack struct {
	File     FeaturedTrackFile `json:"file"`
	Duration float64           `json:"duration"`
}

type FeaturedTrackFile struct {
	Link string `json:"mp3-128"`
}

type Release struct {
	Type          string        `json:"type"`
	Id            int           `json:"id"`
	IsPreorder    int           `json:"is_preorder"`
	BandId        int           `json:"band_id"`
	PublishDate   string        `json:"publish_date"`
	Genre         string        `json:"genre_text"`
	Album         string        `json:"primary_text"`
	Artist        string        `json:"secondary_text"`
	FeaturedTrack FeaturedTrack `json:"featured_track"`
	UrlHints      UrlHints      `json:"url_hints"`
}

type UrlHints struct {
	Subdomain string `json:"subdomain"`
	Slug      string `json:"Slug"`
}

type Releases struct {
	Items []Release `json:"items"`
}

func NewQueryArgs(page int) *queryArgs {
	q := &queryArgs{"all", "top", page, 0, "all", 0, true, "https://bandcamp.com"}
	return q
}

func (q *queryArgs) ToValues() url.Values {
	return url.Values{
		"g":             {q.g},
		"s":             {q.s},
		"p":             {strconv.Itoa(q.p)},
		"gn":            {strconv.Itoa(q.gn)},
		"f":             {q.f},
		"w":             {strconv.Itoa(q.w)},
		"lo":            {strconv.FormatBool(q.lo)},
		"lo_action_url": {q.lo_action_url},
	}
}

func (q *queryArgs) String() string {
	return q.ToValues().Encode()
}

func FetchReleasesFromHome(q *queryArgs) (releases Releases, err error) {
	url := homeUrl + "?" + q.ToValues().Encode()
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &releases)
	if err != nil {
		return releases, err
	}

	return releases, err
}
