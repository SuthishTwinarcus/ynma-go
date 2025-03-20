package yonoma

import "fmt"

type List struct {
	Name string `json:"list_name"`
}

func (c *Client) ListLists() ([]byte, error) {
	return c.request("GET", "lists/list", nil)
}

func (c *Client) CreateList(list List) ([]byte, error) {
	return c.request("POST", "lists/create", list)
}

func (c *Client) UpdateList(listID string, list List) ([]byte, error) {
	endpoint := fmt.Sprintf("lists/%s/update", listID)
	return c.request("POST", endpoint, list)
}

func (c *Client) RetrieveList(listID string) ([]byte, error) {
	endpoint := fmt.Sprintf("lists/%s", listID)
	return c.request("GET", endpoint, nil)
}

func (c *Client) DeleteList(listID string) ([]byte, error) {
	endpoint := fmt.Sprintf("lists/%s/delete", listID)
	return c.request("DELETE", endpoint, nil)
}
