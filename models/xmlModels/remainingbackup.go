package xmlModels

type Remainingbackup struct {
	Remaining Remaining  `xml:"remaining"`
	BZVolumes []BZVolume `xml:"bzvolume"`
}
type Remaining struct {
	RemainingNumFilesForBackup string `xml:"remainingnumfilesforbackup,attr"`
	RemainingNumBytesForBackup string `xml:"remainingnumbytesforbackup,attr"`
}

type BZVolume struct {
	BZVolumeGuid                 string `xml:"bzVolumeGuid,attr"`
	PervolRemainingFilesNumFiles string `xml:"pervol_remaining_files_numfiles,attr"`
	PervolRemainingFilesNumBytes string `xml:"pervol_remaining_files_numbytes,attr"`
}
