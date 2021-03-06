package geojson

// Point is a single set of Position.
type Point Position

// NewPoint returns a Point Feature with the specified longitude and latitude.
func NewPoint(long, lat float64) *Feature {
	return &Feature{
		Geometry: &Point{
			Longitude: long,
			Latitude:  lat,
		},
	}
}

// Type returns the geometry type.
func (*Point) Type() GeometryType {
	return PointGeometryType
}

// NewPointWithElevation returns a Point Feature with the specified longitude, latitude and elevation.
func NewPointWithElevation(long, lat, elevation float64) *Feature {
	return &Feature{
		Geometry: &Point{
			Longitude: long,
			Latitude:  lat,
			Elevation: NewOptionalFloat64(elevation),
		},
	}
}

// MarshalJSON returns the JSON encoding of the Point.
func (p *Point) MarshalJSON() ([]byte, error) {
	return (*Position)(p).MarshalJSON()
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (p *Point) UnmarshalJSON(data []byte) error {
	return (*Position)(p).UnmarshalJSON(data)
}
