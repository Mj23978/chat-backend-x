package cobra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	logger "github.com/mj23978/chat-backend-x/logger/zerolog"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/pkg/errors"
)

var (
	// ErrNilDependency is returned if a dependency is missing.
	ErrNilDependency = errors.New("a dependency was expected to be defined but is nil. Please open an issue with the stack trace")
	// ErrNoPrintButFail is returned to detect a failure state that was already reported to the user in some way
	ErrNoPrintButFail = errors.New("this error should never be printed")
)

// FailSilently is supposed to be used within a commands RunE function.
// It silences cobras error handling and returns the ErrNoPrintButFail error.
func FailSilently(cmd *cobra.Command) error {
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	return errors.WithStack(ErrNoPrintButFail)
}

// Must fatals with the optional message if err is not nil.
func Must(err error, message string, args ...interface{}) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	os.Exit(1)
}

// CheckResponse fatals if err is nil or the response.StatusCode does not match the expectedStatusCode
func CheckResponse(err error, expectedStatusCode int, response *http.Response) {
	Must(err, "Command failed because error occurred: %s", err)

	if response.StatusCode != expectedStatusCode {
		out, err := ioutil.ReadAll(response.Body)
		if err != nil {
			out = []byte{}
		}
		pretty, err := json.MarshalIndent(json.RawMessage(out), "", "\t")
		if err == nil {
			out = pretty
		}

		Fatalf(
			`Command failed because status code %d was expected but code %d was received.

Response payload:

%s`,
			expectedStatusCode,
			response.StatusCode,
			out,
		)
	}
}

// FormatResponse takes an object and prints a json.MarshalIdent version of it or fatals.
func FormatResponse(o interface{}) string {
	out, err := json.MarshalIndent(o, "", "\t")
	Must(err, `Command failed because an error occurred while prettifying output: %s`, err)
	return string(out)
}

// Fatalf prints to os.Stderr and exists with code 1.
func Fatalf(message string, args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, message)
	}
	os.Exit(1)
}

// ExpectDependency expects every dependency to be not nil or it fatals.
func ExpectDependency(dependencies ...interface{}) {
	for _, d := range dependencies {
		if d == nil {
			logger.Errorf("A Fatal Issue : %s", errors.WithStack(ErrNilDependency))
		}
	}
}

// Exec runs the provided cobra command with the given reader as STD_IN and the given args.
// Returns STD_OUT, STD_ERR and the error from the execution.
func Exec(_ *testing.T, cmd *cobra.Command, stdIn io.Reader, args ...string) (string, string, error) {
	stdOut, stdErr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.SetErr(stdErr)
	cmd.SetOut(stdOut)
	cmd.SetIn(stdIn)
	defer cmd.SetIn(nil)
	if args == nil {
		args = []string{}
	}
	cmd.SetArgs(args)
	err := cmd.Execute()
	return stdOut.String(), stdErr.String(), err
}

// ExecNoErr is a helper that assumes a successful run from Exec.
// Returns STD_OUT.
func ExecNoErr(t *testing.T, cmd *cobra.Command, args ...string) string {
	stdOut, stdErr, err := Exec(t, cmd, nil, args...)
	require.NoError(t, err)
	require.Len(t, stdErr, 0, stdOut)
	return stdOut
}

// ExecExpectedErr is a helper that assumes a failing run from Exec returning ErrNoPrintButFail
// Returns STD_ERR.
func ExecExpectedErr(t *testing.T, cmd *cobra.Command, args ...string) string {
	stdOut, stdErr, err := Exec(t, cmd, nil, args...)
	require.True(t, errors.Is(err, ErrNoPrintButFail))
	require.Len(t, stdOut, 0, stdErr)
	return stdErr
}
