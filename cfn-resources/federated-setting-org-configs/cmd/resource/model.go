// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	ApiKeys                  *ApiKeyDefinition   `json:",omitempty"`
	DomainAllowList          []string            `json:",omitempty"`
	DomainRestrictionEnabled *bool               `json:",omitempty"`
	TestMode                 *string             `json:",omitempty"`
	FederationSettingsId     *string             `json:",omitempty"`
	IdentityProviderId       *string             `json:",omitempty"`
	OrgId                    *string             `json:",omitempty"`
	PostAuthRoleGrants       []string            `json:",omitempty"`
	RoleMappings             []RoleMappingView   `json:",omitempty"`
	UserConflicts            []FederatedUserView `json:",omitempty"`
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PrivateKey *string `json:",omitempty"`
	PublicKey  *string `json:",omitempty"`
}

// RoleMappingView is autogenerated from the json schema
type RoleMappingView struct {
	ExternalGroupName *string          `json:",omitempty"`
	Id                *string          `json:",omitempty"`
	RoleAssignments   []RoleAssignment `json:",omitempty"`
}

// RoleAssignment is autogenerated from the json schema
type RoleAssignment struct {
	GroupId *string `json:",omitempty"`
	OrgId   *string `json:",omitempty"`
	Role    *string `json:",omitempty"`
}

// FederatedUserView is autogenerated from the json schema
type FederatedUserView struct {
	ApiKeys              *ApiKeyDefinition `json:",omitempty"`
	EmailAddress         *string           `json:",omitempty"`
	FederationSettingsId *string           `json:",omitempty"`
	FirstName            *string           `json:",omitempty"`
	LastName             *string           `json:",omitempty"`
	UserId               *string           `json:",omitempty"`
}
