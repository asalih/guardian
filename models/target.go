package models

/*Target The target type*/
type Target struct {
	ID              string
	Domain          string
	OriginIPAddress string
	CertKey         string
	CertCrt         string
	UseHTTPS        bool
	WAFEnabled      bool
}

/*FirewallRule The firewall rule*/
type FirewallRule struct {
	ID         string
	Expression string
}
