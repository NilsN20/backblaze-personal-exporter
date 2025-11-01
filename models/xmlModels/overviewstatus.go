package xmlModels

type Overviewstatus struct {
	TransmitState TransmitState `xml:"bztransmit"`
}

type TransmitState struct {
	CurState string `xml:"cur_state,attr"`
}
