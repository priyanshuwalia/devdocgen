package detector

import (
	"encoding/json"
	"os"

	"devdocgen/internal/model"
)

type packageJSON struct {
    Scripts map[string]string `json:"scripts"`
}

func DetectNodeScripts(root string, meta *model.ProjectMetadata) {
    file, err := os.Open(root + "/package.json")
    if err != nil {
        return
    }
    defer file.Close()

    var pkg packageJSON
    if err := json.NewDecoder(file).Decode(&pkg); err != nil {
        return
    }

    if _, ok := pkg.Scripts["start"]; ok {
        meta.RunCommands = append(meta.RunCommands, "npm start")
    } else if _, ok := pkg.Scripts["dev"]; ok {
        meta.RunCommands = append(meta.RunCommands, "npm run dev")
    }
}
