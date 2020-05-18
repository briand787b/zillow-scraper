package controller

type ctxKey int

const (
	propertyCtxKey ctxKey = iota
	skipCtxKey
	takeCtxKey
)
