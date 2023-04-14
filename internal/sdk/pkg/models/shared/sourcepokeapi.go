// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePokeapiPokeapiEnum string

const (
	SourcePokeapiPokeapiEnumPokeapi SourcePokeapiPokeapiEnum = "pokeapi"
)

func (e *SourcePokeapiPokeapiEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pokeapi":
		*e = SourcePokeapiPokeapiEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePokeapiPokeapiEnum: %s", s)
	}
}

// SourcePokeapi - The values required to configure the source.
type SourcePokeapi struct {
	// Pokemon requested from the API.
	PokemonName string                   `json:"pokemon_name"`
	SourceType  SourcePokeapiPokeapiEnum `json:"sourceType"`
}