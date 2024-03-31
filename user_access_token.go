package zammad

import "fmt"

func (c *Client) UserAccessTokenList() (data *[]map[string]interface{}, err error) {
	data = &[]map[string]interface{}{}
	url := fmt.Sprintf("%s/api/v1/user_access_token", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserAccessTokenCreate(t *map[string]interface{}) (data *map[string]interface{}, err error) {
	data = &map[string]interface{}{}
	url := fmt.Sprintf("%s/api/v1/user_access_token", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserAccessTokenDelete(tokenID int) error {
	url := fmt.Sprintf("%s/api/v1/user_access_token/%d", c.Url, tokenID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
