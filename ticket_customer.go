package zammad

import "fmt"

func (c *Client) TicketListByCustomer(customer int) (data *map[string]interface{}, err error) {
	data = &map[string]interface{}{}
	url := fmt.Sprintf("%s/api/v1/ticket_customer?customer_id=%d", c.Url, customer)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, data)
	return
}
