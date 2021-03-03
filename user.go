package tipeee

import (
	"strconv"
)

// Me returns connected user
func (c *Client) Me() (*User, error) {
	resp := &User{}
	err := c.apiget("/partners/me/profile", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetTipsResponse struct {
	Pager *Pager `json:"pager"`
	Tips  []*Tip `json:"items"`
}

// GetTips lists connected user's tips.
// Third value returned (bool) indicates if last page has been reached.
func (c *Client) GetTips(page int) ([]*Tip, error, bool) {

	pageStr := strconv.Itoa(page)

	resp := &GetTipsResponse{}
	err := c.apiget("/partners/tips?perPage=10&page="+pageStr, resp)
	if err != nil {
		return nil, err, false
	}

	return resp.Tips, nil, resp.Pager.NbPages == resp.Pager.CurrentPage
}
