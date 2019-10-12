package namecalculator

import "os"

type Interface interface {
	Rename(os.FileInfo) (string, error)
}
