package ad

import "github.com/go-ldap/ldap/v3"

type ADClient struct {
	Connection *ldap.Conn
}

// NewADClient initializes a new Active Directory client.
func NewADClient(ldapServer, bindDN, bindPass string) (*ADClient, error) {
	l, err := ldap.Dial("tcp", ldapServer)
	if err != nil {
		return nil, err
	}

	err = l.Bind(bindDN, bindPass)
	if err != nil {
		return nil, err
	}

	return &ADClient{Connection: l}, nil
}

// Close closes the LDAP connection.
func (c *ADClient) Close() {
	if c.Connection != nil {
		c.Connection.Close()
	}
}
