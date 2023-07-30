package snapshot

import (
	"context"
)

type Voter interface {
	ActiveProposals(ctx context.Context, req *ActiveProposalsReq) ([]Proposal, error)
	Vote(ctx context.Context, req *VoteReq) error
}
