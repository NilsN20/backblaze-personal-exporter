package xmlModels

type Totalbackup struct {
	Totals    Totals          `xml:"totals"`
	BZVolumes []BZVolumeTotal `xml:"bzvolume"`
}
type Totals struct {
	TotalNumFiles string `xml:"totnumfilesforbackup,attr"`
	TotalNumBytes string `xml:"totnumbytesforbackup,attr"`
}

type BZVolumeTotal struct {
	BZVolumeGuid        string `xml:"bzVolumeGuid,attr"`
	PervolTotalNumFiles string `xml:"pervol_sel_for_backup_numfiles,attr"`
	PervolTotalNumBytes string `xml:"pervol_sel_for_backup_numbytes,attr"`
}
