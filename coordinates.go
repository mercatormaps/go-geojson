package geojson

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Coordinates represents a longitude and latitude with optional elevation/altitude.
type Coordinates struct {
	Longitude float64
	Latitude  float64
	Elevation OptionalFloat64
}

// MarshalJSON returns the JSON encoding of the Coordinates.
// The JSON encoding is an array of numbers with the longitude followed by the latitude, and optional elevation.
func (c *Coordinates) MarshalJSON() ([]byte, error) {
	if c.Elevation.IsSet() {
		return json.Marshal(&coordinates{
			c.Longitude,
			c.Latitude,
			c.Elevation.Value(),
		})
	}

	return json.Marshal(&coordinates{
		c.Longitude,
		c.Latitude,
	})
}

// UnmarshalJSON parses the JSON-encoded data and stores the results.
func (c *Coordinates) UnmarshalJSON(data []byte) error {
	coords := coordinates{}
	if err := json.Unmarshal(data, &coords); err != nil {
		return err
	}

	if len(coords) < 2 {
		return errors.New("incomplete coordinates")
	}

	c.Longitude = coords[0]
	c.Latitude = coords[1]
	if len(coords) > 2 {
		c.Elevation = NewOptionalFloat64(coords[2])
	}
	return nil
}

// OptionalFloat64 is a type that represents a float64 that can be optionally set.
type OptionalFloat64 struct {
	value *float64
}

// NewOptionalFloat64 creates a new OptionalFloat64 set to the specified value.
func NewOptionalFloat64(val float64) OptionalFloat64 {
	return OptionalFloat64{value: &val}
}

// Value returns the value. Should call this method if OptionalFloat64.IsSet() returns true.
func (o OptionalFloat64) Value() float64 {
	return *o.value
}

// IsSet returns true if the value is set, and false if not.
func (o OptionalFloat64) IsSet() bool {
	return o.value != nil
}

// Get the float64 value and whether or not it's set.
func (o OptionalFloat64) Get() (float64, bool) {
	if o.value == nil {
		return 0, false
	}
	return *o.value, true
}

func (o OptionalFloat64) String() string {
	if o.IsSet() {
		return strconv.FormatFloat(o.Value(), 'f', -1, 64)
	}
	return "{unset}"
}

type coordinates []float64
