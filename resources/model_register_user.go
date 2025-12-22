/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RegisterUser struct {
	Key
	Attributes RegisterUserAttributes `json:"attributes"`
}
type RegisterUserResponse struct {
	Data     RegisterUser `json:"data"`
	Included Included     `json:"included"`
}

type RegisterUserListResponse struct {
	Data     []RegisterUser `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustRegisterUser - returns RegisterUser from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRegisterUser(key Key) *RegisterUser {
	var registerUser RegisterUser
	if c.tryFindEntry(key, &registerUser) {
		return &registerUser
	}
	return nil
}
