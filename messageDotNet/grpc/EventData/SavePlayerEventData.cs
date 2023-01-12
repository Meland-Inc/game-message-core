using System;
using UnityEngine;

/// <summary>
/// 玩家数据保持到main server, event json 交互数据结构
/// </summary>
[Serializable]
public class SavePlayerEventData : EventDataBase
{
    public long MsgVersion;
    public ServiceData FormService;
    public int CurHP;
    public float DirX;
    public float DirY;
    public float DirZ;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}