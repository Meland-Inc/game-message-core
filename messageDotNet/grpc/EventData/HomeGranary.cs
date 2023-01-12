using System;
using UnityEngine;

/// <summary>
/// 存放物品到家园仓库
/// </summary>
[Serializable]
public class GranaryStockpileEvent
{
    public long MsgVersion;
    public long HomeOwner;
    public GrpcItemBaseInfo[] Items;
    public long OccupantId;
    public string OccupantName;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}