package itemupgrader

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var (
	//go:embed schemas/*.json
	schemasFS embed.FS
	// schemas is a list of all registered item upgrade schemas.
	schemas         = make(map[int]schemaModel)
	schemasInverted = make(map[int]invertedSchemaModel)
	fileNameRegex   = regexp.MustCompile("^(\\d{4}).*\\.json$")
)

// init ...
func init() {
	files, err := schemasFS.ReadDir("schemas")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		subMatches := fileNameRegex.FindStringSubmatch(f.Name())
		if len(subMatches) == 0 {
			continue
		}
		schemaId, _ := strconv.Atoi(subMatches[1])
		file, err := schemasFS.Open("schemas/" + f.Name())
		if err != nil {
			panic(fmt.Errorf("failed to open schema: %w", err))
		}
		err = RegisterSchema(schemaId, file)
		if err != nil {
			panic(fmt.Errorf("failed to register schema: %w", err))
		}
	}
}

// RegisterSchema attempts to decode and parse a schema from the provided file reader. The file must follow the correct
// specification otherwise an error will be returned.
func RegisterSchema(schemaId int, r io.Reader) error {
	var s schemaModel
	err := json.NewDecoder(r).Decode(&s)
	if err != nil {
		return err
	}
	schemas[schemaId] = s
	schemasInverted[schemaId] = invert(s)
	return nil
}
