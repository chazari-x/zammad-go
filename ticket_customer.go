package zammad

import "fmt"

func (c *Client) TicketListByCustomer(customer int) (*map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/api/v1/ticket_customer?customer_id=%d", c.Url, customer)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return &data, err
	}

	err = c.SendWithAuth(req, &data)
	return &data, err
}
