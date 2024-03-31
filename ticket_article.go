package zammad

import "fmt"

func (c *Client) TicketArticleByTicket(ticketID int) (*[]map[string]interface{}, error) {
	var data []map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_articles/by_ticket/%d", c.Url, ticketID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketArticleShow(ticketArticleID int) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_articles/%d", c.Url, ticketArticleID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}

func (c *Client) TicketArticleCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_articles", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}
