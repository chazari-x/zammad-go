package zammad

import "fmt"

func (c *Client) ObjectList() (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) ObjectShow(objectID int) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes/%d", c.Url, objectID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) ObjectCreate(o *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes", c.Url)
	req, err := c.NewRequest("POST", url, o)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) ObjectUpdate(objectID int, o *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes/%d", c.Url, objectID)
	req, err := c.NewRequest("PUT", url, o)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) ObjectExecuteDatabaseMigration() error {
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes_execute_migrations", c.Url)
	req, err := c.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
