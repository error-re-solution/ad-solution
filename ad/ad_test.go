package ad_test

import (
	"errors"
	"testing"

	"github.com/error-re-solution/ad-solution/ad"
	"github.com/error-re-solution/ad-solution/config"
)

type testCase struct {
	name          string
	LDAPAddress   string
	bindDN        string
	bindPass      string
	expectedError error
}

func TestADClient(t *testing.T) {
	ldap, err := config.LoadLDAPConfig("../")
	if err != nil {
		panic(err)
	}

	testCases := []testCase{
		{
			name:          "testing - AD client",
			LDAPAddress:   ldap.Address,
			bindDN:        ldap.BindDN,
			bindPass:      ldap.BindPassword,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ad.NewADClient(tc.LDAPAddress, tc.bindDN, tc.bindPass)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error: %v, got %v", tc.expectedError, err)
			}
		})
	}

}
