package main

import (
	"testing"
	"errors"
)

func TestSwapper(t *testing.T) {
	liveHosts := map[string]struct{}{
		"abc-34": struct{}{},
		"abd-21": struct{}{},
	}
	spareHosts := map[string]struct{}{
		"abc-32": struct{}{},
		"abd-32": struct{}{},
	}
	s, err := SwapHost("abc-12", liveHosts, spareHosts)
	t.Error(err)
	t.Error(s)
}

func SwapHost(host string, liveHosts map[string]struct{}, spareHosts map[string]struct{}) (string, error) {
	liveRacks := make(map[string]struct{})
	for h := range liveHosts {
		rack := h[0:3]
		liveRacks[rack] = struct{}{}
	}

	for spare := range spareHosts {
		rack := spare[0:3]
		if _, present := liveRacks[rack]; !present {
			return spare, nil
		}

	}
	return "", errors.New("No spares exist")
}
