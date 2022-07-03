package todo_list

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-list-api/internal/jwt"
)

func Get(w http.ResponseWriter, r *http.Request) {
	err := jwt.Auth(w, r)
	if err != nil {
		return
	}

	//count, offset, err := ParseQuery(r)
	//if err != nil {
	//	logger.InfoLog.Println(err)
	//	http.Error(w, err.Error(), 400)
	//}
}

func ParseQuery(r *http.Request) (int, int, error) {
	q := r.URL.Query()

	if len(q["count"]) != 1 || len(q["offset"]) != 1 {
		return 0, 0, fmt.Errorf("bad query")
	}

	q1, q2 := q["count"][0], q["offset"][0]

	count, err := strconv.Atoi(q1)
	if err != nil {
		return 0, 0, fmt.Errorf("bad query")
	}

	offset, err := strconv.Atoi(q2)
	if err != nil {
		return 0, 0, fmt.Errorf("bad query")
	}

	return count, offset, nil
}
