package odoo

// UserContext holds a Odoo user contect object
type UserContext struct {
	Lang     string `json:"user_context"`
	TimeZone string `json:"tz"`
	UserID   int    `json:"uid"`
}

// Session houses a Odoo session object
type Session struct {
	SessionID              string      `json:"session_id"`
	UserID                 int         `json:"uid"`
	IsSystem               bool        `json:"is_system"`
	IsAdmin                bool        `json:"is_admin"`
	Context                UserContext `json:"user_context"`
	DB                     string      `json:"db"`
	ServerVersion          string      `json:"server_version"`
	Name                   string      `json:"name"`
	UserName               string      `json:"username"`
	PartnerDisplayName     string      `json:"partner_display_name"`
	CompanyID              int         `json:"company_id"`
	PartnerID              int         `json:"partner_id"`
	MaxTimeBetweenKeysInMS int         `json:"max_time_between_keys_in_ms"`
	CompanyCurrencyID      int         `json:"company_currency_id"`
	Warning                string      `json:"warning"`
	SessionString          string      `json:"sid"`
}

// Connection holds a Odoo Connection Object
type Connection struct {
	Protocol string
	Host     string
	Port     int
	DB       string
	Username string
	Password string
	Session  Session
	URL      string
}

// AuthResult is the data structure for the authorization request response
type AuthResult struct {
	Version string  `json:"jsonrpc"`
	Result  Session `json:"result"`
}

//Domain data structure to assist with building domain queries.
type Domain struct {
	Field    string
	Operator string
	Value    interface{}
}

// RequestParams are used to enforce proper request structures
type RequestParams struct {
	ID      int
	Model   string
	Method  string
	Domains []Domain
	Fields  []string
	Order   string
	Limit   int
	Offset  int
	Args    map[string]interface{}
}
