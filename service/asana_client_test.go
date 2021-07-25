package service

import (
	"asana-report/model"
	"reflect"
	"testing"
)

func Test_GetTasks(t *testing.T) {
	InitService("http://localhost:3500")

	tests := []struct {
		name         string
		projectId    string
		token        string
		mockFunc     func()
		expected     model.Tasks
		expectingErr bool
	}{
		{
			name:      "All success no error",
			projectId: "1200195508767749",
			token:     "s",
			mockFunc: func() {
			},
			expected: model.GetTasksBy(`
			{
				"data": [
					{
						"gid": "1200616651608041",
						"assignee": {
							"gid": "1127028096108896",
							"resource_type": "user"
						},
						"completed": false,
						"name": "quicktype-vscode",
						"num_subtasks": 0
					},
					{
						"gid": "1200195560177014",
						"assignee": {
							"gid": "1127028096108896",
							"resource_type": "user"
						},
						"completed": false,
						"name": "เรียบเรียงว่าจะเขียนอะไรบ้าง",
						"num_subtasks": 2
					}
				]
			}
			`),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()
			tasksResult := GetTasks(tc.projectId, tc.token)

			// errExist := err != nil
			// if tc.expectingErr != errExist {
			// 	tt.Errorf("Error expectation not met, want %v, got %v", tc.expectingErr, errExist)
			// }

			if !reflect.DeepEqual(tc.expected, tasksResult) {
				tt.Errorf("Error, tasks expectation not met, want %+v, got %+v", tc.expected, tasksResult)
			}
		})
	}
}
