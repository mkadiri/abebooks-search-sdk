package service

import (
	"github.com/mkadiri/abebooks-search-sdk/model"
	"encoding/xml"
	"io/ioutil"
)

func (abebooksClient AbebooksClient) SearchByIsbn(isbn string, filter model.AbebooksSearchFilter) (results model.AbebooksSearchResults, err error) {
	params := filter.GetParams()
	params.Set("isbn", isbn)

	resp, err := abebooksClient.Get(params)

	if err != nil {
		return results, err
	}

	defer resp.Body.Close()

	results = model.AbebooksSearchResults{}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return results, nil
	}

	err = xml.Unmarshal(body, &results)

	return results, err
}