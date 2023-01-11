using System;
using UnityEngine;


/// <summary>
///  更新玩家使用的装备
/// </summary>
[Serializable]
public class UpdateUsedAvatarInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public GrpcPlayerAvatar[] UsingAvatars;
    public GrpcEntityProfile CurProfile;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class UpdateUsedAvatarOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

/// <summary>
/// 更新玩家战斗属性(升级/装备槽等级变更)
/// </summary>
[Serializable]
public class UpdateUserProfileInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;

    public GrpcEntityProfile CurProfile;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class UpdateUserProfileOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

/// <summary>
/// 获取user 详细数据
/// </summary>
[Serializable]
public class GetUserDataInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class GetUserDataOutput
{
    public bool Success;
    public string ErrMsg;
    public GrpcPlayerBaseData BaseData;
    public GrpcEntityProfile Profile;
    public GrpcPlayerFeature Feature;
    public int MapId;
    public GrpcVector3 Pos;
    public GrpcVector3 Dir;
    public GrpcPlayerAvatar[] Avatars;
    public GrpcTalentData TalentData;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}