package models

import (
	"database/sql"
)

/*Target The target type*/
type Target struct {
	ID              string
	Domain          string
	OriginIPAddress string
	CertKey         sql.NullString
	CertCrt         sql.NullString
	AutoCert        bool
	UseHTTPS        bool
	WAFEnabled      bool
	Proto           int
}
