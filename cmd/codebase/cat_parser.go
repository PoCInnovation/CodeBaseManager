package codebase

func CatParseRepo(catStruct parsingRepo) {
	// TODO: change type repo by parsed type
	// TODO: Add opendir management.

	for _, module := range catStruct.repo {

		if FoundAllArgs(catStruct.found) {
			return
		}
	}

	//for _, filePath := range repo {
	//	fileContent, err := codebase.GetFile(filePath)
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Println(*fileContent)
	//	//findTargetFromArgs(filePath, *fileContent, parser)
	//	if foundAllArgs(parser.found) {
	//		return
	//	}
	//}
	//return print content
}
