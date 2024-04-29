package transform

import (
	"io/ioutil"
	"strings"

	"github.com/dattaray-basab/cks-clilib/common"
	"github.com/dattaray-basab/cks-clilib/domain"
	"github.com/dattaray-basab/cks-clilib/globals"
)

var BuildSubcommand = func(templateMap map[string]string) error {
	alterRecord, err := domain.BuildAlterInfrastucture(templateMap, QueryTemplate, TransformControlTemplate)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(alterRecord.FirstFilePath)
	if err != nil {
		return err
	}
	text := string(data)
	changedText := strings.Replace(text, alterRecord.FirstWordInFirstFile, "{{name}}", 1)
	changedData := []byte(changedText)
	err = ioutil.WriteFile(alterRecord.FirstFilePath, changedData, 0644)
	if err != nil {
		return err
	}

	quotedFullQueryId := globals.QUOTE + alterRecord.FullQueryId + globals.QUOTE

	prependString := "{%- set name = val(tokens, " + quotedFullQueryId + ") -%}" + "\n"

	err = common.PrependToFile(alterRecord.FirstFilePath, prependString)
	return err
}
