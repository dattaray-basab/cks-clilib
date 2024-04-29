package domain

import (
	"github.com/dattaray-basab/cks-clilib/common"
	"github.com/dattaray-basab/cks-clilib/globals"
)

var MakeControlFile = func(templateMap map[string]string, content string) error {

	jsonControlFileScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: templateMap[globals.KEY_CONTROL_JSON_PATH],
			Content:  content,
		},
	}

	err := common.CreateFiles(jsonControlFileScaffold)
	if err != nil {
		return err
	}
	return nil
}
