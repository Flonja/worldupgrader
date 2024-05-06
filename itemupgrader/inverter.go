package itemupgrader

// invertedSchemaModel represents the schema (but it's inverted) for loading item upgrade data from a JSON file.
type invertedSchemaModel struct {
	RenamedIDs    map[string]string             `json:"renamedIds,omitempty"`
	RemappedMetas map[string]remappedMetaResult `json:"remappedMetas,omitempty"`
}

type remappedMetaResult struct {
	Identifier    string
	MetadataValue int16
}

func invert(model schemaModel) (schema invertedSchemaModel) {
	schema.RenamedIDs = make(map[string]string)
	schema.RemappedMetas = make(map[string]remappedMetaResult)
	for oldId, newId := range model.RenamedIDs {
		schema.RenamedIDs[newId] = oldId
	}
	for oldId, newId := range model.RemappedMetas {
		for metadata, identifier := range newId {
			schema.RemappedMetas[identifier] = remappedMetaResult{
				Identifier:    oldId,
				MetadataValue: metadata,
			}
		}
	}
	return schema
}
