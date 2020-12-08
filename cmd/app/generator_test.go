package app

import (
	"os"
	"testing"
)

func TestGenerateProject(t *testing.T) {

	t.Run("Generate Project", func(t *testing.T) {
		config := GeneratorConfig{
			Metadata: ProjectMetadata{
				ApplicationName:    "demo-app1",
				GroupID:            "com.sivalabs",
				ArtifactID:         "demo-app1",
				ApplicationVersion: "1.0.0-SNAPSHOT",
				BasePackage:        "com.sivalabs.myapp",
			},
		}
		p := ProjectGenerator{
			templatesDir: "../../templates",
			targetDir:    "../../" + config.Metadata.ApplicationName,
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
