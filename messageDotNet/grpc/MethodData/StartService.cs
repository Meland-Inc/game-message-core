using System;
using UnityEngine;

/// <summary>
/// 向manager service 请求启动服务 request
/// </summary>
[Serializable]
public class StartServiceInput
{
    public GameMessageCore.ServiceType ServiceType;
    public GameMessageCore.SceneServiceSubType SceneSerSubType;
    public long OwnerId;
    public int MapId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


/// <summary>
/// 向manager service 请求启动服务  response
/// </summary>
[Serializable]
public class StartServiceOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}
