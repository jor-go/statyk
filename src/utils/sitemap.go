package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"statyk/src/things"
	"time"
)

/*GenerateSitemap : Adds a sitemap to the build directory*/
func GenerateSitemap(posts []things.Post, config things.SiteConfig, wd string) {
	var sitemap things.Sitemap

	homepage := things.SitemapURL{
		Loc:        config.HomeLocation,
		ChangeFreq: "daily",
		LastMod:    time.Now().Format("2006-01-02"),
		Priority:   "1.0",
	}

	sitemap.URLS = append(sitemap.URLS, homepage)

	for _, post := range posts {
		var link things.SitemapURL
		link.Loc = config.HomeLocation + "/" + post.Config.URL
		link.ChangeFreq = "weekly"
		link.LastMod = time.Now().Format("2006-01-02")
		link.Priority = "0.05"

		sitemap.URLS = append(sitemap.URLS, link)
	}

	sitemap.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

	data, err := xml.Marshal(sitemap)

	if err != nil {
		log.Println(err)
	}

	header := []byte(xml.Header)

	final := append(header, data...)

	err = ioutil.WriteFile(filepath.Join(wd, "build", "sitemap.xml"), final, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Sitemap Generated: ", filepath.Join(wd, "build", "sitemap.xml"))
}
