using System;
using UnityEngine;

/// <summary>
/// 玩家请求进入服务 response event
/// </summary>
[Serializable]
public class UserApplyEnterServiceResEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// 玩家正式进入服务 response event
/// </summary>
[Serializable]
public class UserJoinServiceResEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public bool Success;
    public string ErrMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// 玩家完成服务切换事件
/// </summary>
[Serializable]
public class UserChangeServiceEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public ServiceData FormService;
    public ServiceData ToService;
    public string UserAgentAppId;
    public string UserSocketId;
    public GrpcVector3 UserPosition;
    public GrpcVector3 UserDir;


    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}