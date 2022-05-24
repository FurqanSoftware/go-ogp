// Copyright 2022 Furqan Software Ltd. All rights reserved.

package ogp

import (
	"html/template"
	"strconv"
	"strings"
)

type OpenGraph struct {
	// Basic
	Title  string
	Type   string
	Image  URL // Ignored if Images is set.
	Images []Image
	URL    URL

	// Optional
	Audio           URL
	Description     string
	Determiner      Determiner
	Locale          Locale
	LocaleAlternate []Locale
	SiteName        string
	Video           URL
}

func (o OpenGraph) HTML() template.HTML {
	b := strings.Builder{}

	meta := func(property, content string) {
		b.WriteString(`<meta property="`)
		b.WriteString(template.HTMLEscaper(property))
		b.WriteString(`" content="`)
		b.WriteString(template.HTMLEscaper(content))
		b.WriteString(`" />`)
		b.WriteByte('\n')
	}

	if o.Title != "" {
		meta("og:title", o.Title)
	}
	if o.Type != "" {
		meta("og:type", o.Type)
	}
	if o.Images == nil && o.Image.IsValid() {
		meta("og:type", string(o.Image))
	}
	for _, g := range o.Images {
		if !g.URL.IsValid() {
			continue
		}
		if g.URL != "" {
			// Property og:image:url is identical to og:image.
			meta("og:image", string(g.URL))
		}
		if g.SecureURL != "" {
			meta("og:image:secure_url", string(g.SecureURL))
		}
		if g.Type != "" {
			meta("og:image:type", g.Type)
		}
		if g.Width > 0 {
			meta("og:image:width", strconv.Itoa(g.Width))
		}
		if g.Height > 0 {
			meta("og:image:height", strconv.Itoa(g.Height))
		}
		if g.Alt != "" {
			meta("og:image:alt", g.Alt)
		}
	}
	if o.URL.IsValid() {
		meta("og:url", string(o.URL))
	}

	if o.Audio != "" {
		meta("og:audio", string(o.Audio))
	}
	if o.Description != "" {
		meta("og:description", o.Description)
	}
	if o.Determiner != "" {
		meta("og:determiner", string(o.Determiner))
	}
	if o.Locale != "" {
		meta("og:locale", string(o.Locale))
	}
	for _, loc := range o.LocaleAlternate {
		meta("og:locale:alternate", string(loc))
	}
	if o.SiteName != "" {
		meta("og:site_name", o.SiteName)
	}
	if o.Video.IsValid() {
		meta("og:video", string(o.Video))
	}

	return template.HTML(b.String())
}

type Image struct {
	URL       URL
	SecureURL SecureURL
	Type      string
	Width     int
	Height    int
	Alt       string
}

type Video struct {
	URL       URL
	SecureURL SecureURL
	Type      string
	Width     int
	Height    int
}

type Audio struct {
	URL       URL
	SecureURL SecureURL
	Type      string
}

type URL string

func (u URL) IsValid() bool {
	return strings.HasPrefix(string(u), "http://") || strings.HasPrefix(string(u), "https://")
}

type SecureURL URL

func (s SecureURL) IsValid() bool {
	return strings.HasPrefix(string(s), "https://")
}

type Determiner string

const (
	DetBlank Determiner = ""
	DetAuto  Determiner = "auto"
	DetA     Determiner = "a"
	DetAn    Determiner = "an"
	DetThe   Determiner = "the"
)

type Locale string

const (
	LocEnUS Locale = "en_US"
)
