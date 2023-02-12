package story

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

const defaultTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Choose Your Own Adventure</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{range .Story}}
  <p>{{.}}</p>
  {{end}}
  <ul>
    {{range .Options}}
    <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>
</html>
`

var tmp *template.Template

func init() {
	tmp = template.Must(template.New("").Parse(defaultTemplate))
}

func NewHandler(story Story) http.Handler {
	return handler{story}
}

type handler struct {
	Story Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp.Execute(w, h.Story["intro"])
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
