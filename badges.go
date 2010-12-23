package gostack

import(
)

type BadgesResult struct {
	Badges []Badge
}

type Badge struct{
	Badge_id int
	Rank string
	Name string
	Description string
	Award_count int
	Tag_base bool
	Badges_recipients_url string
}


