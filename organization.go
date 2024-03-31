package zammad

import "fmt"

func (c *Client) OrganizationList() (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/organizations", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) OrganizationSearch(query string, limit int) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/organizations/search?query=%s&limit=%d", c.Url, query, limit)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) OrganizationShow(organizationID int) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/organizations/%d", c.Url, organizationID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) OrganizationCreate(o *map[string]interface{}) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/organizations", c.Url)
	req, err := c.NewRequest("POST", url, o)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) OrganizationUpdate(organizationID int, o *map[string]interface{}) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/organizations/%d", c.Url, organizationID)
	req, err := c.NewRequest("PUT", url, o)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) OrganizationDelete(organizationID int) error {
	url := fmt.Sprintf("%s/api/v1/organizations/%d", c.Url, organizationID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
