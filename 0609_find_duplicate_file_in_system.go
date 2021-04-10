package main

import "strings"

func findDuplicate(paths []string) [][]string {
	duplicates := make(map[string][]string)

	getDir := func(path string) string {
		return strings.Split(path, " ")[0]
	}

	getFiles := func(path string) ([]string, []string) {
		filesWithContent := strings.Split(path, " ")[1:]
		files := make([]string, 0, len(filesWithContent))
		content := make([]string, 0, len(filesWithContent))

		for _, f := range filesWithContent {
			files = append(files, strings.Split(f, "(")[0])
			content = append(content, strings.Split(f, "(")[1])
		}

		return files, content
	}

	for _, p := range paths {
		directory := getDir(p)
		files, content := getFiles(p)

		for i := range files {
			if _, ok := duplicates[content[i]]; !ok {
				duplicates[content[i]] = make([]string, 0)
			}
			duplicates[content[i]] = append(duplicates[content[i]], directory+"/"+files[i])
		}
	}

	result := make([][]string, 0)
	for _, d := range duplicates {
		if len(d) > 1 {
			result = append(result, d)
		}
	}
	return result
}
