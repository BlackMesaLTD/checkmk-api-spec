package main

// ResourceDef defines which schemas belong to a resource
type ResourceDef struct {
	Name        string   // Resource name (e.g., "host")
	Description string   // Human-readable description
	Schemas     []string // Schema names to include for this resource
}

// ResourceDefinitions maps resource names to their schema definitions.
// These are discovered from CheckMK OpenAPI specs and define which
// schemas should be generated for each Terraform resource.
var ResourceDefinitions = map[string]ResourceDef{
	"host": {
		Name:        "host",
		Description: "Host configuration and attributes",
		Schemas: []string{
			// Response types
			"Host",
			"HostConfig",
			"HostConfigCollection",
			"HostExtensions",
			// Request types
			"CreateHost",
			"CreateClusterHost",
			"UpdateHost",
			"BulkCreateHost",
			"BulkUpdateHost",
			"BulkDeleteHost",
			// Attribute types
			"HostCreateAttribute",
			"HostUpdateAttribute",
			"HostViewAttribute",
		},
	},
	"folder": {
		Name:        "folder",
		Description: "Folder structure and configuration",
		Schemas: []string{
			// Response types
			"Folder",
			"FolderCollection",
			"FolderExtensions",
			"FolderMembers",
			// Request types
			"CreateFolder",
			"UpdateFolder",
			"MoveFolder",
			"BulkUpdateFolder",
			// Attribute types
			"FolderCreateAttribute",
			"FolderUpdateAttribute",
			"FolderViewAttribute",
		},
	},
	"aux_tag": {
		Name:        "aux_tag",
		Description: "Auxiliary tags for host classification",
		Schemas: []string{
			// Response types
			"AuxTag",
			"AuxTagOutput",
			"AuxTagResponse",
			"AuxTagResponseCollection",
			// Request types
			"AuxTagAttrsCreate",
			"AuxTagAttrsUpdate",
			"AuxTagAttrsResponse",
		},
	},
	"tag_group": {
		Name:        "tag_group",
		Description: "Tag groups for host classification",
		Schemas: []string{
			"TagGroup",
			"TagGroupCollection",
			"TagGroupExtensions",
			"CreateTagGroup",
			"UpdateTagGroup",
			"HostTag",
		},
	},
	"user": {
		Name:        "user",
		Description: "User accounts and permissions",
		Schemas: []string{
			"UserConfig",
			"UserConfigCollection",
			"UserExtensions",
			"CreateUser",
			"UpdateUser",
			"UserRole",
			"UserContactGroup",
		},
	},
	"contact_group": {
		Name:        "contact_group",
		Description: "Contact groups for notifications",
		Schemas: []string{
			"ContactGroup",
			"ContactGroupCollection",
			"ContactGroupExtensions",
			"CreateContactGroup",
			"UpdateContactGroup",
		},
	},
	"rule": {
		Name:        "rule",
		Description: "Monitoring rules and rulesets",
		Schemas: []string{
			"Rule",
			"RuleCollection",
			"RuleExtensions",
			"RuleProperties",
			"RuleConditions",
			"CreateRule",
			"UpdateRule",
			"MoveRule",
		},
	},
	"time_period": {
		Name:        "time_period",
		Description: "Time period definitions",
		Schemas: []string{
			"TimePeriod",
			"TimePeriodCollection",
			"TimePeriodExtensions",
			"CreateTimePeriod",
			"UpdateTimePeriod",
		},
	},
}

// GetSchemasForResources returns a deduplicated list of schemas for the given resources.
func GetSchemasForResources(resources []string) []string {
	schemaSet := make(map[string]bool)
	var schemas []string

	for _, res := range resources {
		if def, ok := ResourceDefinitions[res]; ok {
			for _, schema := range def.Schemas {
				if !schemaSet[schema] {
					schemaSet[schema] = true
					schemas = append(schemas, schema)
				}
			}
		}
	}

	return schemas
}

// GetAllResources returns all available resource names.
func GetAllResources() []string {
	resources := make([]string, 0, len(ResourceDefinitions))
	for name := range ResourceDefinitions {
		resources = append(resources, name)
	}
	return resources
}
