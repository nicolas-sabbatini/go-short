package urlShort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("MapHandler: ", r.URL.Path)
		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			log.Println("MapHandler: redirecting to ", dest)
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type yamlPathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parcedPathURLs []yamlPathURL
	err := yaml.Unmarshal(yml, &parcedPathURLs)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range parcedPathURLs {
		log.Println("YAMLHandler: ", pu.Path, "to", pu.URL)
		pathsToUrls[pu.Path] = pu.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}
