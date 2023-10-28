package main

import (
	"context"

	"kitexserver/kitex_gen/dragon"
)

// SampleImpl implements the last service interface defined in the IDL.
type SampleImpl struct{}

// DragonSay implements the SampleImpl interface.
func (s *SampleImpl) DragonSay(ctx context.Context, req *dragon.DragonSayRequest) (resp *dragon.DragonSayResponse, err error) {
	return
}
