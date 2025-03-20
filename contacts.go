package yonoma

import "fmt"

type Contact struct {
	Email  string      `json:"email"`
	Status string      `json:"status"`
	Data   ContactData `json:"data"`
}
type ContactData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Zipcode   string `json:"zipcode"`
}
type TagId struct {
	TagId string `json:"tag_id"`
}
type Status struct {
	Status string `json:"status"`
}

func (c *Client) CreateContact(listID string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/create", listID)
	return c.request("POST", endpoint, contact)
}

func (c *Client) UnsubscribeContact(listID string, contactId string, status Status) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/status/%s", listID, contactId)
	return c.request("POST", endpoint, status)
}

func (c *Client) AddContactTag(contactId string, tagId TagId) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/add", contactId)
	return c.request("POST", endpoint, tagId)
}

func (c *Client) RemoveContactTag(contactId string, tagId TagId) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/remove", contactId)
	return c.request("POST", endpoint, tagId)
}

