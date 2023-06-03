package models

type PodItem struct {
	Name     string `json:"service_name"`
	Image    string `json:"Image_name"`
	Tag      string `json:"Image_tag"`
	Network  string `json:"network_name"`
	LabelDir string `json:"vritual_path"`
}
