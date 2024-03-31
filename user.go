package zammad

import "fmt"

func (c *Client) UserMe() (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/users/me", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserList() (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/users", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserSearch(query string, limit int) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/users/search?query=%s&limit=%d", c.Url, query, limit)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserShow(userID int) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/users/%d", c.Url, userID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserCreate(u *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/users", c.Url)
	req, err := c.NewRequest("POST", url, u)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserUpdate(userID int, u *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/users/%d", c.Url, userID)
	req, err := c.NewRequest("PUT", url, u)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) UserDelete(userID int) error {
	url := fmt.Sprintf("%s/api/v1/users/%d", c.Url, userID)
	req, err := c.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
