// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var ApplicationFieldPathsNested = []string{
	"attributes",
	"contact_info",
	"created_at",
	"deleted_at",
	"description",
	"ids",
	"ids.application_id",
	"name",
	"updated_at",
}

var ApplicationFieldPathsTopLevel = []string{
	"attributes",
	"contact_info",
	"created_at",
	"deleted_at",
	"description",
	"ids",
	"name",
	"updated_at",
}
var ApplicationsFieldPathsNested = []string{
	"applications",
}

var ApplicationsFieldPathsTopLevel = []string{
	"applications",
}
var GetApplicationRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"field_mask",
}

var GetApplicationRequestFieldPathsTopLevel = []string{
	"application_ids",
	"field_mask",
}
var ListApplicationsRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"deleted",
	"field_mask",
	"limit",
	"order",
	"page",
}

var ListApplicationsRequestFieldPathsTopLevel = []string{
	"collaborator",
	"deleted",
	"field_mask",
	"limit",
	"order",
	"page",
}
var CreateApplicationRequestFieldPathsNested = []string{
	"application",
	"application.attributes",
	"application.contact_info",
	"application.created_at",
	"application.deleted_at",
	"application.description",
	"application.ids",
	"application.ids.application_id",
	"application.name",
	"application.updated_at",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
}

var CreateApplicationRequestFieldPathsTopLevel = []string{
	"application",
	"collaborator",
}
var UpdateApplicationRequestFieldPathsNested = []string{
	"application",
	"application.attributes",
	"application.contact_info",
	"application.created_at",
	"application.deleted_at",
	"application.description",
	"application.ids",
	"application.ids.application_id",
	"application.name",
	"application.updated_at",
	"field_mask",
}

var UpdateApplicationRequestFieldPathsTopLevel = []string{
	"application",
	"field_mask",
}
var ListApplicationAPIKeysRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"limit",
	"page",
}

var ListApplicationAPIKeysRequestFieldPathsTopLevel = []string{
	"application_ids",
	"limit",
	"page",
}
var GetApplicationAPIKeyRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"key_id",
}

var GetApplicationAPIKeyRequestFieldPathsTopLevel = []string{
	"application_ids",
	"key_id",
}
var CreateApplicationAPIKeyRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"expires_at",
	"name",
	"rights",
}

var CreateApplicationAPIKeyRequestFieldPathsTopLevel = []string{
	"application_ids",
	"expires_at",
	"name",
	"rights",
}
var UpdateApplicationAPIKeyRequestFieldPathsNested = []string{
	"api_key",
	"api_key.created_at",
	"api_key.expires_at",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"api_key.updated_at",
	"application_ids",
	"application_ids.application_id",
	"field_mask",
}

var UpdateApplicationAPIKeyRequestFieldPathsTopLevel = []string{
	"api_key",
	"application_ids",
	"field_mask",
}
var ListApplicationCollaboratorsRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"limit",
	"page",
}

var ListApplicationCollaboratorsRequestFieldPathsTopLevel = []string{
	"application_ids",
	"limit",
	"page",
}
var GetApplicationCollaboratorRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
}

var GetApplicationCollaboratorRequestFieldPathsTopLevel = []string{
	"application_ids",
	"collaborator",
}
var SetApplicationCollaboratorRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.ids",
	"collaborator.ids.ids.organization_ids",
	"collaborator.ids.ids.organization_ids.organization_id",
	"collaborator.ids.ids.user_ids",
	"collaborator.ids.ids.user_ids.email",
	"collaborator.ids.ids.user_ids.user_id",
	"collaborator.rights",
}

var SetApplicationCollaboratorRequestFieldPathsTopLevel = []string{
	"application_ids",
	"collaborator",
}
