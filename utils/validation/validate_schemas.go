package validation

import (
	"fmt"
	"strings"

	io "github.com/TimTwigg/EncounterManagerBackend/utils/io"
	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

func ValidateGroup(asset_path string, group string, schemaFileName string, templateFileName string, hideOutput bool) {
	schemaFile := asset_path + "/schemas/" + schemaFileName
	files, err := io.ListDir(asset_path + "/" + group)
	if err != nil {
		panic(err.Error())
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(schemaFile)
	if err != nil {
		panic(err.Error())
	}

	// validate template
	templateFile := asset_path + "/templates/" + templateFileName
	template, err := io.ReadJSON(templateFile)
	if err != nil {
		panic(err.Error())
	}
	filename := strings.Split(templateFile, "/")[len(strings.Split(templateFile, "/"))-1]
	invalid := schema.Validate(template)
	if invalid != nil {
		log.Error(fmt.Sprintf("Error validating template %s: %s", filename, invalid.Error()))
	} else if !hideOutput {
		log.Info(fmt.Sprintf("Validated template %s", filename))
	}

	// validate stat block files
	for file := range files {
		instanceFile := files[file]
		instance, err := io.ReadJSON(instanceFile)
		if err != nil {
			panic(err.Error())
		}
		filename := strings.Split(instanceFile, "/")[len(strings.Split(instanceFile, "/"))-1]
		invalid := schema.Validate(instance)
		if invalid != nil {
			log.Error(fmt.Sprintf("Error validating %s %s: %s", group, filename, invalid.Error()))
		} else if !hideOutput {
			log.Info(fmt.Sprintf("Validated %s %s", group, filename))
		}
	}
}

func ValidateStatBlocks(asset_path string, hideOutput bool) {
	ValidateGroup(asset_path, "stat_blocks", "stat_block.schema.json", "stat_block.template.json", hideOutput)
}

func ValidateLanguage(asset_path string, hideOutput bool) {
	ValidateGroup(asset_path, "languages", "language.schema.json", "language.template.json", hideOutput)
}

func ValidateDamageTypes(asset_path string, hideOutput bool) {
	ValidateGroup(asset_path, "damage_types", "damage_type.schema.json", "damage_type.template.json", hideOutput)
}

func ValidateConditions(asset_path string, hideOutput bool) {
	ValidateGroup(asset_path, "conditions", "condition.schema.json", "condition.template.json", hideOutput)
}
