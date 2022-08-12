package scout
import (
	"os"
	"ransomwarecheck/contractor"
	"ransomwarecheck/models"
)
func Visit(pathes string, f os.DirEntry, err error) error {
	//cd
	os.Chdir(pathes)
	//mkdir
	Dir, err := os.Open(pathes)
	if err != nil {
		// handle the error and return
	}
	defer Dir.Close()
	// This returns an *os.FileInfo type
	DirInfo, err := Dir.Stat()
	if err != nil {
		// error handling
	}
	// IsDir is short for fileInfo.Mode().IsDir()
	if DirInfo.IsDir() {
		// file is a directory
		models.Pathes = pathes
		contractor.Create()

	} else {
		// file is not a directory
	}
	return nil
}