// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode string

const (
	SourceFaunaUpdateCollectionDeletionModeEnabledDeletionModeDeletedField SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode = "deleted_field"
)

func (e SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode) ToPointer() *SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode {
	return &e
}

func (e *SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleted_field":
		*e = SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode: %v", v)
	}
}

// SourceFaunaUpdateCollectionDeletionModeEnabled - <b>This only applies to incremental syncs.</b> <br>
// Enabling deletion mode informs your destination of deleted documents.<br>
// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
type SourceFaunaUpdateCollectionDeletionModeEnabled struct {
	// Name of the "deleted at" column.
	Column       string                                                     `json:"column"`
	DeletionMode SourceFaunaUpdateCollectionDeletionModeEnabledDeletionMode `json:"deletion_mode"`
}

type SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode string

const (
	SourceFaunaUpdateCollectionDeletionModeDisabledDeletionModeIgnore SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode = "ignore"
)

func (e SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode) ToPointer() *SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode {
	return &e
}

func (e *SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ignore":
		*e = SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode: %v", v)
	}
}

// SourceFaunaUpdateCollectionDeletionModeDisabled - <b>This only applies to incremental syncs.</b> <br>
// Enabling deletion mode informs your destination of deleted documents.<br>
// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
type SourceFaunaUpdateCollectionDeletionModeDisabled struct {
	DeletionMode SourceFaunaUpdateCollectionDeletionModeDisabledDeletionMode `json:"deletion_mode"`
}

type SourceFaunaUpdateCollectionDeletionModeType string

const (
	SourceFaunaUpdateCollectionDeletionModeTypeSourceFaunaUpdateCollectionDeletionModeDisabled SourceFaunaUpdateCollectionDeletionModeType = "source-fauna-update_Collection_Deletion Mode_Disabled"
	SourceFaunaUpdateCollectionDeletionModeTypeSourceFaunaUpdateCollectionDeletionModeEnabled  SourceFaunaUpdateCollectionDeletionModeType = "source-fauna-update_Collection_Deletion Mode_Enabled"
)

type SourceFaunaUpdateCollectionDeletionMode struct {
	SourceFaunaUpdateCollectionDeletionModeDisabled *SourceFaunaUpdateCollectionDeletionModeDisabled
	SourceFaunaUpdateCollectionDeletionModeEnabled  *SourceFaunaUpdateCollectionDeletionModeEnabled

	Type SourceFaunaUpdateCollectionDeletionModeType
}

func CreateSourceFaunaUpdateCollectionDeletionModeSourceFaunaUpdateCollectionDeletionModeDisabled(sourceFaunaUpdateCollectionDeletionModeDisabled SourceFaunaUpdateCollectionDeletionModeDisabled) SourceFaunaUpdateCollectionDeletionMode {
	typ := SourceFaunaUpdateCollectionDeletionModeTypeSourceFaunaUpdateCollectionDeletionModeDisabled

	return SourceFaunaUpdateCollectionDeletionMode{
		SourceFaunaUpdateCollectionDeletionModeDisabled: &sourceFaunaUpdateCollectionDeletionModeDisabled,
		Type: typ,
	}
}

func CreateSourceFaunaUpdateCollectionDeletionModeSourceFaunaUpdateCollectionDeletionModeEnabled(sourceFaunaUpdateCollectionDeletionModeEnabled SourceFaunaUpdateCollectionDeletionModeEnabled) SourceFaunaUpdateCollectionDeletionMode {
	typ := SourceFaunaUpdateCollectionDeletionModeTypeSourceFaunaUpdateCollectionDeletionModeEnabled

	return SourceFaunaUpdateCollectionDeletionMode{
		SourceFaunaUpdateCollectionDeletionModeEnabled: &sourceFaunaUpdateCollectionDeletionModeEnabled,
		Type: typ,
	}
}

func (u *SourceFaunaUpdateCollectionDeletionMode) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceFaunaUpdateCollectionDeletionModeDisabled := new(SourceFaunaUpdateCollectionDeletionModeDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFaunaUpdateCollectionDeletionModeDisabled); err == nil {
		u.SourceFaunaUpdateCollectionDeletionModeDisabled = sourceFaunaUpdateCollectionDeletionModeDisabled
		u.Type = SourceFaunaUpdateCollectionDeletionModeTypeSourceFaunaUpdateCollectionDeletionModeDisabled
		return nil
	}

	sourceFaunaUpdateCollectionDeletionModeEnabled := new(SourceFaunaUpdateCollectionDeletionModeEnabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceFaunaUpdateCollectionDeletionModeEnabled); err == nil {
		u.SourceFaunaUpdateCollectionDeletionModeEnabled = sourceFaunaUpdateCollectionDeletionModeEnabled
		u.Type = SourceFaunaUpdateCollectionDeletionModeTypeSourceFaunaUpdateCollectionDeletionModeEnabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceFaunaUpdateCollectionDeletionMode) MarshalJSON() ([]byte, error) {
	if u.SourceFaunaUpdateCollectionDeletionModeDisabled != nil {
		return json.Marshal(u.SourceFaunaUpdateCollectionDeletionModeDisabled)
	}

	if u.SourceFaunaUpdateCollectionDeletionModeEnabled != nil {
		return json.Marshal(u.SourceFaunaUpdateCollectionDeletionModeEnabled)
	}

	return nil, nil
}

// SourceFaunaUpdateCollection - Settings for the Fauna Collection.
type SourceFaunaUpdateCollection struct {
	// <b>This only applies to incremental syncs.</b> <br>
	// Enabling deletion mode informs your destination of deleted documents.<br>
	// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
	// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
	Deletions SourceFaunaUpdateCollectionDeletionMode `json:"deletions"`
	// The page size used when reading documents from the database. The larger the page size, the faster the connector processes documents. However, if a page is too large, the connector may fail. <br>
	// Choose your page size based on how large the documents are. <br>
	// See <a href="https://docs.fauna.com/fauna/current/learn/understanding/types#page">the docs</a>.
	PageSize int64 `json:"page_size"`
}

type SourceFaunaUpdate struct {
	// Settings for the Fauna Collection.
	Collection *SourceFaunaUpdateCollection `json:"collection,omitempty"`
	// Domain of Fauna to query. Defaults db.fauna.com. See <a href=https://docs.fauna.com/fauna/current/learn/understanding/region_groups#how-to-use-region-groups>the docs</a>.
	Domain string `json:"domain"`
	// Endpoint port.
	Port int64 `json:"port"`
	// URL scheme.
	Scheme string `json:"scheme"`
	// Fauna secret, used when authenticating with the database.
	Secret string `json:"secret"`
}
