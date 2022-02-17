package main

import (
	"IMDB/ent"
	"IMDB/ent/director"
	"IMDB/ent/movie"
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

type M map[string]interface{}

type Page struct {
	Title string
	Body  []byte
}

/*
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
*/
func top10Handler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		top10, err := c.Movie.Query().Order(ent.Desc(movie.FieldRank)).Limit(10).All(r.Context())
		if err != nil {
			panic(err)
		}
		if err := t.Execute(w, top10); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func allHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		top10, err := c.Movie.Query().Order(ent.Asc(movie.FieldName)).All(r.Context())
		if err != nil {
			panic(err)
		}
		if err := t.Execute(w, top10); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func siteHandler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func searchHandler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func addhHandler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func moviePageHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path) // a/b/c/d => d, a/b => b, a => a
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		movie := c.Movie.GetX(r.Context(), int(idInt))
		directorOfMovie := c.Movie.GetX(r.Context(), int(idInt)).QueryDirector().OnlyX(r.Context())

		if err := t.Execute(w, M{
			"movie":           movie,
			"directorOfMovie": directorOfMovie,
		}); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func submissionHandler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}
	})
}

/*

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "Post" {
		http.Redirect(w, r, "/site", http.StatusSeeOther)
		return
	}
	mName := r.FormValue("movie")
	mDirector := r.FormValue("director")
	mRank := r.FormValue("rank")
	mDescription := r.FormValue("extra")


}

*/

func directorsHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		top10, err := c.Director.Query().Order(ent.Asc(director.FieldName)).All(r.Context())
		if err != nil {
			panic(err)
		}
		if err := t.Execute(w, top10); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func directorPageHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path) // a/b/c/d => d, a/b => b, a => a
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		director := c.Director.GetX(r.Context(), int(idInt))
		directorMovies := c.Director.GetX(r.Context(), int(idInt)).QueryMovies().WithDirector().AllX(r.Context())

		if err := t.Execute(w, M{
			"director":       director,
			"directorMovies": directorMovies,
		}); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	directors := client.Director.Query().AllX(ctx)
	movies := client.Movie.Query().AllX(ctx)
	IDS := client.Movie.Query().IDsX(ctx)

	fmt.Println("All Movies:", movies)
	fmt.Println("All Directors:", directors)
	fmt.Println("All IDs:", IDS)

	fmt.Println()
	fmt.Println()

	top10Tpl := template.Must(template.ParseFiles("frontend/top10.html"))
	siteTpl := template.Must(template.ParseFiles("frontend/site.html"))
	searchTpl := template.Must(template.ParseFiles("frontend/search.html"))
	allTpl := template.Must(template.ParseFiles("frontend/all.html"))
	addTpl := template.Must(template.ParseFiles("frontend/add.html"))
	movieTpl := template.Must(template.ParseFiles("frontend/movie-page.html"))
	submissionTpl := template.Must(template.ParseFiles("frontend/submission.html"))
	directorsTpl := template.Must(template.ParseFiles("frontend/directors.html"))
	directorPageTpl := template.Must(template.ParseFiles("frontend/director-page.html"))

	http.Handle("/top10", top10Handler(top10Tpl, client))
	http.Handle("/site", siteHandler(siteTpl))
	http.Handle("/search", searchHandler(searchTpl))
	http.Handle("/all", allHandler(allTpl, client))
	http.Handle("/add", addhHandler(addTpl))
	http.Handle("/movie/", moviePageHandler(movieTpl, client))
	http.Handle("/submission.html", submissionHandler(submissionTpl))
	http.Handle("/directors", directorsHandler(directorsTpl, client))
	http.Handle("/director/", directorPageHandler(directorPageTpl, client))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error running server (%s)", err)

	}
}
