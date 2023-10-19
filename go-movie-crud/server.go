package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type createMovieReq struct {
	Isbn         string `json:"isbn"`
	Title        string `json:"title"`
	DirectorName string `json:"directorname"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	r := gin.Default()

	directors := []Director{
		{FirstName: "rohit", LastName: "sethy"},
		{FirstName: "remo", LastName: "desoza"},
		{FirstName: "arbaz", LastName: "khan"},
	}
	movies := []Movie{
		{"1", "11", "chennai express", &directors[0]},
		{"2", "22", "ABCD", &directors[1]},
		{"3", "33", "Debbang", &directors[2]},
		{"4", "44", "Reddy", &directors[2]},
	}

	getMovies := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": movies,
		})
	}

	removeMovie := func(c *gin.Context) {
		id := c.Param("id")

		for index, item := range movies {
			if item.ID == id {
				movies = append(movies[:index], movies[index+1:]...)
				break
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"data": movies,
		})
	}

	getMovie := func(c *gin.Context) {
		id := c.Param("id")

		for _, item := range movies {
			if item.ID == id {
				c.JSON(http.StatusOK, gin.H{
					"data": item,
				})
				return
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "movie not found",
		})
	}

	createMovie := func(c *gin.Context) {
		var movie createMovieReq

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var d *Director

		for _, item := range directors {
			if item.FirstName == movie.DirectorName {
				d = &item
				break
			}
		}

		if d == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "director not register"})
			return
		}
		a := uuid.NewString()
		m := Movie{a, movie.Isbn, movie.Title, d}
		movies = append(movies, m)

		c.JSON(http.StatusOK, gin.H{
			"data": movies,
		})
	}

	updateMovie := func(c *gin.Context) {
		id := c.Param("id")
		var updateMovie *Movie

		for index, item := range movies {
			if item.ID == id {
				updateMovie = &movies[index]
				break
			}
		}

		if updateMovie == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "movie id doesn't exit"})
			return
		}

		var movie createMovieReq

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var d *Director

		for _, item := range directors {
			if item.FirstName == movie.DirectorName {
				d = &item
				break
			}
		}

		if d == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "director not register"})
			return
		}

		updateMovie.Director = d
		updateMovie.Title = movie.Title
		updateMovie.Isbn = movie.Isbn
		c.JSON(http.StatusOK, gin.H{
			"data": movies,
		})
	}

	r.GET("/movies", getMovies)
	r.GET("/movie/:id", getMovie)
	r.POST("/movie", createMovie)
	r.DELETE("/movie/:id", removeMovie)
	r.PUT("/movie/:id", updateMovie)
	r.Run()
}

// getMovies := func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("getMovies")
// 	if r.Method != "GET" {
// 		http.Error(w, "method not supported", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(movies)
// }

// deleteMovie := func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("deleteMovies")
// 	if r.Method != "DELETE" {
// 		http.Error(w, "method not supported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Println(r)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(movies)
// }

// http.HandleFunc("/movies", getMovies)
// http.HandleFunc("/movies/:id", deleteMovie)
