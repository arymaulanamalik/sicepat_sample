package shared

const (
	CollectionOpaRules    = "OpaRules"
	CollectionUserRoles   = "userRoles"
	CollectionGrants      = "grants"
	CollectionClientRoles = "clientRoles"
	CollectionBranches    = "branches"
	CollectionClients     = "clients"
	CollectionResources   = "resources"
	CollectionRoles       = "roles"
	CollectionUsers       = "users"

	// Sort
	SortASC  = "asc"
	SortDESC = "desc"

	// Redis key title
	TitleAuthorization = "authorization"
	TitleClient        = "client"
	TitleResource      = "resource"
	TitleRole          = "role"
	TitleRoleClient    = "role-client"
	TitleUser          = "user"
	TitleUserRole      = "user-role"
	TitleClientRole    = "client-role"
	TitleGrant         = "grant"
)
