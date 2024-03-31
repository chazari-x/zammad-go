package zammad

import "fmt"

func (c *Client) TicketList() (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/tickets", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketSearch(query string, limit int) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/tickets/search?query=%s&limit=%d", c.Url, query, limit)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketShow(ticketID int) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/tickets/%d", c.Url, ticketID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketCreate(t *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/tickets", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketUpdate(ticketID int, t *map[string]interface{}) (data *map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/tickets/%d", c.Url, ticketID)
	req, err := c.NewRequest("PUT", url, t)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketDelete(ticketID int) error {
	url := fmt.Sprintf("%s/api/v1/tickets/%d", c.Url, ticketID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
