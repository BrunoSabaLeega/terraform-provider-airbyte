// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationScyllaUpdate struct {
	// Address to connect to.
	Address string `json:"address"`
	// Default Scylla keyspace to create data in.
	Keyspace string `json:"keyspace"`
	// Password associated with Scylla.
	Password string `json:"password"`
	// Port of Scylla.
	Port int64 `json:"port"`
	// Indicates to how many nodes the data should be replicated to.
	Replication *int64 `json:"replication,omitempty"`
	// Username to use to access Scylla.
	Username string `json:"username"`
}