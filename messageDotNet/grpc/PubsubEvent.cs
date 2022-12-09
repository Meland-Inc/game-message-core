
/// <summary>
/// MELAND service pubsub event topic name 
/// </summary>
public static class SubscriptionEvent
{
    /// <summary>
    /// 服务启动完成事件
    /// </summary>
    public const string SERVICE_STARTED = "EventTopicServiceStarted";

    /// <summary>
    /// 服务销毁事件
    /// </summary>
    public const string SERVICE_UNREGISTER = "EventTopicServiceUnregister";

    /// <summary>
    /// 玩家请求进入服务 response event
    /// </summary>
    public const string APPLY_ENTER_SERVICE_RES = "EventTopicUserApplyEnterServiceRes";
    /// <summary>
    /// 玩家正式进入服务 response event
    /// </summary>
    public const string JOIN_SERVICE_RES = "EventTopicUserJoinServiceRes";

    /// <summary>
    /// 玩家完成服务切换事件
    /// </summary>
    public const string USER_CHANGE_SERVICE = "EventTopicUserChangeService";

    /// <summary>
    /// 存储玩家数据
    /// </summary>
    public const string SAVE_PLAYER_DATA = "EventTopicSavePlayerData";
    /// <summary>
    /// 击杀怪物
    /// </summary>
    public const string KILL_MONSTER = "EventTopicKillMonster";
    /// <summary>
    /// 玩家死亡
    /// </summary>
    public const string PLAYER_DEATH = "EventTopicPlayerDeath";
    /// <summary>
    /// use NFT event
    /// </summary>
    public const string USE_NFT = "EventTopicUseNFT";
    /// <summary>
    /// user enter game event
    /// </summary>
    public const string USER_ENTER_GAME = "EventTopicUserEnterGame";
    /// <summary>
    /// 玩家离开游戏(断开网络链接时)
    /// </summary>
    public const string USER_LEAVE_GAME = "EventTopicUserLeaveGame";
    /// <summary>
    /// add nft build event
    /// </summary>
    public const string NFT_BUILD_ADD = "EventTopicNftBuildAdd";

    /// <summary>
    /// update nft build event
    /// </summary>
    public const string NFT_BUILD_UPDATE = "EventTopicNftBuildUpdate";
    /// <summary>
    /// remove nft build event
    /// </summary>
    public const string NFT_BUILD_REMOVE = "EventTopicNftBuildRemove";
    /// <summary>
    /// task reward event
    /// </summary>
    public const string TASK_REWARD = "EventTopicUserTaskReward";
    /// <summary>
    /// task finish event
    /// </summary>
    public const string TASK_FINISH = "EventTopicTaskFinish";
    /// <summary>
    /// task list finish event
    /// </summary>
    public const string TASK_LIST_FINISH = "EventTopicTaskListFinish";
    /// <summary>
    /// slot level update event
    /// </summary>
    public const string SLOT_LEVEL_UPGRADE = "EventTopicSlotLevelUpgrade";
    /// <summary>
    /// user level update event
    /// </summary>
    public const string USER_LEVEL_UPGRADE = "EventTopicUserLevelUpgrade";

    public const string TICK_OUT_PLAYER = "EventTopicTickOutPlayer";
    /// <summary>
    /// user level update event
    /// </summary>
    public const string SAVE_HOME_DATA = "EventTopicSaveHomeData";
    /// <summary>
    /// home Granary Stockpile Event
    /// </summary>
    public const string GRANARY_STOCKPILE = "EventTopicGranaryStockpile";



}

