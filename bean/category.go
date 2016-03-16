package bean

import ()

// Attributes:
//  - ID
//  - Name
//  - Desc
//  - Subcategories
type Category struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Desc          string        `json:"desc"`
	Subcategories []SubCategory `json:"subcategories"`
}

// Attributes:
//  - ID
//  - Name
//  - Desc
type SubCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
