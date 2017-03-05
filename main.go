package main

import (
	"github.com/google/go-github/github"
	"github.com/highlanderkev/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/oauth2"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
	"os"
)

var (
	personalAccessToken = os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
)

type TokenSource struct {
	AccessToken string
}

type Server struct {
	ip   string
	port string
}

type Clients struct {
	ghClient *github.Client
	fbClient *firego.Firebase
}

func setupRouter(clients Clients) *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	router.GET("/", handlers.GetStatus())
	router.GET("/status", handlers.GetStatus())
	router.GET("/version", handlers.GetVersion())
	router.GET("/repos", handlers.GetRepos(clients.ghClient))
	router.GET("/posts", handlers.GetPosts(clients.fbClient))
	router.GET("/gists", handlers.GetGists(clients.ghClient))
	return router
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func setupClients() Clients {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return Clients{
		ghClient: github.NewClient(oauthClient),
		fbClient: firego.New("https://highlanderkev-github-io.firebaseio.com/", nil),
	}
}

func main() {
	server := Server{
		ip:   os.Getenv("IP"),
		port: os.Getenv("PORT"),
	}

	if server.port == "" && server.ip == "" {
		log.Fatal("$PORT must be set")
	}

	clients := setupClients()
	router := setupRouter(clients)
	router.Start(server.ip + ":" + server.port)
}
