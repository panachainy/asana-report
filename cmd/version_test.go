package cmd

import (
	"asana-report/cmd"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_versionCmd(t *testing.T) {
	cmd.VERSION = "0.0.0"

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

	// expect = "asana-report" + " " + `\n`
	expect = "asana-report\n"

	actual = buffTmp.String()

	fmt.Println("actual =========================")
	fmt.Println(actual)
	fmt.Println("expect =========================")
	fmt.Println(expect)

	assert.Equal(t, expect, actual,
		"Command 'asana-report' should return empty.",
	)
}
