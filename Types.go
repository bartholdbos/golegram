package Golegram

type Update struct{
	Update_id		int32
	Message 		Message
}

type Bot struct{
	Token			string
	User			User
}

type Message struct{
	Message_id		int32
	From 			User
	Date 			int32
	Chat 			Chat
	Forward_from	User
	Forward_date	int32
	Text 			string
}

type User struct{
	Id				int32
	First_name		string
	Last_name		string
	Username		string
}

type Chat struct{
	Id				int32
	Type			string
	Title			string
	Username		string
	First_name		string
	Last_name		string
}