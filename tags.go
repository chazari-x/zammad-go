package zammad

import "fmt"

func (c *Client) TagList(ticketID int) (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tags?object=Ticket&o_id=%d", c.Url, ticketID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TagSearch(term string) (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tag_search?term=%s", c.Url, term)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TagAdd(t *map[string]interface{}) error {
	url := fmt.Sprintf("%s/api/v1/tags/add", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}

func (c *Client) TagRemove(t *map[string]interface{}) error {
	url := fmt.Sprintf("%s/api/v1/tags/remove", c.Url)
	req, err := c.NewRequest("DELETE", url, t)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}

func (c *Client) TagAdminList() (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tag_list", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TagAdminCreate(o *map[string]interface{}) error {
	url := fmt.Sprintf("%s/api/v1/tag_list", c.Url)
	req, err := c.NewRequest("POST", url, o)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}

func (c *Client) TagAdminRename(tagID int, t *map[string]interface{}) error {
	url := fmt.Sprintf("%s/api/v1/tag_list/%d", c.Url, tagID)
	req, err := c.NewRequest("PUT", url, t)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}

func (c *Client) TagAdminDelete(tagID int) error {
	url := fmt.Sprintf("%s/api/v1/tag_list/%d", c.Url, tagID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
