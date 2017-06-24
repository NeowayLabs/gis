# GeoJSON encoding

GeoJSON encoder/decoder based in the
[RFC7946](https://tools.ietf.org/html/rfc7946).

Features:

- Strict mode
- Parse each geometry into a proper data structure instead of a `map[string]interface{}`
- Enables clients to use geojson geometries without type assertions
