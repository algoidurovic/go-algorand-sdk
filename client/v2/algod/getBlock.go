package algod

import (
	"context"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/types"
)

type BlockParams struct {

	// Format configures whether the response object is JSON or MessagePack encoded.
	Format string `url:"format,omitempty"`
}

type Block struct {
	c *Client

	round uint64

	p BlockParams
}

func (s *Block) Do(ctx context.Context, headers ...*common.Header) (result types.Block, err error) {
	var response models.BlockResponse

	s.p.Format = "msgpack"
	err = s.c.getMsgpack(ctx, &response, fmt.Sprintf("/v2/blocks/%d", s.round), s.p, headers)
	if err != nil {
		return
	}

	result = response.Block
	return
}
