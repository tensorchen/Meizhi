package main

type GankApiToday struct {
	Category []string `json:"category"`
	Error    bool     `json:"error"`
	Results  struct {
		Fuli []Fuli `json:"福利"`
	}
}

type Fuli struct {
	Url string `json:"url"`
}

type Result struct {
	Code    int
	Message string
	Data    *DataEntity
}

type DataEntity struct {
	Image string
}
