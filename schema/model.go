package schema

type UserBasic struct {
	Account   string `binding:"required" json:"account" bson:"account"`
	Password  string `binding:"required" json:"password" bson:"password"`
	Avatar    string `binding:"required" json:"avatar" bson:"avatar"`
	CreatedAt string `binding:"required" json:"createdAt" bson:"created_at"`
	UpdatedAt string `binding:"required" json:"updatedAt" bson:"updated_at"`
	Email     string `binding:"required" json:"email" bson:"email"`
	Nickname  string `binding:"required" json:"nickname" bson:"nickname"`
	Gender    string `binding:"required" json:"gender" bson:"gender"`
}

type Message struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}

type MessageBasic struct {
	UserIdentity string `json:"user_identity" bson:"user_identity"`
	RoomIdentity string `json:"room_identity" bson:"room_identity"`
	Data         string `json:"data" bson:"data"`
	CreatedAt    int64  `json:"created_at" bson:"created_at"`
	UpdatedAt    int64  `json:"emial" bson:"emial"`
}

func (UserBasic) Collection() string {
	return "user_basic"
}

func (MessageBasic) Collection() string {
	return "message_basic"
}
