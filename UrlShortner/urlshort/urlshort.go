package urlshort

import (
	"errors"
	"net/http"
	//yaml "gopkg.in/yaml.v2"
)

//HelloFromPackage is test func
func HelloFromPackage(name string) string {
	return "Hello " + name + ", from the package!"
}

// MapHandler will return a http.HandlerFunc
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathURLs)
	return MapHandler(pathToUrls, fallback), nil
}

func parseYAML(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := errors.New("Test err") //yaml.Unmarshal(data, pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathUrls []pathURL) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
