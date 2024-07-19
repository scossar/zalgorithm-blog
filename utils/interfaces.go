package utils

type FileFetcher interface {
	FilesOfType(dir, fileType string) ([]string, error)
}
