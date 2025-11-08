package html

// Pacote reultilizável

import (
	"io"
	"net/http"
	"regexp"
)

// Titulo obtem o título de uma pagina html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				c <- "Erro ao acessar " + url
				return
			}
			defer resp.Body.Close()

			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)</title>`)
			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- matches[1]
			} else {
				c <- "Sem título"
			}
		}(url)
	}
	return c
}
