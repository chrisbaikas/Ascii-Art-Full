// Initializes and loads banner fonts from the 'banners' folder

package web

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"platform.zone01.gr/git/askordal/ascii-art-web-export-file/utils"
)

var LoadedBanners map[string]map[rune][]string

func init() {
	LoadedBanners = make(map[string]map[rune][]string)
	files, err := os.ReadDir("banners")
	if err != nil {
		log.Fatalf("error reading banners dir: %v", err)
	}

	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".txt") {
			continue
		}
		name := strings.TrimSuffix(f.Name(), ".txt")
		path := filepath.Join("banners", f.Name())
		bannerMap, err := utils.LoadBanner(path)
		if err != nil {
			log.Fatalf("error loading banner %s: %v", name, err)
		}
		LoadedBanners[name] = bannerMap
		log.Printf("Loaded banner: %s", name)
	}
}
