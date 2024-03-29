package validator

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/microlib/simple"
)

// checkEnvars - private function, iterates through each item and checks the required field
func checkEnvar(item string, logger *simple.Logger) error {
	name := strings.Split(item, ",")[0]
	required, _ := strconv.ParseBool(strings.Split(item, ",")[1])
	logger.Trace(fmt.Sprintf("name %s : required %t", name, required))
	if os.Getenv(name) == "" {
		if required {
			logger.Error(fmt.Sprintf("%s envar is mandatory please set it", name))
			return fmt.Errorf(fmt.Sprintf("%s envar is mandatory please set it", name))
		}

		logger.Warn(fmt.Sprintf("%s envar is empty please set it", name))
	}
	return nil
}

// ValidateEnvars : public call that groups all envar validations
// These envars are set via the openshift template
func ValidateEnvars(logger *simple.Logger) error {
	items := []string{
		"LOG_LEVEL,true",
		"WEBHOOK_SECRET,false",
		"REPO_MAPPING,false",
		"PR_OPENED_URL,false",
		"PR_MERGED_URL,false",
		"PRERELEASED_URL,false",
		"RELEASED_URL,false",
	}
	for x := range items {
		if err := checkEnvar(items[x], logger); err != nil {
			return err
		}
	}
	return nil
}
