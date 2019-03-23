package utils

import (
	"fmt"
	"path/filepath"
	"statyk/internal/things"
	"time"

	"github.com/jor-go/sitemap"
)

// GenerateSitemap Adds a sitemap to the build directory
func GenerateSitemap(posts []things.Post, config things.SiteConfig, wd string) {
	var mySitemap sitemap.Sitemap

	homepage := sitemap.URL{}
	err := homepage.New(config.HomeLocation, "daily", 1.0, time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}

	mySitemap.AddURL(homepage)

	for _, post := range posts {
		var link sitemap.URL
		err := link.New(config.HomeLocation+"/"+post.Config.URL, "weekly", 0.5, time.Now())
		if err != nil {
			fmt.Println(err)
			continue
		}
		mySitemap.AddURL(link)
	}

	mySitemap.GenerateAndSave(filepath.Join(wd, "build", "sitemap.xml"))

	fmt.Println("Sitemap Generated: ", filepath.Join(wd, "build", "sitemap.xml"))
}
