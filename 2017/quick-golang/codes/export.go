// This is exported
func IPv6() (*JSONIP, error) {
	return fetch(ipv6Endpoint)
}

// This is not
func fetch(target string) (*JSONIP, error) {
	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ip *JSONIP
	err = json.NewDecoder(resp.Body).Decode(&ip)
	return ip, err
}
