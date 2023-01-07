package modelos

// AzureUser representa um usuário SSO
type AzureUser struct {
	DisplayName       string `json:"displayName,omitempty"`
	GivenName         string `json:"givenName,omitempty"`
	Surname           string `json:"surname,omitempty"`
	JobTitle          string `json:"jobTitle,omitempty"`
	Mail              string `json:"mail,omitempty"`
	OfficeLocation    string `json:"officeLocation,omitempty"`
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
	Id                string `json:"id,omitempty"`
}
