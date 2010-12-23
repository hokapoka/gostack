package gostack

import(
)

type QuestionsResult struct {
	Total int64
	Page int
	Pagesize int
	Questions []Question
}

type Question struct{
	Tags []string
	Answer_count int
	Favorite_count int
	Question_timeline_url string
	Question_comments_url string
	Question_answers_url string
	Question_id int64
	Owner Owner
	Creation_date *jTime
	Last_activity_date *jTime
	Up_vote_count int
	Down_vote_count int
	Vote_count int
	Score int
	Community_owned bool
	Title string
	Answers []Answer
}

type Owner struct{
	User_id int64
	User_type string
	Display_name string
	Reputation int
	Email_hash string
}

type Answer struct{
	Answer_id int64
	Accepted bool
	Answer_comments_url string
	Question_id int64
	Owner Owner
	Creation_date *jTime
	Last_activity_date *jTime
	Up_vote_count int
	Down_vote_count int
	Vote_count int
	Score int
	Community_owned bool
	Title string

}


