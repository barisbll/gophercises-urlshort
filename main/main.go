package main

import (
	"flag"
	"fmt"
	"net/http"

	urlshort "github.com/barisbll/gophercises-urlshort"
)

var pathsToUrls = map[string]string{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

var defaultYaml = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

// TODO: create a serveYaml function that reads yaml files and serve them
// if no port and yaml file provided then use the default yaml
func main() {
	yamlFileFlag := flag.String("yaml", "", "Location of the yaml file passed")
	portFlag := flag.Int("port", 8080, "Port number that is serving the web server")
	flag.Parse()
	fmt.Printf("yamlFileFlag %s\n", *yamlFileFlag)

	mux := defaultMux()

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler([]byte(defaultYaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting the server on :%d", *portFlag)
	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
