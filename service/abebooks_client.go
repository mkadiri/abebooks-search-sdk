package service

import (
	"net/http"
	"net/url"
)

type AbebooksClient struct {
	ClientKey string
	Currency string
	DestinationCountry string
	VendorLocation string
}

func (abebooksClient AbebooksClient) getBase() string {
	return "http://search2.abebooks.com/search"
}

func (abebooksClient AbebooksClient) getParams() url.Values {
	return url.Values{
		"clientkey": {abebooksClient.ClientKey},
		"currency": {abebooksClient.Currency},
		"destinationcountry": {abebooksClient.DestinationCountry},
		"vendorlocation": {abebooksClient.VendorLocation},
	}
}

func (abebooksClient AbebooksClient) Get(params url.Values) (resp *http.Response, err error) {
	requestUrl := abebooksClient.getBase() + "?" + abebooksClient.getParams().Encode() + "&" + params.Encode()

	req, err := http.NewRequest("GET", requestUrl, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err = client.Do(req)

	return resp, err
}