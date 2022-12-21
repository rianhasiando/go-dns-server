package lookup

type DNSDatabase struct {
	// map[domain name]map[class]data
	// ex. map[rianhasiando.space.]map[IN]data
	ListZones map[string]map[string]DNSClassData `json:"list_zones"`
}

type DNSClassData struct {
	A  []RecordA
	NS []RecordNS
}

type RecordA struct {
	Name  string `json:"name"`
	TTL   int    `json:"ttl"`
	Value string `json:"value"`
}

type RecordNS struct {
	TTL   int    `json:"ttl"`
	Value string `json:"value"`
}
