package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Token = le badge reçu de Spotify.
type Token struct {
	AccessToken string    `json:"access_token"` // le badge
	ExpiresIn   int       `json:"expires_in"`   // durée en secondes
	ExpiresAt   time.Time `json:"-"`            // heure d'expiration
}

// FetchToken demande un token à Spotify (Client Credentials).
func (c *Client) FetchToken() (Token, error) {
	// corps du formulaire : comme Postman Q1
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", c.ClientID)
	form.Set("client_secret", c.ClientSecret)

	// prépare le POST avec le formulaire en corps
	req, err := http.NewRequest(http.MethodPost, c.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return Token{}, err
	}

	// dit à Spotify que le corps est un formulaire
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// envoie ; erreur = réseau coupé ou timeout
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("demande de token : %w", err)
	}
	// ferme la connexion en sortant, quel que soit le chemin
	defer resp.Body.Close()

	// refus de Spotify : on ne met ni corps ni secret dans le message
	if resp.StatusCode != http.StatusOK {
		return Token{}, fmt.Errorf("token refusé : code %d", resp.StatusCode)
	}

	// range le JSON reçu dans tok
	var tok Token
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return Token{}, fmt.Errorf("décodage token : %w", err)
	}

	// heure d'expiration = maintenant + durée de vie
	tok.ExpiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
	return tok, nil
}
