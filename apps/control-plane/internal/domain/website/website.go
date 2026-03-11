package website

type Website struct {
	ID       string `json:"id"`
	Domain   string `json:"domain"`
	Upstream string `json:"upstream"`
}
