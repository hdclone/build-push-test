package model

import "time"

func PtrTime(t time.Time) *time.Time { return &t }
func PtrString(s string) *string     { return &s }
