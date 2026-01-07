package did

type Document struct {
	ID        string   `json:"id"`
	PublicKey string   `json:"publicKey"`
	Services  []string `json:"services"`
}
