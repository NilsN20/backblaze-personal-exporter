package xmlModels

type Lastbackupcompleted struct {
	LastBackup lastbackupcompleted `xml:"lastbackupcompleted"`
}

type lastbackupcompleted struct {
	GmtMillis     string `xml:"gmt_millis,attr"`
	Localdatetime string `xml:"localdatetime,attr"`
}
