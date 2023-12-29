package ad_test

import (
	"errors"
	"testing"

	"github.com/error-re-solution/ad-solution/ad"
	"github.com/error-re-solution/ad-solution/config"
)

func Test_GetAllUsers(t *testing.T) {
	type testCase struct {
		test          string
		ou            string
		expectedError error
	}

	ldap, err := config.LoadLDAPConfig("../")
	if err != nil {
		panic(err)
	}

	client, err := ad.NewADClient(ldap.Address, ldap.BindDN, ldap.BindPassword)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	testCases := []testCase{
		{
			test:          "testing - GetAllUsers",
			ou:            "customOU", // Valid Organization Unit
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := client.GetAllUsers(tc.ou)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
