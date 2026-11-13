package spotify

import (
	"fmt"
	"strings"
)

// Duree renvoie la durée au format mm:ss
func (t Track) Duree() string {
	total := t.DurationMs / 1000
	minutes := total / 60
	secondes := total % 60
	return fmt.Sprintf("%02d:%02d", minutes, secondes)
}

// Annee renvoie l'année de sortie (les 4 premiers caractères de "2019-06-28").
func (t Track) Annee() string {
	d := t.Album.ReleaseDate
	if len(d) < 4 {
		return ""
	} else {
		return d[:4] //return les 4 premiers caractères de releaseDate
	}
}

// AFeat vaut true s'il y a au moins un artiste autre que JUL.
func (t Track) AFeat() bool {
	for _, a := range t.Artists {
		if strings.ToLower(a.Name) != "jul" {
			return true
		}
	}
	return false
}
