package zammad

import "fmt"

func (c *Client) OnlineNotificationList() (data []map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/online_notifications", c.Url)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, &data)
	return
}

func (c *Client) OnlineNotificationShow(notificationID int) (data map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/online_notifications/%d", c.Url, notificationID)
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, &data)
	return
}

func (c *Client) OnlineNotificationUpdate(notificationID int, n *map[string]interface{}) (data map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/api/v1/online_notifications/%d", c.Url, notificationID)
	req, err := c.NewRequest("PUT", url, n)
	if err != nil {
		return
	}

	err = c.SendWithAuth(req, &data)
	return
}

func (c *Client) OnlineNotificationDelete(notificationID int) error {
	url := fmt.Sprintf("%s/api/v1/online_notifications/%d", c.Url, notificationID)
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}

func (c *Client) OnlineNotificationMarkAllAsRead() error {
	url := fmt.Sprintf("%s/api/v1/online_notifications/mark_all_as_read", c.Url)
	req, err := c.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAuth(req, nil)
}
