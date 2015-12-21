package golegram

type UpdateHandler func(Update)

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
	Reply_to_message *Message
	Text 			string
	Audio			Audio
	Document		Document
	Photo			[]PhotoSize
	Sticker			Sticker
	Video			Video
	Voice			Voice
	Caption			string
	Contact			Contact
	Location		Location
	New_chat_participant User
	Left_chat_participant User
	New_chat_title	string
	New_chat_photo	[]PhotoSize
	Delete_chat_photo bool
	Group_chat_created bool
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

type PhotoSize struct{
	File_id			string
	Width			int32
	Height			int32
	File_size		int32
}

type Audio struct{
	File_id			string
	Duration		int32
	Performer		string
	Title			string
	Mime_type		string
	File_size		int32
}

type Document struct{
	File_id			string
	Thumb			PhotoSize
	File_name		string
	Mime_type		string
	File_size		int32
}

type Sticker struct{
	File_id			string
	Width			int32
	Height			int32
	Thumb			PhotoSize
	File_size		int32
}

type Video struct{
	File_id			string
	Width			int32
	Height			int32
	Duration		int32
	Thumb			PhotoSize
	Mime_type		string
	File_size		int32
}

type Voice struct{
	File_id			string
	Duration		int32
	Mime_type		string
	File_size		int32
}

type Contact struct{
	Phone_number	string
	First_name		string
	Last_name		string
	User_id			int32
}

type Location struct{
	Longitude		float64
	Latitude		float64
}

type UserProfilePhotos struct{
	Total_count		int32
	Photos			[][]PhotoSize
}

type File struct{
	File_id			string
	File_size		int32
	File_path		string
}