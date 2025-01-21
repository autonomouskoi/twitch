package twitch

import (
	"fmt"
	"os"

	"github.com/nicklaw5/helix/v2"
)

// client wraps a helix client and maintains its auth token
type client struct {
	*helix.Client
	token *Token
}

// UserID returns the user ID associated with this client, from the token
func (c *client) UserID() string {
	return c.token.UserId
}

// create a new client based on an active Token
func newClient(token *Token) (*client, error) {
	apiURL := ""
	if v := os.Getenv("TWITCH_API_URL"); v != "" {
		apiURL = v // allow overriding the API URL to use the twitch dev tool
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

// create a client from profile and save it internally
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
