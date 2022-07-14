package model

import "github.com/rs/xid"

type ID string

func NewID(kind string) ID {
	return ID(kind + "_" + xid.New().String())
}
