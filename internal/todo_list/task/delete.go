package task

import (
	"context"
	"net/http"
	"todo-list-api/pkg/logger"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	task, err := AccessTask(w, r)
	if err != nil {
		return
	}

	err = DeleteTaskByID(context.Background(), task.ID)
	if err != nil {
		logger.InfoLog.Println(err)
		http.Error(w, "500 internal server error", 500)
		return
	}

	w.WriteHeader(204)
}
