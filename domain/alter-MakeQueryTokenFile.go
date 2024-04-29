package domain

import (
	"github.com/dattaray-basab/cks-clilib/common"
	"github.com/dattaray-basab/cks-clilib/globals"
)

var MakeQueryTokenFile = func(templateMap map[string]string, content string, queryFilePath string) error {

	queryTokenScaffold := globals.ScaffoldInfoTListT{

		{
			Filepath: queryFilePath,
			Content:  content,
		},
	}

	err := common.CreateFiles(queryTokenScaffold)
	if err != nil {
		return err
	}

	return nil
}
