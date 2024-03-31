package zammad

import "fmt"

func (c *Client) TicketList() (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tickets", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketSearch(query string, limit int) (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tickets/search?query=%s&limit=%d", c.Url, query, limit)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketShow(ticketID int) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tickets/%d", c.Url, ticketID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tickets", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketUpdate(ticketID int, t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/tickets/%d", c.Url, ticketID)
	req, err := c.NewRequest("PUT", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketDelete(ticketID int) error {
	url := fmt.Sprintf("%s/api/v1/tickets/%d", c.Url, ticketID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
