package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_astCmd(t *testing.T) {
	os.Setenv("ASAR_WORKSPACE_ID", "111")
	os.Setenv("ASAR_TOKEN", "1/token")
	os.Setenv("ASAR_ASSIGNEE_ID", "1127028194303123")
	os.Setenv("ASAR_ASANA_URL", "http://localhost:3500")

	var (
		astCmd  = createAstCmd()
		argsTmp = []string{"ast"}
		buffTmp = new(bytes.Buffer)

		expect string
		actual string
	)

	astCmd.SetOut(buffTmp)
	astCmd.SetArgs(argsTmp)

	if err := astCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'astCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = "[Configuration]\n  GLOBAL_CONFIG: {WorkspaceId:111 Token:1/token AsanaUrl:http://localhost:3500 AssigneeId:1127028194303123}\n================================================\nProject: Learning in progress...\nDone.\n================================================\nProject: Vuejs+Lalarvel in progress...\nDone.\n================================================\nAll Done.\n==== Short Report ====\nSumTask: 4\nSumCompleted: 2\nSumSubTask: 2\nSumSubTaskCompleted: 1\n"

	actual = buffTmp.String()

	assert.Equal(t, expect, actual,
		"Command 'asar' should return 'asar development'.",
	)
}
