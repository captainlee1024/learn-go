package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

// HandleFileList filehandler
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		// return errors.New("path must start " +
		// 	"with " + prefix)
		return userError("path must start " +
			"with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		// panic(err)

		// http.Error(writer,
		// 	err.Error(),
		// 	http.StatusInternalServerError)
		// return

		return err
	}

	all, err := ioutil.ReadAll(file)

	if err != nil {
		// panic(err)
		return err
	}

	writer.Write(all)
	return err
}
