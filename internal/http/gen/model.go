// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package gen

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewSushi defines model for NewSushi.
type NewSushi struct {
	Name  string   `json:"name"`
	Price int      `json:"price"`
	Sozai []string `json:"sozai"`
}

// Sushi defines model for Sushi.
type Sushi struct {
	// Embedded struct due to allOf(#/components/schemas/NewSushi)
	NewSushi `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Id int64 `json:"id"`
}

// FindSushisParams defines parameters for FindSushis.
type FindSushisParams struct {

	// 昇順にするかどうか
	Asc *bool `json:"asc,omitempty"`

	// 取得する件数
	Limit *int32 `json:"limit,omitempty"`
}

// AddSushiJSONBody defines parameters for AddSushi.
type AddSushiJSONBody NewSushi

// AddSushiJSONRequestBody defines body for AddSushi for application/json ContentType.
type AddSushiJSONRequestBody AddSushiJSONBody

