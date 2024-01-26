package code

var (
	// 用户服务状态码定义
	UserErrNotExistOrPassword        = New(10000, "用户不存在或密码错误")
	UserErrLoginFailed               = New(10001, "登录失败，请重试")
	UserErrLocked                    = New(10002, "用户已被锁定")
	UserErrDeleted                   = New(10003, "用户已被删除")
	UserErrDisabled                  = New(10004, "用户已被禁用")
	UserErrStatusException           = New(10005, "用户状态异常")
	UserErrEmailAlreadyRegistered    = New(10006, "邮箱已被注册")
	UserErrRegistrationFailed        = New(10007, "注册失败，请重试")
	UserErrGetUserInfoFailed         = New(10008, "获取用户信息失败，请重试")
	UserErrUnableToGetUserListInfo   = New(10009, "无法获取用户列表信息")
	UserErrPublicKeyNotExist         = New(10010, "用户公钥不存在")
	UserErrGetUserPublicKeyFailed    = New(10011, "获取用户公钥失败，请重试")
	UserErrSaveUserPublicKeyFailed   = New(10012, "保存用户公钥失败，请重试")
	UserErrNotExist                  = New(10013, "用户不存在")
	UserErrGetUserSecretBundleFailed = New(10014, "获取用户秘钥包失败，请重试")
	UserErrSetUserSecretBundleFailed = New(10015, "设置用户秘钥包失败，请重试")
	UserErrErrLogoutFailed           = New(10016, "退出登录失败，请重试")
	UserErrOldPassword               = New(10017, "旧密码错误")
	UserErrSwapPublicKeyFailed       = New(10018, "用户交换公钥失败")
	UserErrGetUserLoginClientsFailed = New(10019, "获取用户登录客户端失败，请重试")

	// 文件存储服务状态码定义
	StorageErrParseFilePathFailed    = New(11000, "解析文件路径失败")
	StorageErrCreateFileRecordFailed = New(11001, "保存文件失败")
	StorageErrGetFileInfoFailed      = New(11002, "获取文件信息失败")
	StorageErrDeleteFileFailed       = New(11003, "删除文件失败")

	// 关系服务状态码定义
	RelationErrUserNotFound                          = New(13000, "用户不存在")
	RelationUserErrFriendRelationNotFound            = New(13001, "好友关系不存在")
	RelationErrRelationNotFound                      = New(13002, "关系不存在")
	RelationErrAddFriendFailed                       = New(13003, "添加好友失败")
	RelationErrFriendRequestAlreadyPending           = New(13004, "好友请求已经发送，等待确认")
	RelationErrAlreadyFriends                        = New(13005, "已经是好友")
	RelationErrConfirmFriendFailed                   = New(13006, "管理好友申请失败")
	RelationErrDeleteFriendFailed                    = New(13007, "删除好友失败")
	RelationErrAddBlacklistFailed                    = New(13008, "添加黑名单失败")
	RelationErrDeleteBlacklistFailed                 = New(13009, "删除黑名单失败")
	RelationErrGetFriendListFailed                   = New(13010, "获取好友列表失败")
	RelationErrGetBlacklistFailed                    = New(13011, "获取黑名单列表失败")
	RelationErrGetUserRelationFailed                 = New(13012, "获取用户关系失败")
	RelationUserErrGetRequestListFailed              = New(13013, "获取用户申请列表失败")
	RelationUserErrNoFriendRequestRecords            = New(13014, "未找到好友申请记录")
	RelationErrRejectFriendFailed                    = New(13015, "拒绝好友申请失败")
	RelationErrSetUserFriendSilentNotificationFailed = New(13016, "设置用户好友静默通知失败")
	RelationErrNotInBlacklist                        = New(13017, "没有在黑名单")
	RelationErrSendFriendRequestFailed               = New(13018, "发送好友请求失败")
	RelationErrManageFriendRequestFailed             = New(13019, "管理好友请求失败")

	RelationErrCreateGroupFailed                         = New(13101, "创建群聊失败")
	RelationErrGetGroupIDsFailed                         = New(13102, "获取群聊成员")
	RelationGroupErrRequestFailed                        = New(13103, "申请加入群聊失败")
	RelationGroupErrRequestAlreadyPending                = New(13104, "等待同意申请")
	RelationGroupErrAlreadyInGroup                       = New(13105, "已经在群聊中")
	RelationGroupErrApproveJoinGroupFailed               = New(13106, "同意加入群聊失败")
	RelationGroupErrNoJoinRequestRecords                 = New(13107, "没有申请加入群聊记录")
	RelationGroupErrRejectJoinGroup                      = New(13108, "拒绝加入群聊")
	RelationGroupErrRemoveUserFromGroupFailed            = New(13109, "将用户移除群聊失败")
	RelationGroupErrLeaveGroupFailed                     = New(13110, "退出群聊失败")
	RelationGroupErrGetJoinRequestListFailed             = New(13111, "获取群聊申请列表失败")
	RelationGroupErrGroupRelationFailed                  = New(13112, "获取用户群组关系失败")
	RelationGroupErrGetGroupInfoFailed                   = New(13113, "获取群聊信息失败")
	RelationGroupErrManageJoinFailed                     = New(13114, "管理群聊申请失败")
	RelationGroupErrInviteFailed                         = New(13115, "邀请入群失败")
	RelationGroupErrRelationNotFound                     = New(13116, "关系不存在")
	RelationGroupErrSetUserGroupSilentNotificationFailed = New(13117, "设置群聊消息静默通知失败")
	RelationGroupErrNotInGroup                           = New(13118, "没有在群聊中")

	// 消息服务错误码定义
	MsgErrInsertUserMessageFailed                   = New(14000, "发送消息失败")
	MsgErrInsertGroupMessageFailed                  = New(14001, "发送群聊消息失败")
	MsgErrGetUserMessageListFailed                  = New(14002, "获取用户消息列表失败")
	MsgErrGetLastMsgsForUserWithFriends             = New(14003, "获取消息失败")
	MsgErrGetLastMsgsForGroupsWithIDs               = New(14004, "获取群聊消息失败")
	MsgErrGetLastMsgsByDialogIds                    = New(14005, "获取消息失败")
	DialogErrCreateDialogFailed                     = New(14100, "创建对话失败")
	DialogErrJoinDialogFailed                       = New(14101, "加入对话失败")
	DialogErrGetUserDialogListFailed                = New(14102, "获取用户对话列表失败")
	DialogErrDeleteDialogFailed                     = New(14103, "删除对话失败")
	DialogErrDeleteDialogUsersFailed                = New(14104, "删除对话用户失败")
	DialogErrGetDialogUserByDialogIDAndUserIDFailed = New(14105, "获取对话用户失败")
	MsgErrEditUserMessageFailed                     = New(14006, "编辑用户消息失败")
	MsgErrEditGroupMessageFailed                    = New(14007, "编辑群聊消息失败")
	MsgErrDeleteUserMessageFailed                   = New(14008, "撤回用户消息失败")
	MsgErrDeleteGroupMessageFailed                  = New(14009, "撤回群聊消息失败")
	GetMsgErrGetGroupMsgByIDFailed                  = New(14010, "获取群聊消息失败")
	GetMsgErrGetUserMsgByIDFailed                   = New(14011, "获取用户消息失败")
	SetMsgErrSetUserMsgLabelFailed                  = New(14012, "设置用户消息标注失败")
	SetMsgErrSetGroupMsgLabelFailed                 = New(14013, "设置群聊消息标注失败")
	GetMsgErrGetUserMsgLabelByDialogIdFailed        = New(14014, "获取用户消息标注失败")
	GetMsgErrGetGroupMsgLabelByDialogIdFailed       = New(14015, "获取群聊消息标注失败")
	SetMsgErrSetUserMsgsReadStatusFailed            = New(14016, "批量已读消息失败")
	SetMsgErrSetUserMsgReadStatusFailed             = New(14017, "修改消息已读状态失败")
	GetMsgErrGetUnreadUserMsgsFailed                = New(14018, "获取未读消息失败")
	MsgErrTimeoutExceededCannotRevoke               = New(14019, "超过时间限制不能撤回")

	// 群组服务错误码定义
	GroupErrGetGroupInfoByGidFailed             = New(15000, "获取群聊信息失败")
	GroupErrGetBatchGroupInfoByIDsFailed        = New(15001, "获取群聊列表信息失败")
	GroupErrUpdateGroupFailed                   = New(15002, "更新群聊信息失败")
	GroupErrInsertGroupFailed                   = New(15003, "创建群组失败")
	GroupErrDeleteGroupFailed                   = New(15004, "删除群聊失败")
	GroupErrGroupNotFound                       = New(15005, "群聊不存在")
	GroupErrDeleteUserGroupRelationFailed       = New(15006, "删除用户群聊关系失败")
	GroupErrDeleteUserGroupRelationRevertFailed = New(15007, "删除用户群聊关系回滚失败")
	GroupErrGroupStatusNotAvailable             = New(15008, "群聊状态不可用")
	GroupErrUserIsMuted                         = New(15009, "用户禁言中")
)
