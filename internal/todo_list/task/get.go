package task

import (
	"net/http"
	"todo-list-api/pkg/logger"
)

func Get(w http.ResponseWriter, r *http.Request) {

	task, err := AccessTask(w, r)
	if err != nil {
		return
	}

	b, err := task.Marshal()
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "500 internal server error", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}
