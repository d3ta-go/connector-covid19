package types

// WHOCountryByGeometry represent WHO Country By Geometry
type WHOCountryByGeometry struct {
	Type    string  `json:"type"`
	Objects Objects `json:"objects"`
}

// Objects represent Objects
type Objects struct {
	Countries Countries `json:"countries"`
}

// Countries represent Countries
type Countries struct {
	Type       string            `json:"type"`
	Geometries []GeometryElement `json:"geometries"`
}

// GeometryElement represent Geometry Element
type GeometryElement struct {
	Arcs       interface{}        `json:"arcs"`
	Type       *Type              `json:"type"`
	Properties GeometryProperties `json:"properties"`
}

// GeometryProperties represent Geometry Properties
type GeometryProperties struct {
	ID        interface{} `json:"id"`
	WHORegion WHORegion   `json:"WHO_REGION"`
	ISO2Code  string      `json:"ISO_2_CODE"`
	Adm0Name  string      `json:"ADM0_NAME"`
	ISO3Code  string      `json:"ISO_3_CODE"`
}

// WHORegion represent WHORegion
type WHORegion string

const (
	// Afro represent AFRO WHO Regian
	Afro WHORegion = "AFRO"
	// Amro represent AMRO WHO Regian
	Amro WHORegion = "AMRO"
	// Emro represent EMRO WHO Regian
	Emro WHORegion = "EMRO"
	// Euro represent EURO WHO Regian
	Euro WHORegion = "EURO"
	// Searo represent SEARO WHO Regian
	Searo WHORegion = "SEARO"
	// Wpro represent WPRO WHO Regian
	Wpro WHORegion = "WPRO"
)

// Type represent geometry type
type Type string

const (
	// MultiPolygon represent MultiPolygon geometry type
	MultiPolygon Type = "MultiPolygon"
	// Polygon represent Polygon geometry type
	Polygon Type = "Polygon"
)
