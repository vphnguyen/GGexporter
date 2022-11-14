package model

var RootURL = "http://gg-svmgr.io"

type Config struct {
	Bind    string
	MgrHost string
	MgrPort string
	RootURL string
}

const (
	TYPE_PMSRVR    = "14"
	TYPE_EXTRACT   = "2"
	TYPE_REPLICAT  = "3"
	TYPE_PUMP      = "4"
	TYPE_MGR       = "1"
	STATUS_RUNNING = "3"
)
