package app

import (
	"os"
	"testing"
)

func TestGenerateProject(t *testing.T) {

	t.Run("Generate Maven Project", func(t *testing.T) {
		config := GeneratorConfig{
			Metadata: ProjectMetadata{
				ApplicationName:    "maven-app",
				GroupID:            "com.sivalabs",
				ArtifactID:         "maven-app",
				ApplicationVersion: "1.0.0-SNAPSHOT",
				BasePackage:        "com.sivalabs.myapp",
				BuildTool:          "maven",
			},
			Options: Options{
				UseRDBMS:        true,
				ProdDbType:      "postgresql",
				TestDbType:      "h2",
				DBMigrationType: "flyway",
				Features:        []string{},
			},
		}
		p := ProjectGenerator{
			templatesDir: "../../templates",
			targetDir:    "../../output/" + config.Metadata.ApplicationName,
			config:       config,
		}
		defer func() {
			_ = os.RemoveAll(p.targetDir)
		}()
		if err := p.GenerateProject(); err != nil {
			t.Errorf("GenerateProject() error = %v", err)
		}
	})

	t.Run("Generate Gradle Project", func(t *testing.T) {
		config := GeneratorConfig{
			Metadata: ProjectMetadata{
				ApplicationName:    "gradle-app",
				GroupID:            "com.sivalabs",
				ArtifactID:         "gradle-app",
				ApplicationVersion: "1.0.0-SNAPSHOT",
				BasePackage:        "com.sivalabs.myapp",
				BuildTool:          "gradle",
			},
			Options: Options{
				UseRDBMS:        true,
				ProdDbType:      "mysql",
				TestDbType:      "hsqldb",
				DBMigrationType: "liquibase",
				Features:        []string{},
			},
		}
		p := ProjectGenerator{
			templatesDir: "../../templates",
			targetDir:    "../../output/" + config.Metadata.ApplicationName,
			config:       config,
		}
		defer func() {
			_ = os.RemoveAll(p.targetDir)
		}()
		if err := p.GenerateProject(); err != nil {
			t.Errorf("GenerateProject() error = %v", err)
		}
	})
}
