// Code generated by goctl. DO NOT EDIT.
package types

type RecommendReq struct {
}

type RecommendResply struct {
	Code    int64          `json:"Code"`
	Message string         `json:"Message"`
	Data    []*ProductItem `json:"Data"`
}

type ProductsReq struct {
	Page        int64  `json:"page"`
	PageSize    int64  `json:"pageSize"`
	ProductName string `json:"productName"`
}

type ProductsResply struct {
	Code       int64          `json:"Code"`
	Message    string         `json:"Message"`
	Page       int64          `json:"Page"`
	PageSize   int64          `json:"PageSize"`
	TotalCount int64          `json:"TotalCount"`
	Data       []*ProductItem `json:"Data"`
}

type ProductItem struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	CoverPic    string  `json:"coverPic"`
	Property    string  `json:"property"`
	MtPrice     float64 `json:"mtPrice"`
	DisPrice    float64 `json:"disPrice"`
	Stock       int64   `json:"stock"`
	State       int64   `json:"state"`
	SalesVolume int64   `json:"salesVolume"`
	Images      string  `json:"images"`
	Detail      string  `json:"detail"`
}

type DetailReq struct {
	PorductId int64 `json:"porductId"`
}

type DetailResply struct {
	Code    int64       `json:"Code"`
	Message string      `json:"Message"`
	Data    ProductItem `json:"Data"`
}

type CateListReq struct {
	Pid int64 `json:"Pid"`
}

type CategoryResply struct {
	Code    int64           `json:"Code"`
	Message string          `json:"Message"`
	Data    []*CategoryItem `json:"Data"`
}

type CategoryItem struct {
	ID       int64  `json:"Id"`
	ParentId int64  `json:"ParentId"`
	Name     string `json:"Name"`
	Status   int64  `json:"Status"`
}
