// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.1.0
// - protoc             v3.9.1
// source: lorawan-stack/api/application.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the Application message to JSON.
func (x *Application) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		s.WriteTime(x.CreatedAt)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("updated_at")
		s.WriteTime(x.UpdatedAt)
	}
	if x.DeletedAt != nil || s.HasField("deleted_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted_at")
		if x.DeletedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.DeletedAt)
		}
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if x.Description != "" || s.HasField("description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description")
		s.WriteString(x.Description)
	}
	if x.Attributes != nil || s.HasField("attributes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.Attributes {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	if x.DevEuiCounter != 0 || s.HasField("dev_eui_counter") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_eui_counter")
		s.WriteUint32(x.DevEuiCounter)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the Application message from JSON.
func (x *Application) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "created_at", "createdAt":
			s.AddField("created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.CreatedAt = *v
		case "updated_at", "updatedAt":
			s.AddField("updated_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.UpdatedAt = *v
		case "deleted_at", "deletedAt":
			s.AddField("deleted_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.DeletedAt = v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "description":
			s.AddField("description")
			x.Description = s.ReadString()
		case "attributes":
			s.AddField("attributes")
			x.Attributes = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.Attributes[key] = s.ReadString()
			})
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		case "dev_eui_counter", "devEuiCounter":
			s.AddField("dev_eui_counter")
			x.DevEuiCounter = s.ReadUint32()
		}
	})
}

// MarshalProtoJSON marshals the Applications message to JSON.
func (x *Applications) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Applications) > 0 || s.HasField("applications") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("applications")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Applications {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("applications"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the Applications message from JSON.
func (x *Applications) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "applications":
			s.AddField("applications")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Applications = append(x.Applications, nil)
					return
				}
				v := &Application{}
				v.UnmarshalProtoJSON(s.WithField("applications", false))
				if s.Err() != nil {
					return
				}
				x.Applications = append(x.Applications, v)
			})
		}
	})
}

// MarshalProtoJSON marshals the CreateApplicationRequest message to JSON.
func (x *CreateApplicationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Application != nil || s.HasField("application") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application")
		x.Application.MarshalProtoJSON(s.WithField("application"))
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.Collaborator)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the CreateApplicationRequest message from JSON.
func (x *CreateApplicationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application":
			if !s.ReadNil() {
				x.Application = &Application{}
				x.Application.UnmarshalProtoJSON(s.WithField("application", true))
			}
		case "collaborator":
			s.AddField("collaborator")
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Collaborator = v
		}
	})
}

// MarshalProtoJSON marshals the UpdateApplicationRequest message to JSON.
func (x *UpdateApplicationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Application != nil || s.HasField("application") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application")
		x.Application.MarshalProtoJSON(s.WithField("application"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the UpdateApplicationRequest message from JSON.
func (x *UpdateApplicationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application":
			if !s.ReadNil() {
				x.Application = &Application{}
				x.Application.UnmarshalProtoJSON(s.WithField("application", true))
			}
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// MarshalProtoJSON marshals the CreateApplicationAPIKeyRequest message to JSON.
func (x *CreateApplicationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ApplicationIds)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the CreateApplicationAPIKeyRequest message from JSON.
func (x *CreateApplicationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "rights":
			s.AddField("rights")
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}

// MarshalProtoJSON marshals the UpdateApplicationAPIKeyRequest message to JSON.
func (x *UpdateApplicationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ApplicationIds)
	}
	if x.ApiKey != nil || s.HasField("api_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		x.ApiKey.MarshalProtoJSON(s.WithField("api_key"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the UpdateApplicationAPIKeyRequest message from JSON.
func (x *UpdateApplicationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "api_key", "apiKey":
			if !s.ReadNil() {
				x.ApiKey = &APIKey{}
				x.ApiKey.UnmarshalProtoJSON(s.WithField("api_key", true))
			}
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// MarshalProtoJSON marshals the SetApplicationCollaboratorRequest message to JSON.
func (x *SetApplicationCollaboratorRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ApplicationIds)
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		x.Collaborator.MarshalProtoJSON(s.WithField("collaborator"))
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the SetApplicationCollaboratorRequest message from JSON.
func (x *SetApplicationCollaboratorRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "collaborator":
			if !s.ReadNil() {
				x.Collaborator = &Collaborator{}
				x.Collaborator.UnmarshalProtoJSON(s.WithField("collaborator", true))
			}
		}
	})
}