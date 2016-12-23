package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// htmlとして解釈
	// 自動的な解釈としては、デフォルトはtext/plain,
	// 出力の先頭にタグがある場合そのままtext/htmlとなる
	w.Header().Set("Content-Type", "text/html")

	Parse(w, r)
	Panic(w)
	Example(w)
	ExampleEscape(w)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

func Parse(w http.ResponseWriter, r *http.Request) {

	t := template.New("test")
	t = template.Must(t.Parse("{{define \"base\"}}test : {{.foo}}, {{.bar}}<br>\n{{end}}"))

	fmt.Println("name : " + t.Name())
	dat := map[string]string{"foo": "bar"}

	err := t.ExecuteTemplate(w, "base", dat)
	if err != nil {
		fmt.Println(err)
	}

	// html/template: "base3" is undefined
	err = t.ExecuteTemplate(w, "base3", dat)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, template.HTML("<b>HTML</b>"))

	if t.Lookup("lookup") == nil {
		fmt.Println("no template lookup")
	}

	if t.Lookup("test") == nil {
		fmt.Println("no template test")
	}
}

func Panic(w http.ResponseWriter) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprint(w, "Error::", err, template.HTML(`<br>`))
		}
	}()
	t := template.Must(template.New("").Parse("{{"))
	fmt.Println(t)
}

func Example(w http.ResponseWriter) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>
`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(w, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(w, noItems)
	check(err)

}

func ExampleEscape(w http.ResponseWriter) {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	v := []interface{}{`"Fran & Freddie's Diner"`, ' ', `<tasty@example.com>`}

	fmt.Println(template.HTMLEscapeString(s))
	template.HTMLEscape(w, []byte(s))
	fmt.Fprintln(w, "")
	fmt.Println(template.HTMLEscaper(v...))

	fmt.Println(template.JSEscapeString(s))
	template.JSEscape(w, []byte(s))
	fmt.Fprintln(w, "")
	fmt.Println(template.JSEscaper(v...))

	fmt.Println(template.URLQueryEscaper(v...))
}
