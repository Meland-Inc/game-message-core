using System;
using UnityEngine;

/// <summary>
/// save user home json data to main server
/// </summary>
[Serializable]
public class SaveHomeEvent : EventDataBase
{
    public long MsgVersion;
    public GrpcHomeData Data;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}