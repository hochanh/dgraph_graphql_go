package apitest

import (
	"testing"

	"github.com/romshark/dgraph_graphql_go/apitest/setup"
	"github.com/stretchr/testify/require"
)

// TestCreateUserErr tests all possible user account creation errors
func TestCreateUserErr(t *testing.T) {
	// Test duplicate email on creation
	t.Run("duplicateEmail", func(t *testing.T) {
		ts := setup.New(t, tcx)
		defer ts.Teardown()

		ts.Help.OK.CreateUser("fooBarowich", "foo@bar.buz")
		res, errs := ts.Help.CreateUser("bazBazowich", "foo@bar.buz")
		require.Nil(t, res)
		require.Len(t, errs, 1)
	})
}
