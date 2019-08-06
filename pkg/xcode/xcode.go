package xcode

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	defaultAPIBase = "http://localhost:20343/api/"
	defaultAPIBots = "bots/"
)

// Client ...
type Client struct {
	http *http.Client
	base *sling.Sling
	opts *Opts

	Bot *Bot
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Base string
}

// New returns a new Xcode Client. If you provide an URL. For example,#
//
//		xcodeURL := "https://localhost"
// 		httpClient := &http.Client{}
// 		xcodeClient := chronos.New(chronosURL, httpClient)
//
// chronosURL: is the URL of your Chronos with a slash at the end.
func New(c *http.Client, opts ...Opt) *Client {
	options := new(Opts)

	p := new(Client)
	p.opts = options
	p.http = c

	configure(p, opts...)

	p.Bot = newBot(p.base)

	return p
}

// Base ...
func Base(path string) func(o *Opts) {
	return func(o *Opts) {
		o.Base = path
	}
}

func configure(p *Client, opts ...Opt) error {
	for _, o := range opts {
		o(p.opts)
	}

	p.base = sling.New().Client(p.http).Base(p.opts.Base)

	if p.opts.Base == "" {
		p.base = sling.New().Client(p.http).Base(defaultAPIBase)
	}

	return nil
}
