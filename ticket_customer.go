package zammad

import "fmt"

func (c *Client) TicketListByCustomer(customer int) (data *[]map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/ticket_customer?customer_id=%d", c.Url, customer)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	return data, c.SendWithAuth(req, data)
}
