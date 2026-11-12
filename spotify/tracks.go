package spotify

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// GetAlbumTracks renvoie tous les titres d'un album, avec l'album recopié dans chaque titre.
func (c *Client) GetAlbumTracks(token string, album Album) ([]Track, error) {
	var all []Track
	limit := 50
	offset := 0

	for {
		// : même construction que pour les albums
		q := url.Values{}
		q.Set("limit", strconv.Itoa(limit))
		q.Set("offset", strconv.Itoa(offset))
		path := "/albums/" + album.ID + "/tracks?" + q.Encode()
		// on crée une variable page de type TracksPage pour stocker la réponse de l'API
		//  et on appelle la fonction get pour récupérer la page de titres
		var page TracksPage
		err := c.get(path, token, &page)
		if err != nil {
			return nil, fmt.Errorf("titres de %s : %w", album.Name, err)
		}

		//  pour chaque titre de la page
		for i := range page.Items {
			page.Items[i].Album = album
		}

		all = append(all, page.Items...)
		if page.Next == "" || len(page.Items) == 0 {
			break
		}
		offset += limit
	}
	return all, nil
}

// GetAllTracks parcourt tous les albums et rassemble leurs titres.
func (c *Client) GetAllTracks(token string, albums []Album) ([]Track, error) {
	var all []Track
	for _, album := range albums {
		tracks, err := c.GetAlbumTracks(token, album) // on récupère les titres de l'album avec la fonction précédente
		if err != nil {
			return nil, err
		}
		all = append(all, tracks...)       // on ajoute les titres à la liste complète
		time.Sleep(100 * time.Millisecond) // pour ne pas dépasser le quota de l'API je pense ...
	}
	return all, nil // on renvoie la liste complète
}
