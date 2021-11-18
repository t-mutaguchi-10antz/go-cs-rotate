package model

import (
	"fmt"
)

type platform string

const (
	// PlatformGCP = platform("gcp")
	PlatformAWS = platform("aws")
)

func Platform(v string) (platform, error) {
	o := platform(v)
	switch o {
	case
		// PlatformGCP,
		PlatformAWS:
		return o, nil
	default:
		return platform(""), fmt.Errorf("failed to create platform value: invalid value")
	}
}
