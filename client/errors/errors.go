package broadcaster_client_errors

import "errors"

type Error error

var (
	BroadcastRequestSignatureEmpty       Error = errors.New("signature can not be empty")
)
