package model

type AbebooksSearchResults struct {
	ResultCount string `xml:"resultCount"`
	Books []AbebooksSearchBook `xml:"Book"`
}

type AbebooksSearchBook struct {
	BookId 					string `xml:"bookId"`
	Isbn10 					string `xml:"isbn10"`
	Isbn13 					string `xml:"isbn13"`
	ListingCondition 		string `xml:"listingCondition"`
	ItemCondition 			string `xml:"itemCondition"`
	Quantity 				string `xml:"quantity"`
	VendorCurrency 			string `xml:"vendorCurrency"`
	ListingPrice 			string `xml:"listingPrice"`
	FirstBookShipCost		string `xml:"firstBookShipCost"`
	ExtraBookShipCost		string `xml:"extraBookShipCost"`
	MinShipDays				string `xml:"minShipDays"`
	MaxShipDays				string `xml:"maxShipDays"`
	TotalListingPrice		string `xml:"totalListingPrice"`
	ListingUrl				string `xml:"listingUrl"`
	Author					string `xml:"author"`
	Title					string `xml:"title"`
	PublisherName			string `xml:"publisherName"`
	VendorImage				string `xml:"vendorImage"`
	CatalogImage			string `xml:"catalogImage"`
	VendorName				string `xml:"vendorName"`
	VendorLocation			string `xml:"vendorLocation"`
	VendorId				string `xml:"vendorId"`
	SellerRating			string `xml:"sellerRating"`
	Keywords				string `xml:"keywords"`
	BookJacket				string `xml:"bookJacket"`
	BindingType				string `xml:"bindingType"`
	EditionType				string `xml:"editionType"`
	Signed					string `xml:"signed"`
	PublicationYear			string `xml:"publicationYear"`
	VendorPrice				string `xml:"vendorPrice"`
	VendorDescription		string `xml:"vendorDescription"`
	ProductType				string `xml:"productType"`
}