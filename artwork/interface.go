package artwork

//Interface is interface to describe how to handle artworks
type Interface interface {
	CreateArtwork(artwork Artwork) (string, error)
	GetArtwork(id string) (Artwork, error)
	GetArtworkFromUserID(userID string) ([]Artwork, error)
	UpdateArtwork(artwork Artwork) error
}

//Status of an artwork. Active, removed, sold
type Status int8

// Description of Status
const (
	ACTIVE Status = iota
	REMOVED
	SOLD
)

// Artwork should be representation of model in database
type Artwork struct {
	id          string
	titel       string
	description string
	seller      string
	images      []string
	price       int32
	status      Status
}
