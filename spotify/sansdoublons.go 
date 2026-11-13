package spotify
import "strings"
// SansDoublons supprime les doublons 
func SansDoublons(tracks []Track) []Track {
vu := map[string]bool{}
var uniques []Track
for _,t := range tracks {
	cle := strings.ToLower(strings.TrimSpace(t.Name))
	if vu[cle] {
		continue
	}
	vu[cle] = true
	uniques = append(uniques, t)
}
return uniques
}