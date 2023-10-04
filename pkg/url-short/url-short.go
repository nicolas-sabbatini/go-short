package urlShort

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
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

type pathUrlJsonOrYaml struct {
	Path string `json:"path" yaml:"path"`
	URL  string `json:"url" yaml:"url"`
}

func pathUrlJsonOrYamlSliceToMap(pu []pathUrlJsonOrYaml) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pu {
		log.Println("pathUrlJsonOrYamlSliceToMap: ", pu.Path, "to", pu.URL)
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

func parceYaml(yamlFile []byte) ([]pathUrlJsonOrYaml, error) {
	var parcedPathURLs []pathUrlJsonOrYaml
	err := yaml.Unmarshal(yamlFile, &parcedPathURLs)
	if err != nil {
		return nil, err
	}
	return parcedPathURLs, nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parcedPathURLs, err := parceYaml(yml)
	if err != nil {
		return nil, err
	}
	pathsToUrls := pathUrlJsonOrYamlSliceToMap(parcedPathURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

func parceJson(jsonFile []byte) ([]pathUrlJsonOrYaml, error) {
	var parcedPathURLs []pathUrlJsonOrYaml
	err := json.Unmarshal(jsonFile, &parcedPathURLs)
	if err != nil {
		return nil, err
	}
	return parcedPathURLs, nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parcedPathURLs, err := parceJson(json)
	if err != nil {
		return nil, err
	}
	pathsToUrls := pathUrlJsonOrYamlSliceToMap(parcedPathURLs)
	return MapHandler(pathsToUrls, fallback), nil
}
