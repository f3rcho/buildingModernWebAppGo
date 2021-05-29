# Building a modern web app with Golang

## Personal notes from my Go learning path

To start running the server just clone the repo and run

```bash
$ go run main.go
```

## This section pretends to show some snippes of code. Maybe a how to do it!!

### Rendering content (templates)

we could create a function called renderTemplate a call the built-in method template, parsing the files:

```go
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error rendering template", err)
		return
	}
}
```

And the call this function in our routes handlers

```go
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}
```
