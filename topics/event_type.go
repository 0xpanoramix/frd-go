package topics

type EventType string

const (
	BuilderBidValid            EventType = "builder_bid_valid"
	ProposerGetHeader          EventType = "proposer_getheader"
	ProposerSubmitBlindedBlock EventType = "proposer_submit_blinded_block"
)
