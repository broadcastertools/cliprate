package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/broadcastertools/cliprate/core/api"
	"github.com/broadcastertools/cliprate/core/fe/pkg/app"
	"github.com/broadcastertools/cliprate/core/fe/pkg/authz"
	"github.com/broadcastertools/cliprate/pkg/storage"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
)

type config struct {
	// When false, gifted subscriptions are not authorized.
	AllowGifted bool `envconfig:"allow_gifted"`
	// MongoDB DSN.
	MongoDSN string `envconfig:"mongodb_dsn", "mongodb://localhost:27017/"`
	// What to listen on for incomming requests.
	ListenTo string `envconfig:"listen", ":8000"`
	// The domain name the app is running on.
	DomainName string `envconfig:"domainname", "clips.mutexisthegoat.com"`

	// Twitch API client id.
	TwitchClientID string `envconfig:"twitch.client_id"`
	// Twitch API secret.
	TwitchSecret string `envconfig:"twitch.secret"`
	// The base url the application is running on and should accept a redirect.
	TwitchRedirectURL string `envconfig:"twitch.redirect", "http://localhost:3000"`

	// A hex colour code for the app bar.
	AppBarColour        string `envconfig:"appbarcolour", "1f1d43"`
	LogoUri             string `envconfig:"logouri", "/imgs/logo.png"`
	StreamerDisplayName string `envconfig:"steamer_display_name", "MuTeX"`
	StreamerId          string `envconfig:"streamer_id", "98506045"`
	StreamerLogin       string `envconfig:"streamer_login", "mutex"`
}

func validateConfiguration(c config) {

}

func main() {
	ctx := context.Background()

	var cnf config
	err := envconfig.Process("cliprate", &cnf)
	if err != nil {
		log.Fatal(err.Error())
	}
	validateConfiguration(cnf)

	nl, err := net.Listen("tcp", cnf.ListenTo)
	if err != nil {
		log.Fatal(err)
	}

	sc, err := storage.NewMongoStorageClient(ctx, cnf.MongoDSN)
	if err != nil {
		log.Fatal("unable to connect to MongoDB instance", err)
	}

	authorization, err := authz.NewTwitchAuth(cnf.TwitchRedirectURL, cnf.TwitchClientID, cnf.TwitchSecret)
	if err != nil {
		log.Fatal("unable to initialise Twitch authorization", err)
	}

	// Hard coded site configuration for now.
	a := app.NewApp(authorization, sc, &api.SiteConfig{
		AppbarColor:         cnf.AppBarColour,
		Domain:              cnf.DomainName,
		LogoUri:             cnf.LogoUri,
		StreamerDisplayName: cnf.StreamerDisplayName,
		StreamerId:          cnf.StreamerId,
		StreamerLogin:       cnf.StreamerLogin,
	})

	var hndlr http.Handler = api.HandlerWithOptions(a, api.ChiServerOptions{
		BaseURL: "/v1",
	})

	hndlr = cors.Default().Handler(hndlr)
	hndlr = logRequestHandler(hndlr)

	srv := http.Server{
		Handler:           hndlr,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	log.Fatal(srv.Serve(nl))
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("DEBUG: %s %s\n", r.Method, r.URL)
	}

	return http.HandlerFunc(fn)
}
