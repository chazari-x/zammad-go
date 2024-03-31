package zammad

import "fmt"

func (c *Client) TicketStateList() (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_states")
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) TicketStateShow(ticketStateID int) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/ticket_states/%d", c.Url, ticketStateID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) TicketStateCreate(t *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/ticket_states", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) TicketStateUpdate(ticketStateID int, t *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/ticket_states/%d", c.Url, ticketStateID)
	req, err := c.NewRequest("PUT", url, t)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}

func (c *Client) TicketStateDelete(ticketStateID int) error {
	url := fmt.Sprintf("%s/api/v1/ticket_states/%d", c.Url, ticketStateID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
