package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type Yaml struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func contains(m map[string]string, str string) bool {
	for k := range m {
		if k == str {
			return true
		}
	}

	return false
}

func containsYaml(v []Yaml, str string) string {
	for _, v := range v {
		if v.Path == str {
			return v.URL
		}
	}

	return ""
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contains(pathsToUrls, r.URL.Path) {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusSeeOther)
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	values := []Yaml{}

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
