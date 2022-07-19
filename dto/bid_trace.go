package dto

type message struct {
	Slot                 string `json:"slot"`
	ParentHash           string `json:"parent_hash"`
	BlockHash            string `json:"block_hash"`
	BuilderPublicKey     string `json:"builder_pubkey"`
	ProposerPublicKey    string `json:"proposer_pubkey"`
	ProposerFeeRecipient string `json:"proposer_fee_recipient"`
	Value                string `json:"value"`
}

type BidTrace struct {
	Signature string  `json:"signature"`
	Message   message `json:"message"`
}
