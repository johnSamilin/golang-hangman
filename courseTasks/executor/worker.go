package executor

import (
	"fmt"
	"strconv"
)

func (w *Worker) Work(configId int, task ITask) {
	commonInfo := "Worker #" + strconv.Itoa(w.Id) + " (task from cmd pool #" + strconv.Itoa(configId) + ") executing " + task.String()
	w.Lock.Lock()
	defer w.Lock.Unlock()
	var result string
	w.IsActive, result = task.Execute()
	fmt.Println(commonInfo, ":", result)
}
