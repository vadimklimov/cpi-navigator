package api

type artifactType struct {
	Name          string
	ResourceType  string
	EntitySetName string
}

type ArtifactTypes struct {
	Designtime struct {
		IntegrationFlow  artifactType
		ValueMapping     artifactType
		MessageMapping   artifactType
		ScriptCollection artifactType
	}
}

func SupportedArtifactTypes() *ArtifactTypes {
	artifactTypes := new(ArtifactTypes)

	artifactTypes.Designtime.IntegrationFlow.Name = "integration_flow"
	artifactTypes.Designtime.IntegrationFlow.ResourceType = "integrationflows"
	artifactTypes.Designtime.IntegrationFlow.EntitySetName = "IntegrationDesigntimeArtifacts"

	artifactTypes.Designtime.ValueMapping.Name = "value_mapping"
	artifactTypes.Designtime.ValueMapping.ResourceType = "valuemappings"
	artifactTypes.Designtime.ValueMapping.EntitySetName = "ValueMappingDesigntimeArtifacts"

	artifactTypes.Designtime.MessageMapping.Name = "message_mapping"
	artifactTypes.Designtime.MessageMapping.ResourceType = "messagemappings"
	artifactTypes.Designtime.MessageMapping.EntitySetName = "MessageMappingDesigntimeArtifacts"

	artifactTypes.Designtime.ScriptCollection.Name = "script_collection"
	artifactTypes.Designtime.ScriptCollection.ResourceType = "scriptcollections"
	artifactTypes.Designtime.ScriptCollection.EntitySetName = "ScriptCollectionDesigntimeArtifacts"

	return artifactTypes
}
