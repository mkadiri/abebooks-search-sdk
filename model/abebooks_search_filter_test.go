package model

import (
	"testing"
)

func TestGetParams(t *testing.T) {
	filter := AbebooksSearchFilter{
		BookCondition: "new",
		MinSellerRating: 0,
		OutputSize: "short",
		VendorId: "",
	}

	values := filter.GetParams()

	if values.Get("maxprice") != "" {
		t.Errorf("MaxPrice shouldn't be set")
	}

	if values.Get("minsellerrating") != "" {
		t.Errorf("MinSellerRating shouldn't be set")
	}

	if values.Get("bookcondition") != "new" {
		t.Errorf("BookCondition was not set")
	}

	if values.Get("vendorid") != "" {
		t.Errorf("VendorId shouldn't be set")
	}
}