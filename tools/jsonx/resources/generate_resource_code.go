package resources

import (
	"encoding/json"
	"fmt"
	"strings"
)

type resource struct {
	Code     string `json:"code"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func buildResourceCodes(resources []resource) {
	b, err := json.Marshal(resources)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("json is \n", string(b))
}

func NewResources() {

	var category = "battle"
	var codePrefix = "action." + strings.Title(category) + "."

	var names = []string{
		"ListAllTrainingBattles",
		"StopTrainingBattle",
	}
	resources := []resource{}
	for _, name := range names {
		resources = append(resources, resource{
			Code:     codePrefix + name,
			Type:     "action",
			Name:     name,
			Category: category,
		})
	}

	buildResourceCodes(resources)
}
