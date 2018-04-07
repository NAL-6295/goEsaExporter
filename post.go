package main

//Post ポスト
type Post struct {
	Number         int      `json:"number"`
	Name           string   `json:"name"`
	FullName       string   `json:"full_name"`
	Wip            bool     `json:"wip"`
	BodyMd         string   `json:"body_md"`
	BodyHTML       string   `json:"body_html"`
	Message        string   `json:"message"`
	URL            string   `json:"url"`
	Tags           []string `json:"tags"`
	RevisionNumber int      `json:"revision_number"`
}

//Posts Posts
type Posts struct {
	Posts      []Post `json:"posts"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	MaxPerPage int    `json:"max_per_page"`
	NextPage   int    `json:"next_page"`
}
