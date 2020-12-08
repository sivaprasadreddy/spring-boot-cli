package app

type GeneratorConfig struct {
	Metadata ProjectMetadata
}

type ProjectMetadata struct {
	ApplicationName    string
	GroupID            string
	ArtifactID         string
	ApplicationVersion string
	BasePackage        string
}
