package models

type PodItem struct {
	Name     string `json:"service name"`
	Image    string `json:"Image name"`
	Tag      string `json:"Image tag"`
	Network  string `json:"network name"`
	LabelDir string `json:"vritual path"`
}
