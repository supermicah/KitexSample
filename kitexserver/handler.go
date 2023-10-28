package main

import (
	"context"
	"fmt"

	"kitexserver/kitex_gen/dragon"
)

// SampleImpl implements the last service interface defined in the IDL.
type SampleImpl struct{}

// DragonSay implements the SampleImpl interface.
func (s *SampleImpl) DragonSay(ctx context.Context, req *dragon.DragonSayRequest) (resp *dragon.DragonSayResponse, err error) {
	if len(req.GetMessage()) == 0 {
		return nil, fmt.Errorf("error")
	}
	resp = &dragon.DragonSayResponse{Message: "i'm call back."}
	return
}
