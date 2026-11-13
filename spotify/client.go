package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Client regroupe tout ce qu'il faut pour parler à Spotify.
type Client struct {
	HTTP         *http.Client // le coursier (avec son chrono)
	BaseURL      string       // adresse de base de l'API
	ClientID     string       // identifiant de l'application
	ClientSecret string       // secret de l'application
	TokenURL     string       // adresse où demander le token
}

// NewClient construit un client avec un timeout.
func NewClient(timeout time.Duration) *Client {
	return &Client{
		HTTP:     &http.Client{Timeout: timeout},
		BaseURL:  "https://api.spotify.com/v1",
		TokenURL: "https://accounts.spotify.com/api/token",
	}
}

// get envoie GET baseURL+path avec le token, et décode la réponse JSON dans out.
// En cas de 429 (trop de requêtes), il attend le délai Retry-After puis réessaie (3 essais max).
func (c *Client) get(path, token string, out any) error {
	// prépare la requête ; erreur si l'URL est mauvaise
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return fmt.Errorf("mauvais url : %w", err)
	}
	// badge d'accès, comme Bearer Token dans Postman
	req.Header.Set("Authorization", "Bearer "+token)

	for essai := 1; essai <= 3; essai++ {
		// envoie la requête ; erreur = réseau coupé ou timeout
		resp, err := c.HTTP.Do(req)
		if err != nil {
			return fmt.Errorf("appel spotify : %w", err)
		}

		// 429 : trop de requêtes, Spotify demande d'attendre
		if resp.StatusCode == http.StatusTooManyRequests {
			// Retry-After = nombre de secondes à attendre (reçu sous forme de texte)
			delai, err := strconv.Atoi(resp.Header.Get("Retry-After"))
			if err != nil {
				resp.Body.Close()
				return fmt.Errorf("quota Spotify dépassé, impossible de lire Retry-After : %w", err)
			}
			// attente trop longue : on s'arrête avec un message clair
			if delai > 30 {
				resp.Body.Close()
				return fmt.Errorf("quota Spotify dépassé, réessayer dans %d s", delai)
			}
			// ferme cette réponse, attend, puis refait un tour de boucle
			resp.Body.Close()
			time.Sleep(time.Duration(delai) * time.Second)
			continue
		}

		// tout code autre que 200 = erreur ; on lit le message de Spotify
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			return fmt.Errorf("spotify a répondu %d : %s", resp.StatusCode, body)
		}

		// range le JSON reçu dans out (le struct de l'appelant)
		err = json.NewDecoder(resp.Body).Decode(out)
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("réponse illisible : %w", err)
		}
		// tout s'est bien passé
		return nil
	}

	// 3 essais, toujours des 429
	return fmt.Errorf("spotify : trop de tentatives (429)")
}
