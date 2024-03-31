package zammad

import "fmt"

func (c *Client) TicketStateList() (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_states")
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketStateShow(ticketStateID int) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_states/%d", c.Url, ticketStateID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketStateCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_states", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketStateUpdate(ticketStateID int, t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_states/%d", c.Url, ticketStateID)
	req, err := c.NewRequest("PUT", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketStateDelete(ticketStateID int) error {
	url := fmt.Sprintf("%s/api/v1/ticket_states/%d", c.Url, ticketStateID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
