// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiToken represents a resolvable DeFi token instance.
type DefiToken struct {
	repo repository.Repository
	types.DefiToken
}

// NewDefiToken creates a new instance of resolvable DeFi token.
func NewDefiToken(tk *types.DefiToken, repo repository.Repository) *DefiToken {
	return &DefiToken{
		repo:      repo,
		DefiToken: *tk,
	}
}

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (rs *rootResolver) DefiTokens() ([]*DefiToken, error) {
	// pass the call to repository
	tkList, err := rs.repo.DefiTokens()
	if err != nil {
		return nil, err
	}

	// prep the container
	list := make([]*DefiToken, len(tkList))
	for ix, tk := range tkList {
		list[ix] = NewDefiToken(&tk, rs.repo)
	}

	return list, nil
}

// Price resolves the value of the token in ref. denomination
// using on-chain price oracle.
func (dt *DefiToken) Price() (hexutil.Big, error) {
	return dt.repo.DefiTokenPrice(&dt.Address)
}

// Price resolves the value of the token in ref. denomination
// using on-chain price oracle.
func (dt *DefiToken) AvailableBalance(args *struct{ Owner common.Address }) (hexutil.Big, error) {
	return dt.repo.Erc20Balance(&args.Owner, &dt.Address)
}

// DefiConfiguration resolves the current DeFi contract settings.
func (rs *rootResolver) DefiConfiguration() (*types.DefiSettings, error) {
	// pass the call to repository
	return rs.repo.DefiConfiguration()
}

// ErcTokenBalance resolves the current available balance of the specified token
// for the specified owner.
func (rs *rootResolver) ErcTokenBalance(args *struct {
	Owner common.Address
	Token common.Address
}) (hexutil.Big, error) {
	return rs.repo.Erc20Balance(&args.Owner, &args.Token)
}