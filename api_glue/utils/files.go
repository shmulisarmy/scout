package utils

import "os"

func Load_from_file(file_name string) string {
	file, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	return string(file)
}
