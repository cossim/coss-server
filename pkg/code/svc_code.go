package code

var (
	// 用户服务状态码定义
	UserErrNotExistOrPassword      = New(10000, "用户不存在或密码错误")
	UserErrLoginFailed             = New(10001, "登录失败，请重试")
	UserErrLocked                  = New(10002, "用户已被锁定")
	UserErrDeleted                 = New(10003, "用户已被删除")
	UserErrDisabled                = New(10004, "用户已被禁用")
	UserErrStatusException         = New(10005, "用户状态异常")
	UserErrEmailAlreadyRegistered  = New(10006, "邮箱已被注册")
	UserErrRegistrationFailed      = New(10007, "注册失败，请重试")
	UserErrGetUserInfoFailed       = New(10008, "获取用户信息失败，请重试")
	UserErrUnableToGetUserListInfo = New(10009, "无法获取用户列表信息")
	UserErrPublicKeyNotExist       = New(10010, "用户公钥不存在")
	UserErrGetUserPublicKeyFailed  = New(10011, "获取用户公钥失败，请重试")
	UserErrSaveUserPublicKeyFailed = New(10012, "保存用户公钥失败，请重试")

	// 文件存储服务状态码定义
	StorageErrParseFilePathFailed    = New(11000, "解析文件路径失败")
	StorageErrCreateFileRecordFailed = New(11001, "保存文件失败")
	StorageErrGetFileInfoFailed      = New(11002, "获取文件信息失败")
	StorageErrDeleteFileFailed       = New(11003, "删除文件失败")

	// 关系服务状态码定义
	RelationErrUserNotFound                = New(13000, "用户不存在")
	RelationErrFriendNotFound              = New(13001, "好友不存在")
	RelationErrRelationNotFound            = New(13002, "关系不存在")
	RelationErrAddFriendFailed             = New(13003, "添加好友失败")
	RelationErrFriendRequestAlreadyPending = New(13004, "好友请求已经发送，等待确认")
	RelationErrAlreadyFriends              = New(13005, "已经是好友")
	RelationErrConfirmFriendFailed         = New(13006, "确认好友失败")
	RelationErrDeleteFriendFailed          = New(13007, "删除好友失败")
	RelationErrAddBlacklistFailed          = New(13008, "添加黑名单失败")
	RelationErrDeleteBlacklistFailed       = New(13009, "删除黑名单失败")
	RelationErrGetFriendListFailed         = New(13010, "获取好友列表失败")
	RelationErrGetBlacklistFailed          = New(13011, "获取黑名单列表失败")
	RelationErrGetUserRelationFailed       = New(13012, "获取用户关系失败")
	RelationErrInsertUserGroupFailed       = New(13013, "创建群聊失败")
	RelationErrGetUserGroupIDsFailed       = New(13014, "获取群聊列表失败")

	// 消息服务错误码定义
	MsgErrInsertUserMessageFailed       = New(14000, "发送消息失败")
	MsgErrInsertGroupMessageFailed      = New(14001, "发送群聊消息失败")
	MsgErrGetUserMessageListFailed      = New(14002, "获取用户消息列表失败")
	MsgErrGetLastMsgsForUserWithFriends = New(14003, "获取消息失败")
	MsgErrGetLastMsgsForGroupsWithIDs   = New(14004, "获取群聊消息失败")
	MsgErrGetLastMsgsByDialogIds        = New(14005, "获取消息失败")
	DialogErrCreateDialogFailed         = New(14100, "创建对话失败")
	DialogErrJoinDialogFailed           = New(14101, "加入对话失败")
	DialogErrGetUserDialogListFailed    = New(14102, "获取用户对话列表失败")

	// 群组服务错误码定义
	GroupErrGetGroupInfoByGidFailed      = New(15000, "获取群聊信息失败")
	GroupErrGetBatchGroupInfoByIDsFailed = New(15001, "获取群聊列表信息失败")
	GroupErrUpdateGroupFailed            = New(15002, "更新群聊信息失败")
	GroupErrInsertGroupFailed            = New(15003, "创建群组失败")
	GroupErrDeleteGroupFailed            = New(15004, "删除群聊失败")
)
