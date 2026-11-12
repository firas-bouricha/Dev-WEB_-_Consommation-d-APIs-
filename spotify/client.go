package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client regroupe tout ce qu'il faut pour parler à Spotify.
type Client struct {
	HTTP         *http.Client
	BaseURL      string
	ClientID     string
	ClientSecret string
}

// NewClient construit un client avec un timeout.
func NewClient(timeout time.Duration) *Client {
	return &Client{
		HTTP:    &http.Client{Timeout: timeout},
		BaseURL: "https://api.spotify.com/v1",
	}
}

// get envoie GET baseURL+path avec le token, et décode la réponse JSON dans out.
func (c *Client) get(path, token string, out any) error {
	// get baseURL+path et le stocké dans req si error on return  err
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return fmt.Errorf("mauvais url : %w", err)
	}
	//C’est ce que j'ai  fait dans Postman : l’onglet Authorization → Bearer Token
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return fmt.Errorf("appel spotify : %w", err)
	}
	//quand la fonction est fini fermé la connexion
	defer resp.Body.Close()
	// tout code autre que 200 = erreur ; on lit le message de Spotify
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("spotify a répondu %d : %s", resp.StatusCode, body)
	}
	// range le JSON reçu dans out (le struct de l'appelant)
	err = json.NewDecoder(resp.Body).Decode(out)
	if err != nil {
		return fmt.Errorf("réponse illisible : %w", err)
	}
	// tout s'est bien passé
	return nil
}
