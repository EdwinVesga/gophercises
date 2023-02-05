package handler

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		if v, ok := pathsToUrls[path]; ok {
			http.Redirect(rw, req, v, http.StatusFound)
		} else {
			fallback.ServeHTTP(rw, req)
		}
	}
}

type Urlshort struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)

	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]Urlshort, error) {
	arr := make([]Urlshort, 0)

	err := yaml.Unmarshal(yml, &arr)

	if err != nil {
		return nil, err
	}

	return arr, nil
}

func buildMap(arr []Urlshort) map[string]string {
	res := make(map[string]string)
	for _, v := range arr {
		res[v.Path] = v.Url
	}

	return res
}
