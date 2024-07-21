package aws

type Signer interface {
	SignURL(url string) (string, error)
}
