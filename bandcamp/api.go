package bandcamp

import (
	"fmt"
	"net/url"
	"strconv"
)

const home_url = "https://bandcamp.com/api/discover/3/get_web"

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

func NewQueryArgs(page int) url.Values {
	q := &queryArgs{"all", "top", page, 0, "all", 0, true, "https://bandcamp.com"}
	v := q.ToValues()
	return v
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
	return fmt.Sprintf("g=%s&s=%s&p=%d&gn=%d&f=%s&w=%d&lo=%t&lo_action_url=%s", q.g, q.s, q.p, q.gn, q.f, q.w, q.lo, q.lo_action_url)
}
