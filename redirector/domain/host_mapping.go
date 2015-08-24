package domain

// HostMapRepository is the interface that will be used in our
// usecases to define database requests
type HostMapRepository interface {
	Create(*HostMap) error
	Read(host string) (*HostMap, error)
	Update(updatedHostMap *HostMap) error
	Delete(toBeDeletedHostMap *HostMap) error
}

// HostMap is a simple structure that maps the URL we want to
// take the request From and where it should be redirected to
type HostMap struct {
	From string `json:"From"`
	To   string `json:"To"`
}
