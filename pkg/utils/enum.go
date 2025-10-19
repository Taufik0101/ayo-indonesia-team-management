package utils

type UserType string

const (
	ADMIN UserType = "admin"
	USER  UserType = "user"
)

type PlayerPositionType string

const (
	PENYERANG      PlayerPositionType = "penyerang"
	GELANDANG      PlayerPositionType = "gelandang"
	BERTAHAN       PlayerPositionType = "bertahan"
	PENJAGA_GAWANG PlayerPositionType = "penjaga_gawang"
)
