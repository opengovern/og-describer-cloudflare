//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import "time"

type AccessApplicationDescription struct {
	ID                     string
	Name                   string
	AccountID              string
	AccountName            string
	Domain                 string
	CreatedAt              *time.Time
	Aud                    string
	AutoRedirectToIdentity bool
	CustomDenyMessage      string
	CustomDenyURL          string
	EnableBindingCookie    bool
	SessionDuration        string
	UpdatedAt              *time.Time
	AllowedIDPs            interface{}
	CORSHeaders            interface{}
}
