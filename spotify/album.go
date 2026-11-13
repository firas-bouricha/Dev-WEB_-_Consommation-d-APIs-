package spotify

import (
	"fmt"
	"net/url"
	"strconv"
)

// accessible depuis main c pour J en majuscules, sinon c'est privé
const JulID = "3IW7ScrzXmPvZhB27hmfgy"

// GetArtistAlbums renvoie TOUS les albums et singles d'un artiste .
func (c *Client) GetArtistAlbums(token, artistID string) ([]Album, error) {
	var all []Album // la liste complète, vide au départ
	limit := 10
	offset := 0

	for {
		q := url.Values{}                                        // prépare les paramètres de la requête
		q.Set("include_groups", "album,single")                  // on veut les albums et les singles
		q.Set("limit", strconv.Itoa(limit))                      // 10 par page
		q.Set("offset", strconv.Itoa(offset))                    //  on commence à la page 0, puis 10, puis 20...
		path := "/artists/" + artistID + "/albums?" + q.Encode() // c le lien complet de l'appel à Spotify

		var page AlbumsPage              // la page reçue de Spotify
		err := c.get(path, token, &page) // on appelle la fonction get pour récupérer la page d'albums
		if err != nil {
			return nil, fmt.Errorf("albums offset %d : %w", offset, err)
		}

		all = append(all, page.Items...) //  ajouter les albums de la page à la liste complète un par un

		if page.Next == "" || len(page.Items) == 0 {
			break
		}

		offset = offset + limit //  on passe à la page suivante
	}

	return all, nil
}
