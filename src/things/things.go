package things

import (
	"encoding/xml"
	"html/template"
)

/*Flags : Flags passed in through the terminal*/
type Flags struct {
	Name string
}

/*SiteConfig : Config for the Site*/
type SiteConfig struct {
	Name          string `yaml:"name,omitempty"`
	StyleLocation string `yaml:"style-location,omitempty"`
	HomeLocation  string `yaml:"home-location,omitempty"`
	Port          string `yaml:"port"`
}

/*PostConfig : Configuration for a post*/
type PostConfig struct {
	Title       string            `yaml:"title"`
	Date        string            `yaml:"date"`
	URL         string            `yaml:"url"`
	Markdown    string            `yaml:"markdown"`
	Description string            `yaml:"description"`
	Custom      map[string]string `yaml:"custom"`
}

/*Post : Defines a post*/
type Post struct {
	SiteConfig SiteConfig
	Config     PostConfig
	HTML       template.HTML
}

/*MultiPost : A page with more than one post*/
type MultiPost struct {
	SiteConfig
	Posts []Post
}

/*GetPostsAfter : Returns posts after index of*/
func (m *MultiPost) GetPostsAfter(i int) []Post {
	return m.Posts[i:]
}

/*SitemapURL : The struct representing the xml for a single link*/
type SitemapURL struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod"`
	ChangeFreq string `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

/*Sitemap : Struct for the entire sitemap*/
type Sitemap struct {
	Format  string       `xml:",innerxml"`
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLS    []SitemapURL `xml:"url"`
}
