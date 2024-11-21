//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import (
	"github.com/cloudflare/cloudflare-go"
	"time"
)

type Metadata struct{}

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

type AccountDescription struct {
	ID       string
	Name     string
	Type     string
	Settings *cloudflare.AccountSettings
}

type AccountMemberDescription struct {
	UserEmail string
	ID        string
	Status    string
	AccountID string
	Code      string
	User      cloudflare.AccountMemberUserDetails
	Roles     []cloudflare.AccountRole
	Title     string
}

type AccountRoleDescription = struct {
	ID          string
	Name        string
	Description string
	Permissions map[string]cloudflare.AccountRolePermission
	AccountID   string
	Title       string
}

type ApiTokenDescription struct {
	ID         string
	Name       string
	Status     string
	Condition  *cloudflare.APITokenCondition
	ExpiresOn  *time.Time
	IssuedOn   *time.Time
	ModifiedOn *time.Time
	NotBefore  *time.Time
	Policies   []cloudflare.APITokenPolicies
}

type DNSRecordDescription struct {
	ZoneID     string
	ZoneName   string
	ID         string
	Type       string
	Name       string
	Content    string
	TTL        int
	CreatedOn  time.Time
	Locked     bool
	ModifiedOn time.Time
	Priority   *uint16
	Proxiable  bool
	Proxied    *bool
	Data       interface{}
	Meta       interface{}
}

type FireWallRuleDescription struct {
	ID          string
	Paused      bool
	Description string
	Action      string
	Title       string
	Priority    interface{}
	Filter      cloudflare.Filter
	Products    []string
	CreatedOn   time.Time
	ModifiedOn  time.Time
	ZoneID      string
}

type LoadBalancerDescription struct {
	ID                        string
	Name                      string
	ZoneName                  string
	ZoneID                    string
	TTL                       int
	Enabled                   *bool
	CreatedOn                 *time.Time
	Description               string
	FallbackPool              string
	ModifiedOn                *time.Time
	Proxied                   bool
	SessionAffinity           string
	SessionAffinityTTL        int
	SteeringPolicy            string
	DefaultPools              []string
	PopPools                  map[string][]string
	RegionPools               map[string][]string
	SessionAffinityAttributes *cloudflare.SessionAffinityAttributes
}

type LoadBalancerMonitorDescription struct {
	ID              string
	CreatedOn       *time.Time
	ModifiedOn      *time.Time
	Type            string
	Description     string
	Method          string
	Path            string
	Header          map[string][]string
	Timeout         int
	Retries         int
	Interval        int
	Port            uint16
	ExpectedBody    string
	ExpectedCodes   string
	FollowRedirects bool
	AllowInsecure   bool
	ProbeZone       string
}

type LoadBalancerPoolDescription struct {
	ID                string
	Name              string
	Enabled           bool
	Monitor           string
	CreatedOn         *time.Time
	Description       string
	Latitude          *float32
	Longitude         *float32
	MinimumOrigins    int
	ModifiedOn        *time.Time
	NotificationEmail string
	CheckRegions      []string
	LoadShedding      *cloudflare.LoadBalancerLoadShedding
	Origins           []cloudflare.LoadBalancerOrigin
}
