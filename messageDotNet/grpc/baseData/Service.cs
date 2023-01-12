using System;
using UnityEngine;

// service data
[Serializable]
public class ServiceData
{
    public string APPID;
    public GameMessageCore.ServiceType ServiceType;
    public GameMessageCore.SceneServiceSubType SceneSerSubType;
    public long Owner;
    public string Host;
    public int Port;
    public int MapId;
    public int Online;
    public int MaxOnline;
    public long CreatedAt;
    public long UpdatedAt;


    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}