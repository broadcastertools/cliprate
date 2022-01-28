package authz

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nicklaw5/helix"
)

var scopes = []string{
	// We need e-mail address to notify the subscriber.
	"user:read:email",

	// We need to be able to read subscription info to ensure they've subscribed to the stream.
	"user:read:subscriptions",
	"channel:read:subscriptions",
}

type (
	TwitchAuthorization struct {
		clientOptions helix.Options
	}
)

func NewTwitchAuth(
	baseURI string,
	twitchClientID string,
	twitchSecret string,
) (*TwitchAuthorization, error) {
	_ = context.Background()

	opts := helix.Options{
		RedirectURI:  baseURI + "/oauth/callback",
		ClientID:     twitchClientID,
		ClientSecret: twitchSecret,
	}
	_, err := helix.NewClient(&opts)
	if err != nil {
		return nil, err
	}

	return &TwitchAuthorization{
		clientOptions: opts,
	}, nil
}

// The library isn't safe to use concurrently.
func (t *TwitchAuthorization) client() *helix.Client {
	opts := t.clientOptions
	c, _ := helix.NewClient(&opts)
	return c
}

func (t *TwitchAuthorization) GetAuthorizationURL() string {
	return t.client().GetAuthorizationURL(&helix.AuthorizationURLParams{
		ResponseType: "code",
		Scopes:       scopes,
		ForceVerify:  false,
	})
}

type RequestUATResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (t *TwitchAuthorization) RequestUAT(ctx context.Context, code string) (*RequestUATResponse, error) {
	uat, err := t.client().RequestUserAccessToken(code)
	if err != nil {
		return nil, err
	}

	if uat.StatusCode > 299 {
		return nil, fmt.Errorf("error obtaining access token: %s", uat.Error)
	}

	return &RequestUATResponse{
		AccessToken:  uat.Data.AccessToken,
		RefreshToken: uat.Data.RefreshToken,
		ExpiresIn:    uat.Data.ExpiresIn,
	}, nil
}

type GetUserResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (t *TwitchAuthorization) GetUser(ctx context.Context, accessToken string) (*GetUserResponse, error) {
	c := t.client()
	c.SetUserAccessToken(accessToken)

	res, err := c.GetUsers(&helix.UsersParams{})
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("error obtaining user: %s", res.Error)
	}

	if len(res.Data.Users) != 1 {
		return nil, fmt.Errorf("error, expecting 1 user got %d", len(res.Data.Users))
	}

	return &GetUserResponse{
		ID:          res.Data.Users[0].ID,
		DisplayName: res.Data.Users[0].DisplayName,
		Email:       res.Data.Users[0].Email,
	}, nil
}

type GetSubscriptionResponse struct {
	Exists bool
	Tier   string
	Gifted bool
}

func (t *TwitchAuthorization) GetSubscription(ctx context.Context, accessToken, subscriberID, streamerId string) (*GetSubscriptionResponse, error) {
	c := t.client()
	c.SetUserAccessToken(accessToken)

	res, err := c.CheckUserSubsription(&helix.UserSubscriptionsParams{
		BroadcasterID: streamerId,
		UserID:        subscriberID,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		if res.StatusCode == http.StatusNotFound {
			return &GetSubscriptionResponse{
				Exists: false,
				Tier:   "0",
				Gifted: false,
			}, nil
		}

		return nil, fmt.Errorf("error obtaining user: %s", res.Error)
	}

	if v := len(res.Data.UserSubscriptions); v != 1 {
		return nil, fmt.Errorf("error, expecting a single subscription got %d", v)
	}

	return &GetSubscriptionResponse{
		Exists: true,
		Tier:   res.Data.UserSubscriptions[0].Tier,
		Gifted: res.Data.UserSubscriptions[0].IsGift,
	}, nil
}
