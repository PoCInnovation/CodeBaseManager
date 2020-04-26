package codebase

type findFctArray struct {
	language string
	fct      targetFctParser
}

type targetFctParser func(fileContent, toFind string) (content *string, err error)
type contentFound map[string]string

type parsingRepo struct {
	content contentFound
	args    []string
	found   []bool
	repo    []string
}

func findTargetFromArgs(filepath, fileContent string, parser *parsingRepo) (err error) {
	for idx, toFind := range parser.args {
		if filepath == toFind {
			parser.content[toFind] = fileContent
			parser.found[idx] = true
			continue
		}
		if content, err := parser.target(fileContent, toFind); err == nil {
			parser.content[toFind] = *content
			parser.found[idx] = true
		} else {
			return err
		}
	}
	return nil
}
