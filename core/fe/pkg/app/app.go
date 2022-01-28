package app

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/broadcastertools/cliprate/core/api"
	"github.com/broadcastertools/cliprate/core/fe/pkg/authz"
	"github.com/broadcastertools/cliprate/pkg/storage"
)

type App struct {
	authorization *authz.TwitchAuthorization
	store         *storage.MongoStorageClient
	config        *api.SiteConfig
}

var (
	_ api.ServerInterface = (*App)(nil)
)

func NewApp(a *authz.TwitchAuthorization, s *storage.MongoStorageClient, c *api.SiteConfig) *App {
	return &App{
		authorization: a,
		store:         s,
		config:        c,
	}
}

// Marshal `d` into JSON and write to the response.
// No errors are handled.
func writeJSON(w http.ResponseWriter, d interface{}) {
	w.Header().Set("content-type", "application/json")
	rb, _ := json.Marshal(d)
	_, _ = w.Write(rb)
}

func (a *App) GetSiteConfiguration(w http.ResponseWriter, r *http.Request) {
	cnf := *a.config
	cnf.AuthorizationUrl = a.authorization.GetAuthorizationURL()

	w.WriteHeader(http.StatusOK)
	writeJSON(w, cnf)
}

func (a *App) GetClips(w http.ResponseWriter, r *http.Request) {
}

func (a *App) LoginWithTwitchCode(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	rb, _ := ioutil.ReadAll(r.Body)
	var data api.LoginWithTwitchCodeRequest

	if err := json.Unmarshal(rb, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, &api.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
		})
		return
	}

	uat, err := a.authorization.RequestUAT(ctx, data.Code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, &api.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request, code was invalid",
		})
		return
	}

	usr, err := a.authorization.GetUser(ctx, uat.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(w, &api.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal error, unable to retrieve details from Twitch",
		})
		return
	}

	println("uat: ", uat.AccessToken)

	s, err := a.authorization.GetSubscription(ctx, uat.AccessToken, usr.ID, a.config.StreamerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(w, &api.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal error, unable to retrieve subscription from Twitch",
		})
		return
	}

	if !a.config.IsGiftedAuthorized && s.Gifted {
		w.WriteHeader(http.StatusForbidden)
		writeJSON(w, &api.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Whoops, this application requires your subscription not to be gifted.",
		})
		return
	}

	self := &api.Subscriber{}

	writeJSON(w, &api.LoginWithTwitchCodeResponse{
		Token: "BOO",
		Self:  *self,
	})
}

func (a *App) GetSelf(w http.ResponseWriter, r *http.Request) {
}
