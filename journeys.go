package meander

import "strings"

type j struct {
	Name       string
	PlaceTypes []string
}

// Journeys : 出かける目的と場所のマッピング
// 指定できる施設(https://developers.google.com/places/supported_types)
var Journeys = []interface{}{
	&j{
		Name:       "ロマンティック",
		PlaceTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"},
	},
	&j{
		Name:       "ショッピング",
		PlaceTypes: []string{"department_store", "cafe", "clothing_store", "jewely_store", "shoe_store"},
	},
	&j{
		Name:       "ナイトライフ",
		PlaceTypes: []string{"bar", "casino", "food", "bar", "night_club", "bar", "hospital"},
	},
	&j{
		Name:       "カルチャー",
		PlaceTypes: []string{"meseum", "cafe", "cemetery", "library", "art_gallery"},
	},
	&j{
		Name:       "リラックス",
		PlaceTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"},
	},
}

func (j *j) Public() interface{} {
	return map[string]interface{}{
		"name":     j.Name,
		"journeys": strings.Join(j.PlaceTypes, "|"),
	}
}
