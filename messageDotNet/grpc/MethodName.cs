
/// <summary>
/// MELAND service call method API name
/// </summary>
public static class ManagerServiceAction
{
    /// <summary>
    /// 获取服务时间戳
    /// </summary>
    public const string SERVICE_TIMESTAMP = "QueryServerTime";
    /// <summary>
    /// 注册服务
    /// </summary>
    public const string REGISTER_SERVICE = "ManagerActionServiceRegister";
    /// <summary>
    /// 查询服务信息
    /// </summary>
    public const string SELECT_SERVICE = "ManagerActionSelectService";

    public const string MULTI_SELECT_SERVICE = "ManagerActionMultiSelectService";

}

/// <summary>
///  main service call method API name
/// </summary>
public static class MainServiceAction
{
    /// <summary>
    /// 扣除nft
    /// </summary>
    public const string TAKE_NFT = "MainServiceActionTakeNFT";
    /// <summary>
    /// 获取所有建造物数据 for main service
    /// </summary>
    public const string GET_ALL_BUILD = "MainServiceActionGetAllBuild";

}

/// <summary>
/// MELAND proto message Action calls method name 
/// </summary>
public static class ProtoMessageAction
{
    /// <summary>
    /// 转发客户端proto Message
    /// </summary>
    public const string PULL_CLIENT_MESSAGE = "PullClientMessage";
    /// <summary>
    /// services 推送 proto message 到客户端 
    /// </summary>
    public const string BROAD_CAST_TO_CLIENT = "BroadCastToClient";
    /// <summary>
    /// services 批量推送 proto message 到客户端 
    /// </summary>
    public const string MULTIPLE_BROAD_CAST_TO_CLIENT = "MultipleBroadCastToClient";
}

/// <summary>
/// MELAND user action method api
/// </summary>
public static class UserAction
{
    /// <summary>
    /// 更新玩家战斗属性(升级/装备槽等级变更)
    /// </summary>
    public const string UPDATE_USER_PROFILE = "UpdateUserProfile";
    /// <summary>
    /// 更新玩家使用的装备
    /// </summary>
    public const string UPDATE_USED_AVATAR = "UpdateUsedAvatar";
    /// <summary>
    ///  查询玩家详细数据 for main service
    /// </summary>
    public const string GET_USER_DATA = "GetUserData";

}