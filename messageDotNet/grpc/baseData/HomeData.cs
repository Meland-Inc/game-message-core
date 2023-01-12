using System;
using UnityEngine;

// service data
[Serializable]
public class GrpcHomeData
{
    public string SoilJson;
    public string ResourceJson;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}