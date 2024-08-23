package twitch

import (
	"fmt"
	"os"

	"github.com/nicklaw5/helix/v2"
)

type client struct {
	*helix.Client
	token *Token
}

func (c *client) UserID() string {
	return c.token.UserId
}

func newClient(token *Token) (*client, error) {
	apiURL := ""
	if v := os.Getenv("TWITCH_API_URL"); v != "" {
		apiURL = v
	}
	hClient, err := helix.NewClient(&helix.Options{
		APIBaseURL:      apiURL,
		ClientID:        token.ClientId,
		UserAccessToken: token.Access,
	})
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}
	client := &client{
		Client: hClient,
		token:  token,
	}

	return client, nil
}

func (t *Twitch) addProfile(profile *Profile) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	profile.Token.ClientId = clientID
	client, err := newClient(profile.Token)
	if err != nil {
		return fmt.Errorf("creating client: %w", err)
	}
	t.clients[profile.Name] = client
	return nil
}
