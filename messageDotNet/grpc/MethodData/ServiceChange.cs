/*
 * @Author: alex
 * @Date: 2022-12-9 10:00
 * @Description: user 切换 scene service 需要的grpc calls 相关交换协议
 * @FilePath: game-message-core/messageDotNet/grpc/MethodData/ServiceChange.cs
 */

using System;
using UnityEngine;

/// <summary>
/// user请求进入目标scene service request（预进入阶段）
/// </summary>
[Serializable]
public class UserApplyEnterServiceInput
{
    public long ApplyUser;
    public GrpcVector3 UserPosition;
    public GrpcVector3 UserDir;
    public ServiceData FromSer;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
/// <summary>
/// user 请求进入目标scene service  response （预进入阶段）
/// </summary>
[Serializable]
public class UserApplyEnterServiceOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// user正式进入目标scene service request（正式进入阶段）
/// </summary>
[Serializable]
public class UserJoinServiceInput
{
    public long UserId;
    public string UserAgentAppId;
    public string UserSocketId;
    public GrpcVector3 UserPosition;
    public GrpcVector3 UserDir;
    public ServiceData FromSer;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
/// <summary>
/// user正式进入目标scene service request（正式进入阶段）
/// </summary>
[Serializable]
public class UserJoinServiceOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}


