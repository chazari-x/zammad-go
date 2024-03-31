package zammad

import "fmt"

func (c *Client) GroupList() (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/groups", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) GroupShow(groupID int) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/groups/%d", c.Url, groupID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) GroupCreate(g *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/groups", c.Url)
	req, err := c.NewRequest("POST", url, g)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) GroupUpdate(groupID int, g *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/groups/%d", c.Url, groupID)
	req, err := c.NewRequest("PUT", url, g)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) GroupDelete(groupID int) error {
	url := fmt.Sprintf("%s/api/v1/groups/%d", c.Url, groupID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
