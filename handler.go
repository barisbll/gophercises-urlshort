package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func contains(m map[string]string, str string) bool {
	for k, _ := range m {
		if k == str {
			return true
		}
	}

	return false
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contains(pathsToUrls, r.URL.Path) {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusSeeOther)
		}
		fallback.ServeHTTP(w, r)
	}
}

type Yaml struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func containsYaml(v []Yaml, str string) string {
	for _, v := range v {
		if v.Path == str {
			return v.URL
		}
	}

	return ""
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var values []Yaml

	err := yaml.Unmarshal(yml, &values)
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		yamlUrl := containsYaml(values, r.URL.Path)
		if len(yamlUrl) > 0 {
			http.Redirect(w, r, yamlUrl, http.StatusSeeOther)
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
