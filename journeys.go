package meander

type j struct {
	Name       string
	PlaceTypes []string
}

var Journeys = []interface{}{
	&j{Name: "Romantic", PlaceTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name: "Shopping", PlaceTypes: []string{"department_store", "cafe", "clothing_store", "jewely_store", "shoe_store"}},
	&j{Name: "NightLife", PlaceTypes: []string{"bar", "casino", "food", "bar", "naight_club", "bar", "bar", "hospital"}},
	&j{Name: "Calture", PlaceTypes: []string{"museum", "cafe", "cemetery", "library", "art_gallery"}},
	&j{Name: "Relax", PlaceTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}
