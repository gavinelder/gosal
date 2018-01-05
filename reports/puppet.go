package reports

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// PuppetFacts are puppet facts duh
type PuppetFacts map[string]interface{}

// GetPuppetFacts will get puppet facts duh
func GetPuppetFacts() (PuppetFacts, error) {
	cmd := exec.Command("puppet", "facts")

	// cmd.Stderr = os.Stderr
	o, err := cmd.Output()
	if err != nil {
		return PuppetFacts{}, errors.Wrap(err, "exect puppet facts")
	}

	var pf PuppetFacts

	if err := json.Unmarshal(o, &pf); err != nil {
		return PuppetFacts{}, errors.Wrap(err, "failed unmarshalling Puppet Facts")
	}

	pf = Flatten(pf)

	fmt.Println(pf)

	return pf, nil
}

// Flatten takes a map[string]interface and flattens it
func Flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := Flatten(child)
			for nk, nv := range nm {
				o[k+"=>"+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}
