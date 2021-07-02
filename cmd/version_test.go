package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_versionCmd(t *testing.T) {
	var (
		versionCmd = createVersionCmd()
		argsTmp    = []string{"version"}
		buffTmp    = new(bytes.Buffer)

		expect string
		actual string
	)

	versionCmd.SetOut(buffTmp)  // set output from os.Stdout -> buffTmp
	versionCmd.SetArgs(argsTmp) // set command args

	if err := versionCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'helloCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = "asar development\n"

	actual = buffTmp.String()

	assert.Equal(t, expect, actual,
		"Command 'asar' should return 'asar development'.",
	)
}
