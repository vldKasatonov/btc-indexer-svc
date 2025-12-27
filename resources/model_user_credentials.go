/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserCredentials struct {
	Key
	Attributes UserCredentialsAttributes `json:"attributes"`
}
type UserCredentialsResponse struct {
	Data     UserCredentials `json:"data"`
	Included Included        `json:"included"`
}

type UserCredentialsListResponse struct {
	Data     []UserCredentials `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustUserCredentials - returns UserCredentials from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserCredentials(key Key) *UserCredentials {
	var userCredentials UserCredentials
	if c.tryFindEntry(key, &userCredentials) {
		return &userCredentials
	}
	return nil
}
