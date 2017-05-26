package hostedzones

type HostedZones struct {
	API     string `json:"api" yaml:"api"`
	Etcd    string `json:"etcd" yaml:"etcd"`
	Ingress string `json:"ingress" yaml:"ingress"`
}
