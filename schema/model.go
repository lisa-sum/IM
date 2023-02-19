package schema

type UserFormBasic struct {
	Account   string `bson:"account" form:"account"`
	Password  string `bson:"password" form:"password"`
	Email     string `bson:"email" form:"email"`
	Nickname  string `bson:"nickname" form:"nickname"`
	Gender    string `bson:"gender" form:"gender"`
	CreatedAt int    `bson:"created_at" form:"createdAt"`
	UpdatedAt int    `bson:"updated_at" form:"updatedAt"`
}

type UserBasic struct {
	Account   string `json:"account" bson:"account"`
	Password  string `json:"password" bson:"password"`
	Avatar    string `json:"avatar" bson:"avatar"`
	CreatedAt int64  `json:"createdAt" bson:"created_at"`
	UpdatedAt int64  `json:"updatedAt" bson:"updated_at"`
	Email     string `json:"email" bson:"email"`
	Nickname  string `json:"nickname" bson:"nickname"`
	Gender    string `json:"gender" bson:"gender"`
}

type Message struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"roomIdentity"`
}

type RoomBasic struct {
	UserIdentity string `json:"userIdentity" bson:"user_identity"` // 创建者/房主/群主
	Avatar       string `json:"avatar" bson:"avatar"`              //房间头像/群头像
	Info         string `json:"info" bson:"info"`                  // 房间简介/群介绍
	Name         string `json:"name" bson:"name"`                  // 房间名/群名
	Number       string `json:"number" bson:"number"`              // 房间号/群号
	CreatedAt    int64  `json:"createdAt" bson:"created_at"`       // 创建时间
	UpdatedAt    int64  `json:"updatedAt" bson:"updated_at"`       // 更新时间
}

type MessageBasic struct {
	UserIdentity string `json:"userIdentity" bson:"user_identity"`
	RoomIdentity string `json:"roomIdentity" bson:"room_identity"`
	Data         string `json:"data" bson:"data"`
	CreatedAt    int64  `json:"createdAt" bson:"created_at"`
	UpdatedAt    int64  `json:"updatedAt" bson:"updated_at"`
}

func (UserBasic) Collection() string {
	return "user_basic"
}

func (MessageBasic) Collection() string {
	return "message_basic"
}

func (RoomBasic) Collection() string {
	return "room_basic"
}
