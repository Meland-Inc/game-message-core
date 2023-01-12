using System;
using UnityEngine;

/// <summary>
/// select user home data for main service
/// </summary>
[Serializable]
public class MainServiceActionGetHomeDataInput
{
    public long UserId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionGetHomeDataOutput
{
    public bool Success;
    public string ErrMsg;
    public long UserId;
    public GrpcHomeData Data;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}