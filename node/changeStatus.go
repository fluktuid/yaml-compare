package node

type ChangeStatus byte

const (
	UNDEFINED ChangeStatus = ' '
	REMOVED   ChangeStatus = '-'
	ADDED     ChangeStatus = '+'
	CHANGED   ChangeStatus = '~'
)
