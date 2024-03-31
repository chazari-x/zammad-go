package zammad

import "fmt"

func (c *Client) ObjectList() (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) ObjectShow(objectID int) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes/%d", c.Url, objectID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) ObjectCreate(o *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes", c.Url)
	req, err := c.NewRequest("POST", url, o)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) ObjectUpdate(objectID int, o *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes/%d", c.Url, objectID)
	req, err := c.NewRequest("PUT", url, o)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) ObjectExecuteDatabaseMigration() error {
	url := fmt.Sprintf("%s/api/v1/object_manager_attributes_execute_migrations", c.Url)
	req, err := c.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
