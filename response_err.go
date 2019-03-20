package communication

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func ResponseErrMarshalJSON(w http.ResponseWriter, data interface{}) { spew.Dump(data) }
func ResponseErr(w http.ResponseWriter, data interface{})            { spew.Dump(data) }

func ResponseErrNotFound(w http.ResponseWriter) {
	setHeaderForJSON(w)
	w.WriteHeader(http.StatusNotFound)
}
func ResponseErrNoResponse(w http.ResponseWriter) {
	setHeaderForJSON(w)
	w.WriteHeader(http.StatusRequestTimeout)
}
func ResponseErrorTokenInvalid(w http.ResponseWriter) {
	setHeaderForJSON(w)
	w.WriteHeader(http.StatusUnauthorized)
}
func ResponseErrorMarhsallingJSON(w http.ResponseWriter) {
	setHeaderForJSON(w)
	w.WriteHeader(http.StatusUnprocessableEntity)
	log.WithFields(getErrorLoggingFields(3)).Debug("error marshalling json")
}

func generateErrorID() string {
	return time.Now().String()
}

func getErrorLoggingFields(count int) (fields log.Fields) {
	fields = getFormatedCallTrace(count)
	fields["eid"] = generateErrorID()
	fields["ts"] = time.Now()
	return
}

func getFormatedCallTrace(count int) (fields log.Fields) {
	fields = make(map[string]interface{})
	pc := make([]uintptr, count)
	filled := runtime.Callers(2, pc)
	if filled > count {
		filled = count
	}
	for i := 0; i < filled; i++ {
		prefix := fmt.Sprintf("err%d_", i)
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		// fields[prefix+"file"] = file
		// fields[prefix+"line"] = line
		// fields[prefix+"func"] = f.Name()
		fields[prefix] = fmt.Sprintf("%s:%d %s", file, line, f.Name())
	}
	return
}
