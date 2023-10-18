package models

type Feed struct{
	posts []FeedPost
}

type FeedPost struct{
	//bulletin service data structure
	Content string
}