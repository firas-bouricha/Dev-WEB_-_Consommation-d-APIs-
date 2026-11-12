package spotify

import (
	"encoding/json"
	"testing"
)

func TestDecodeTrack(t *testing.T) {
	data := []byte(`{
		"id": "0goer5e8B6SJaTV2B51k6f",
		"name": "La bandite",
		"duration_ms": 160893,
		"external_ids": {"isrc": "FR0U91900069"}
	}`)

	var tr Track
	err := json.Unmarshal(data, &tr)

	if err != nil {
		t.Fatalf("décodage impossible : %v", err)
	}
	if tr.Name != "La bandite" {
		t.Errorf("nom = %q", tr.Name)
	}
	if tr.DurationMs != 160893 {
		t.Errorf("durée = %d", tr.DurationMs)
	}
	if tr.ExternalIDs.ISRC != "FR0U91900069" {
		t.Errorf("isrc = %q", tr.ExternalIDs.ISRC)
	}
}
