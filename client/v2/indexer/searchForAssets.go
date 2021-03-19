package indexer

import (
	"context"

	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

type SearchForAssetsParams struct {

	// AssetID asset ID
	AssetID uint64 `url:"asset-id,omitempty"`

	// Creator filter just assets with the given creator address.
	Creator string `url:"creator,omitempty"`

	// IncludeAll include all items including closed accounts, deleted applications,
	// destroyed assets, opted-out asset holdings, and closed-out application
	// localstates.
	IncludeAll bool `url:"include-all,omitempty"`

	// Limit maximum number of results to return.
	Limit uint64 `url:"limit,omitempty"`

	// Name filter just assets with the given name.
	Name string `url:"name,omitempty"`

	// Next the next page of results. Use the next token provided by the previous
	// results.
	Next string `url:"next,omitempty"`

	// Unit filter just assets with the given unit.
	Unit string `url:"unit,omitempty"`
}

type SearchForAssets struct {
	c *Client

	p SearchForAssetsParams
}

// AssetID asset ID
func (s *SearchForAssets) AssetID(AssetID uint64) *SearchForAssets {
	s.p.AssetID = AssetID
	return s
}

// Creator filter just assets with the given creator address.
func (s *SearchForAssets) Creator(Creator string) *SearchForAssets {
	s.p.Creator = Creator
	return s
}

// IncludeAll include all items including closed accounts, deleted applications,
// destroyed assets, opted-out asset holdings, and closed-out application
// localstates.
func (s *SearchForAssets) IncludeAll(IncludeAll bool) *SearchForAssets {
	s.p.IncludeAll = IncludeAll
	return s
}

// Limit maximum number of results to return.
func (s *SearchForAssets) Limit(Limit uint64) *SearchForAssets {
	s.p.Limit = Limit
	return s
}

// Name filter just assets with the given name.
func (s *SearchForAssets) Name(Name string) *SearchForAssets {
	s.p.Name = Name
	return s
}

// Next the next page of results. Use the next token provided by the previous
// results.
func (s *SearchForAssets) Next(Next string) *SearchForAssets {
	s.p.Next = Next
	return s
}

// Unit filter just assets with the given unit.
func (s *SearchForAssets) Unit(Unit string) *SearchForAssets {
	s.p.Unit = Unit
	return s
}

func (s *SearchForAssets) Do(ctx context.Context, headers ...*common.Header) (response models.AssetsResponse, err error) {
	err = s.c.get(ctx, &response, "/v2/assets", s.p, headers)
	return
}
