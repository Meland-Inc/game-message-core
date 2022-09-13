using System;
using UnityEngine;

/// <summary>
/// player exit game 
/// </summary>
[Serializable]
internal class PlayerLeaveGameInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string AgentAppId;
    public long UserId;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


/// <summary>
/// agent service forward client message response
/// </summary>
[Serializable]
internal class PlayerLeaveGameOutput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string AgentAppId;
    public long UserId;
    public bool Success;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}