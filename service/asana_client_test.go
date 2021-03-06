package service

import (
	"asana-report/model"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTasks(t *testing.T) {
	InitService("http://localhost:3500")

	tests := []struct {
		name            string
		projectId       string
		token           string
		mockFunc        func()
		expected        model.Tasks
		expectingErr    bool
		expectingErrMsg string
	}{
		{
			name:      "Success",
			projectId: "111",
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
		{
			name:      "404 Not found",
			projectId: "wrong-project-id",
			token:     "",
			mockFunc: func() {
			},
			expectingErr:    true,
			expectingErrMsg: "Something wrong from asana status code is 404 at GetTasks()\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()

			if tc.expectingErr {
				assert.PanicsWithValue(t, tc.expectingErrMsg, func() { GetTasks(tc.projectId, tc.token) }, "The code did not panic or mistake message of panic")
				return
			}

			tasksResult := GetTasks(tc.projectId, tc.token)

			if !reflect.DeepEqual(tc.expected, tasksResult) {
				tt.Errorf("Error, tasks expectation not met, want %+v, got %+v", tc.expected, tasksResult)
			}
		})
	}
}

func Test_GetSubTasks(t *testing.T) {
	InitService("http://localhost:3500")

	tests := []struct {
		name            string
		projectId       string
		token           string
		mockFunc        func()
		expected        model.Tasks
		expectingErr    bool
		expectingErrMsg string
	}{
		{
			name:      "Success",
			projectId: "111",
			token:     "",
			mockFunc: func() {
			},
			expected: model.GetTasksBy(`
			{
				"data": [
					{
						"gid": "1200653409861926",
						"assignee": {
							"gid": "1127028096108896",
							"resource_type": "user"
						},
						"completed": true,
						"name": "cobra-101 (part 1 setup project)"
					},
					{
						"gid": "1200653409861948",
						"assignee": {
							"gid": "1127028096108896",
							"resource_type": "user"
						},
						"completed": false,
						"name": "cobra-101 (part 2 create command)"
					}
				]
			}
			`),
		},
		{
			name:      "404 Not found",
			projectId: "wrong-project-id",
			token:     "",
			mockFunc: func() {
			},
			expectingErr:    true,
			expectingErrMsg: "Something wrong from asana status code is 404 at GetSubTasks()\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()

			if tc.expectingErr {
				assert.PanicsWithValue(t, tc.expectingErrMsg, func() { GetSubTasks(tc.projectId, tc.token) }, "The code did not panic or mistake message of panic")
				return
			}

			tasksResult := GetSubTasks(tc.projectId, tc.token)

			if !reflect.DeepEqual(tc.expected, tasksResult) {
				tt.Errorf("Error, tasks expectation not met, want %+v, got %+v", tc.expected, tasksResult)
			}
		})
	}
}

func Test_GetWorkspace(t *testing.T) {
	InitService("http://localhost:3500")

	tests := []struct {
		name            string
		workspaceId     string
		token           string
		mockFunc        func()
		expected        model.Workspace
		expectingErr    bool
		expectingErrMsg string
	}{
		{
			name:        "Success",
			workspaceId: "111",
			token:       "",
			mockFunc: func() {
			},
			expected: model.GetWorkspaceBy(`
			{
				"data": [
					{
						"gid": "1200195508767749",
						"name": "Learning",
						"resource_type": "project"
					},
					{
						"gid": "1200485948764582",
						"name": "Vuejs+Lalarvel",
						"resource_type": "project"
					}
				],
				"next_page": null
			}
			`),
		},
		{
			name:        "404 Not found",
			workspaceId: "wrong-workspace-id",
			token:       "",
			mockFunc: func() {
			},
			expectingErr:    true,
			expectingErrMsg: "Something wrong from asana status code is 404 at getWorkspace()\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()

			if tc.expectingErr {
				assert.PanicsWithValue(t, tc.expectingErrMsg, func() { GetWorkspace(tc.workspaceId, tc.token) }, "The code did not panic or mistake message of panic")
				return
			}

			tasksResult := GetWorkspace(tc.workspaceId, tc.token)

			if !reflect.DeepEqual(tc.expected, tasksResult) {
				tt.Errorf("Error, tasks expectation not met, want %+v, got %+v", tc.expected, tasksResult)
			}
		})
	}
}

func Test_UpdateTasks_Success(t *testing.T) {
	InitService("http://localhost:3500")

	tests := []struct {
		name            string
		taskId          string
		assigneeId      string
		token           string
		mockFunc        func()
		expectingErr    bool
		expectingErrMsg string
	}{
		{
			name:       "Success",
			taskId:     "111",
			assigneeId: "assignee-mock",
			token:      "111",
			mockFunc: func() {
			},
		},
		{
			name:       "404 Not found",
			taskId:     "worng-task-id",
			assigneeId: "assignee-mock",
			token:      "",
			mockFunc: func() {
			},
			expectingErr:    true,
			expectingErrMsg: "Something wrong from asana status code is 404 at UpdateTasks()\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockFunc()

			if tc.expectingErr {
				assert.PanicsWithValue(t, tc.expectingErrMsg, func() { UpdateTasks(tc.taskId, tc.assigneeId, tc.token) }, "The code did not panic or mistake message of panic")
				return
			}

			UpdateTasks(tc.taskId, tc.assigneeId, tc.token)
		})
	}
}
