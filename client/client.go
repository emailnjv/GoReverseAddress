package client

import (
	"fmt"

	"GoReverseAddress/resources"
)

type Client struct {
	SiteModule SiteModule
}

type SiteModule interface {
	RequestRALU(resources.AddressAndName, error) (string, error)
	ParseRALU(string, error) (resources.PersonalInfo, error)
}

type Caller interface {
	Call(string) (string, error)
}

func (c Client)	ReverseAddressLookUp(addressNameInput resources.AddressAndName) (resources.PersonalInfo, error) {
	return resources.PersonalInfo{}, fmt.Errorf("Error")
}

