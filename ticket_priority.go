package zammad

import "fmt"

func (c *Client) TicketPriorityList() (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities")
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketPriorityShow(ticketPriorityID int) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_priorities/%d", c.Url, ticketPriorityID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketPriorityCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities")
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketPriorityUpdate(ticketPriorityID int, t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_priorities/%d", c.Url, ticketPriorityID)
	req, err := c.NewRequest("PUT", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketPriorityDelete(ticketPriorityID int) error {
	url := fmt.Sprintf("%s/api/v1/ticket_priorities/%d", c.Url, ticketPriorityID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
