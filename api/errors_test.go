package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Error(t *testing.T) {
	t.Parallel()
	errorMessage := "Account not found at the requested version"
	errorCode := "AccountNotFound"
	vmErrorCode := uint64(0)

	testJson := fmt.Sprintf(`{
		"message": "%s",
		"error_code": "%s"
	}`, errorMessage, errorCode)
	data := &Error{}
	err := json.Unmarshal([]byte(testJson), &data)
	require.NoError(t, err)
	assert.Equal(t, errorMessage, data.Message)
	assert.Equal(t, errorCode, data.ErrorCode)
	assert.Equal(t, vmErrorCode, data.VmErrorCode)
}

func Test_ErrorWithVm(t *testing.T) {
	t.Parallel()
	errorMessage := "Invalid transaction: Type: The transaction has a bad signature Code: 1"
	errorCode := "VmError"
	vmErrorCode := uint64(1)

	testJson := fmt.Sprintf(`{
		"message": "%s",
		"error_code": "%s",
		"vm_error_code": %d
	}`, errorMessage, errorCode, vmErrorCode)
	data := &Error{}
	err := json.Unmarshal([]byte(testJson), &data)
	require.NoError(t, err)
	assert.Equal(t, errorMessage, data.Message)
	assert.Equal(t, errorCode, data.ErrorCode)
	assert.Equal(t, vmErrorCode, data.VmErrorCode)
}
