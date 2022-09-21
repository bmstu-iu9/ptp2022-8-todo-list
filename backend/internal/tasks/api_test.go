package tasks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/entity"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/errors"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/log"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/router"
	"github.com/bmstu-iu9/ptp2022-8-todo-list/backend/internal/test"
)

func TestApi(t *testing.T) {
	logger := log.New()
	mux := *router.New(logger)
	s := NewService(&mockRepository{
		items:    []entity.Task{task_examples[0]},
		task_id:  2,
		label_id: 2,
	})

	RegisterHandlers(&mux, s, logger)

	toJson := func(data interface{}) string {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(data)
		if err != nil {
			panic(err)
		}
		return buf.String()
	}

	test_cases := []test.ApiTestCase{
		{
			Name:     "get",
			Method:   "GET",
			Url:      "/users/1/tasks",
			WantBody: toJson([]entity.Task{task_examples[0]}),
			WantCode: http.StatusOK,
		},
		{
			Name:     "get by id",
			Method:   "GET",
			Url:      "/users/1/tasks/1",
			WantBody: toJson(task_examples[0]),
			WantCode: http.StatusOK,
		},
		{
			Name:       "create",
			Method:     "POST",
			Url:        "/users/1/tasks",
			Body:       toJson(task_examples[3]),
			WantCode:   http.StatusCreated,
			WantHeader: http.Header{"Location": {"https://ptp.starovoytovai.ru/api/v1/users/1/tasks/2"}},
		},
		{
			Name:     "update",
			Method:   "PATCH",
			Url:      "/users/1/tasks/1",
			Body:     toJson(task_examples[2]),
			WantCode: http.StatusOK,
			WantBody: toJson(task_examples[1]),
		},
		{
			Name:     "delete",
			Method:   "DELETE",
			Url:      "/users/1/tasks/2",
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "get after delete",
			Method:   "GET",
			Url:      "/users/1/tasks/2",
			WantCode: http.StatusNotFound,
			WantBody: toJson(errors.Problem{
				Title:  "Not found",
				Status: http.StatusNotFound,
			}),
		},
	}

	test.Endpoint(t, test_cases, &mux)
}
