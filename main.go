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

var OutID int

func moviePageHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path) // a/b/c/d => d, a/b => b, a => a
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err)
		}

		OutID = int(idInt)

		movie := c.Movie.GetX(r.Context(), int(idInt))
		directorOfMovie := c.Movie.GetX(r.Context(), int(idInt)).QueryDirector().OnlyX(r.Context())
		reviewsOfMovie := c.Movie.GetX(r.Context(), int(idInt)).QueryReview().AllX(r.Context())
		sumOfRanks := c.Movie.GetX(r.Context(), int(idInt)).Rank
		for _, r := range reviewsOfMovie {
			sumOfRanks += r.Rank
		}
		ranksOfMovie := sumOfRanks / (len(reviewsOfMovie) + 1)

		if err := t.Execute(w, M{
			"movie":           movie,
			"directorOfMovie": directorOfMovie,
			"reviewsOfMovie":  reviewsOfMovie,
			"rankOfMovie":     ranksOfMovie,
		}); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func submitReviewHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}

		if r.Method != "POST" {
			http.Redirect(w, r, "/site", http.StatusSeeOther)
			return
		}

		er := r.ParseForm()
		if er != nil {
			log.Fatal(er)
		}

		mReview := r.PostForm.Get("txt")
		mRank, _ := strconv.Atoi(r.PostForm.Get("rnk"))
		newReview := c.Review.Create().SetText(mReview).SetRank(mRank).SaveX(r.Context())
		newReviewToMovie := c.Movie.UpdateOneID(OutID).AddReview(newReview).SaveX(r.Context())

		fmt.Println("new review added", newReviewToMovie)
	})
}

/*
func submitRankHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}
	})
}
*/

func submitHandler(t *template.Template, c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, fmt.Sprintf("error excuting template (%s)", err), http.StatusInternalServerError)
		}
		if r.Method != "POST" {
			http.Redirect(w, r, "/site", http.StatusSeeOther)
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		mDescription := r.PostForm.Get("extra")
		mName := r.PostForm.Get("movie")
		mDirector := r.PostForm.Get("director")
		mRank, _ := strconv.Atoi(r.PostForm.Get("ranking"))
		mReview := r.PostForm.Get("rev")

		newMovie := c.Movie.Create().SetName(mName).SetDescription(mDescription).SetRank(mRank).SaveX(r.Context())
		newDirector := c.Director.Create().SetName(mDirector).SaveX(r.Context())
		newMovieID := c.Movie.Query().Where(movie.Name(newMovie.Name)).OnlyIDX(r.Context())
		newMovieToDirector := c.Director.UpdateOne(newDirector).AddMovieIDs(newMovieID).SaveX(r.Context())
		newReview := c.Review.Create().SetText(mReview).SetRank(mRank).SaveX(r.Context())
		newReviewToMovie := c.Movie.UpdateOne(newMovie).AddReview(newReview).SaveX(r.Context())

		fmt.Println("new movie added:", newMovie, "new Director added:", newDirector)
		fmt.Println("new conecctions made:", newMovieToDirector, newReviewToMovie)
		fmt.Println("new review added:", newReview)

	})
}

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

func cssHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("frontend/css/"))
	tmpl.Execute(w, tmpl)
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
	submitTpl := template.Must(template.ParseFiles("frontend/submission.html"))
	submitReviewTpl := template.Must(template.ParseFiles("frontend/submissionRev.html"))
	directorsTpl := template.Must(template.ParseFiles("frontend/directors.html"))
	directorPageTpl := template.Must(template.ParseFiles("frontend/director-page.html"))
	//submitRankTpl := template.Must(template.ParseFiles("frontend/submissionRank.html"))

	http.Handle("/top10", top10Handler(top10Tpl, client))
	http.Handle("/site", siteHandler(siteTpl))
	http.Handle("/search", searchHandler(searchTpl))
	http.Handle("/all", allHandler(allTpl, client))
	http.Handle("/add", addhHandler(addTpl))
	http.Handle("/movie/", moviePageHandler(movieTpl, client))
	http.Handle("/submission.html", submitHandler(submitTpl, client))
	http.Handle("/submissionRev.html", submitReviewHandler(submitReviewTpl, client))
	http.Handle("/directors", directorsHandler(directorsTpl, client))
	http.Handle("/director/", directorPageHandler(directorPageTpl, client))
	//http.Handle("/submissionRank.html", submitRankHandler(submitRankTpl, client))

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/frontend/css/", http.StripPrefix("/frontend/css/", fs))
	http.HandleFunc("/css", cssHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error running server (%s)", err)

	}
}
