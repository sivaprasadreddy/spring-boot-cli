package app

type GeneratorConfig struct {
	Metadata ProjectMetadata
	Options  Options
}

type ProjectMetadata struct {
	ApplicationName    string
	GroupID            string
	ArtifactID         string
	ApplicationVersion string
	BasePackage        string
	BuildTool          string
}

type Options struct {
	UseRDBMS        bool
	ProdDbType      string
	TestDbType      string
	DBMigrationType string
	Features        []string
}
