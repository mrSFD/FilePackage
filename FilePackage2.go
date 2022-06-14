package FilePackage
import (
	"fmt"
	"io"
	"os"
)

var path = "G:\\CRWDF\\TestFile.txt"

func CreateFile() (string){
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { os.Exit(1) }
		defer file.Close()
	}

	return "file successfully created."
}

func ReadFile() (string){
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644) 
	if isError(err) { os.Exit(1) }
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		
		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}
		
		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}
	
	fmt.Println("done reading from file")
	return string(text)
}

func WriteFile() (string){
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if isError(err) { os.Exit(1) }
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString("halo\n")
	if isError(err) { os.Exit(1) }
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) { os.Exit(1) }

	// save changes
	err = file.Sync()
	if isError(err) { os.Exit(1) }

	return "done writing to file"
}

func DeleteFile() (string){
	// delete file
	var err = os.Remove(path)
	if isError(err) { os.Exit(1) }

	return "done deleting file"
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
