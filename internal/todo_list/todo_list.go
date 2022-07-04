package todo_list

import (
	"context"
	"fmt"
	. "todo-list-api/config"
)

func GetUserTasks(ctx context.Context, id int) ([]int, error) {
	rows, err := DBPool.Query(ctx, fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1;", DBTask.ID, DBTask.Table, DBTask.UserID), id)
	if err != nil {
		return nil, err
	}

	var (
		arr    []int
		taskID int
	)

	fmt.Println(fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1;", DBTask.ID, DBTask.Table, DBTask.UserID), id)

	for f := rows.Next(); f == true; {
		err := rows.Scan(&taskID)
		if err != nil {
			return arr, nil
		}
		arr = append(arr, taskID)
		f = rows.Next()
	}

	return arr, nil
}
