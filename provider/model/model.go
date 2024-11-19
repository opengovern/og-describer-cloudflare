//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import (
	"github.com/cloudflare/cloudflare-go"
	"time"
)

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
	AllowedIDPs            []string
	CORSHeaders            *cloudflare.AccessApplicationCorsHeaders
}

type AccessGroupDescription struct {
	ID          string
	Name        string
	AccountID   string
	AccountName string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Exclude     []interface{}
	Include     []interface{}
	Require     []interface{}
}

type AccessPolicyDescription struct {
	ID                           string
	Name                         string
	ApplicationID                string
	ApplicationName              string
	AccountID                    string
	CreatedAt                    *time.Time
	Decision                     string
	Precedence                   int
	PurposeJustificationPrompt   *string
	PurposeJustificationRequired *bool
	UpdatedAt                    *time.Time
	ApprovalGroups               []cloudflare.AccessApprovalGroup
	Exclude                      []interface{}
	Include                      []interface{}
	Require                      []interface{}
}
