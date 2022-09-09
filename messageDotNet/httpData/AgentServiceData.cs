using System;
using UnityEngine;

/// <summary>
/// http请求网关数据的response
/// </summary>
[Serializable]
internal class AgentServiceResp
{
    public int ErrorCode;//TODO: NEED USE GLOBAL ERROR CODE
    public string ErrorMessage;
    public string Host;
    public int Port;
    public int Online;
    public int MaxOnline;
    public long CreatedAt;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}