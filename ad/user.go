package ad

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

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

// GetUserByEmail retrieves a user from Active Directory based on the email address.
func (c *ADClient) GetUserByEmail(email string) (*ADUser, error) {
	searchFilter := fmt.Sprintf("(mail=%s)", email)
	attributes := []string{"dn", "cn", "givenName", "sn", "mail", "memberOf"}

	searchResult, err := c.Search("DC=xcompany,DC=local", searchFilter, attributes)
	if err != nil {
		return nil, err
	}

	if len(searchResult.Entries) == 0 {
		return nil, ErrUserNotFound
	}

	entry := searchResult.Entries[0]
	user := &ADUser{
		FirstName: entry.GetAttributeValue("givenName"),
		LastName:  entry.GetAttributeValue("sn"),
		Email:     entry.GetAttributeValue("mail"),
		MemberOf:  entry.GetAttributeValues("memberOf"),
	}

	return user, nil
}

// AddUser adds a new user to Active Directory.
// !TOFIX: test fails when trying to add a new user with a password.
func (c *ADClient) AddUser(username /*password*/, email, ou, baseDN, domain string) error {
	dn := fmt.Sprintf("CN=%s,OU=%s,%s", username, ou, baseDN)
	fmt.Printf("DN: %s", dn)
	entry := ldap.NewAddRequest(dn, nil)

	entry.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "user"})
	entry.Attribute("cn", []string{username})
	entry.Attribute("sAMAccountName", []string{username})
	entry.Attribute("userPrincipalName", []string{username + "@" + domain})
	entry.Attribute("mail", []string{email})

	// encodedPassword, err := encodeUTF16LE(password)
	// if err != nil {
	// 	return err
	// }
	// entry.Attribute("unicodePwd", []string{string(encodedPassword)})

	return c.Connection.Add(entry)
}

// func encodeUTF16LE(s string) ([]byte, error) {
// 	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
// 	utf16Encoded, err := encoder.Bytes([]byte(s))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Remove UTF-16 BOM (Byte Order Mark)
// 	if len(utf16Encoded) >= 2 && utf16Encoded[0] == 0xFE && utf16Encoded[1] == 0xFF {
// 		utf16Encoded = utf16Encoded[2:]
// 	}

// 	return utf16Encoded, nil
// }
