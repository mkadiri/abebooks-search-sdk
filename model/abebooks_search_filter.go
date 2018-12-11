package model

import "net/url"

type AbebooksSearchFilter struct {
	BookCondition string
	MinSellerRating int
	OutputSize string
	VendorId string
	MaxPrice int
}

const OutputSizeLong = "long"
const OutputSizeShot = "short"

const ConditionNew = "newonly"
const ConditionUsed = "used"


func (filter AbebooksSearchFilter) GetParams() url.Values{
	values := url.Values{}

	if filter.BookCondition != "" {
		values.Set("bookcondition", filter.BookCondition)
	}

	if filter.MinSellerRating != 0 {
		values.Set("minsellerrating", string(filter.MinSellerRating))
	}

	if filter.OutputSize != "" {
		values.Set("outputsize", filter.OutputSize)
	}

	if filter.VendorId != "" {
		values.Set("vendorid", filter.VendorId)
	}

	if filter.MaxPrice != 0 {
		values.Set("maxprice", string(filter.MaxPrice))
	}

	return values
}