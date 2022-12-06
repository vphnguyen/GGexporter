// ConstsAndVariables
//   - Định nghĩa các hằng số và struct dùng để chứa cấu hình cho GGexporter.
package model

var RootURL = "http://gg-svmgr.io"

// Struct dùng để chứa cấu hình cho GGexporter.
type Config struct {
	Bind    string
	MgrHost string
	MgrPort string
	RootURL string
}

// Chuyển đổi các hằng số này sẽ ứng với các giá trị trong xml.
const (
	TYPE_PMSRVR    = "14"
	TYPE_EXTRACT   = "2"
	TYPE_REPLICAT  = "3"
	TYPE_PUMP      = "4"
	TYPE_MGR       = "1"
	STATUS_RUNNING = "3"
)
