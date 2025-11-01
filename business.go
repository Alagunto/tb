package tb

import (
	"encoding/json"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
)

// BusinessConnection returns the information about the connection of the bot with a business account.
func (b *Bot[RequestType]) BusinessConnection(id string) (*telegram.BusinessConnection, error) {
	params := map[string]string{
		"business_connection_id": id,
	}

	data, err := b.Raw( "getBusinessConnection", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result *telegram.BusinessConnection
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}
	return resp.Result, nil
}
