package story

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("template.html"))
}

type HandlerOpt func(h *handler)

func WithTemplate(t *template.Template) HandlerOpt {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(f func(r *http.Request) string) HandlerOpt {
	return func(h *handler) {
		h.pathFunc = f
	}
}

func NewHandler(s Story, opts ...HandlerOpt) http.Handler {
	h := handler{s, tmp, defaultPathFunc}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

type handler struct {
	s        Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

func defaultPathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	story := h.pathFunc(r)

	if st, ok := h.s[story]; ok {
		if err := h.t.Execute(w, st); err != nil {
			log.Println(err)
			http.Error(w, "Something went wront!", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Chapter not found!", http.StatusNotFound)
	}
}

func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)

	var story Story

	err := decoder.Decode(&story)

	if err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Arc

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
