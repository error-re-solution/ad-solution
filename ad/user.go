package ad

import "fmt"

type ADUser struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	MemberOf  []string `json:"memberOf"`
}

// GetAllUsers retrieves all users from the specified Organizational Unit (OU) in Active Directory.
func (c *ADClient) GetAllUsers(ou string) ([]*ADUser, error) {
	searchFilter := "(objectClass=user)"
	attributes := []string{"dn", "cn", "givenName", "sn", "mail", "memberOf"}

	searchBase := fmt.Sprintf("OU=%s,DC=xcompany,DC=local", ou)
	searchResult, err := c.Search(searchBase, searchFilter, attributes)
	if err != nil {
		return nil, err
	}

	users := make([]*ADUser, len(searchResult.Entries))
	for i, entry := range searchResult.Entries {
		user := &ADUser{
			FirstName: entry.GetAttributeValue("givenName"),
			LastName:  entry.GetAttributeValue("sn"),
			Email:     entry.GetAttributeValue("mail"),
			MemberOf:  entry.GetAttributeValues("memberOf"),
		}
		users[i] = user
	}

	return users, nil
}
