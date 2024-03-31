package zammad

import "fmt"

func (c *Client) TicketArticleByTicket(ticketID int) (data *[]map[string]interface{}, err error) {
	data = &[]map[string]interface{}{}
	url := fmt.Sprintf("%s/api/v1/ticket_articles/by_ticket/%d", c.Url, ticketID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketArticleShow(ticketArticleID int) (data *map[string]interface{}, err error) {
	data = &map[string]interface{}{}
	url := fmt.Sprintf("%s/api/v1/ticket_articles/%d", c.Url, ticketArticleID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}

func (c *Client) TicketArticleCreate(t *map[string]interface{}) (data *map[string]interface{}, err error) {
	data = &map[string]interface{}{}
	url := fmt.Sprintf("%s/api/v1/ticket_articles", c.Url)
	req, err := c.NewRequest("POST", url, t)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}
