package pick

import (
	"github.com/dattaray-basab/cks-clilib/domain"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	_, err := domain.BuildAlterInfrastucture(templateMap, QueryTemplate, ControlTemplate)
	return err
}
