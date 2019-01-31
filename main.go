package main

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
	"strconv"
)

func quadkeyString(t maptile.Tile) string {
	s := strconv.FormatInt(int64(t.Quadkey()),4)
	// for zero padding
	zeros := "000000000000000000000000000000"
	return zeros[:((int(t.Z)+1)-len(s))/2] + s
}

func main() {
	name := "皇居"
	address := "東京都千代田区千代田１−１"

	// Creating a point
	lat := 35.685323
	lng := 139.752768
	p := orb.Point{lng, lat}

	// Change to a maptile.
	t := maptile.At(p, 17)
	println(quadkeyString(t)) // 13300211231022032

	// viewing as geojson
	c := geojson.NewFeatureCollection()
	gp := geojson.NewFeature(p)
	gp.Properties["name"] = name
	gp.Properties["address"] = address

	gt := geojson.NewFeature(t.Bound())
	gt.Properties["quadkey"] = quadkeyString(t)

	c = c.Append(gp).Append(gt)
	b, _ := c.MarshalJSON()
	println(string(b)) //{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[139.752768,35.685323]},"properties":{"address":"東京都千代田区千代田１−１","name":"皇居"}},{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[139.7515869140625,35.684071533140965],[139.75433349609375,35.684071533140965],[139.75433349609375,35.68630240145626],[139.7515869140625,35.68630240145626],[139.7515869140625,35.684071533140965]]]},"properties":{"quadkey":"13300211231022032"}}]}
}