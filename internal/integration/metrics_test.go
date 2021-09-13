package integration_test

import (
	"bytes"
	"encoding/json"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationGetMetrics(t *testing.T) {
	for testname, id := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			client := &MdwClient{}
			mtrcs, err := goieeeapi.GetMetrics(client, id)

			// Ensure no error
			assert.Nil(err)

			// Reencode to JSON
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(mtrcs)

			// Decode both to interfaces
			var expected interface{}
			var actual interface{}
			json.Unmarshal(client.LastBody, &expected)
			json.Unmarshal(buf.Bytes(), &actual)

			// Compares both to check valid API (and not nil)
			assert.NotNil(expected)
			assert.Equal(expected, actual)
		})
	}
}
