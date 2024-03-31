package zammad

import "fmt"

func (c *Client) UserAccessTokenList() (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/user_access_token", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) UserAccessTokenCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/user_access_token", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) UserAccessTokenDelete(tokenID int) error {
	url := fmt.Sprintf("%s/api/v1/user_access_token/%d", c.Url, tokenID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
