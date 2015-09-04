package goxtremio

import xms "github.com/emccode/goxtremio/api/v3"

type Initiator *xms.Initiator
type NewInitiatorOptions xms.PostInitiatorsReq
type NewInitiatorResult *xms.PostInitiatorsResp

//GetInitiator returns a specific initiator by name or ID
func (c *Client) GetInitiator(id string, name string) (Initiator, error) {
	i, err := c.api.GetInitiator(id, name)
	if err != nil {
		return nil, err
	}

	return i.Content, nil
}

//GetInitiators returns a list of initiators
func (c *Client) GetInitiators() (Refs, error) {
	i, err := c.api.GetInitiators()
	if err != nil {
		return nil, err
	}
	return i.Initiators, nil
}

//NewInitiator creates a volume
func (c *Client) NewInitiator(opts *NewInitiatorOptions) (NewInitiatorResult, error) {
	req := xms.PostInitiatorsReq(*opts)
	res, err := c.api.PostInitiators(&req)
	if err != nil {
		return nil, err
	}

	nir := NewInitiatorResult(res)
	return nir, nil
}

//DeleteInitiator deletes a volume
func (c *Client) DeleteInitiator(id string, name string) error {
	return c.api.DeleteInitiators(id, name)
}
