using System;
using UnityEngine;


[Serializable]
public class GrpcTalentLevel
{
    public GameMessageCore.TalentType TalentType;
    public int Level;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
    public GameMessageCore.TalentLevel ToProtoData()
    {
        return new GameMessageCore.TalentLevel()
        {
            Type = TalentType,
            Level = Level,
        };
    }
}


[Serializable]
public class GrpcTalentNodeData
{
    public int NodeId;
    public int Level;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
    public GameMessageCore.TalentNode ToProtoData()
    {
        return new GameMessageCore.TalentNode()
        {
            NodeId = NodeId,
            Level = Level,
        };
    }
}



[Serializable]
public class GrpcTalentTree
{
    public GameMessageCore.TalentType TalentType;
    public GrpcTalentNodeData[] Nodes;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public GameMessageCore.TalentTree ToProtoData()
    {
        GameMessageCore.TalentTree tree = new() { Type = TalentType };
        if (Nodes != null)
        {
            for (int i = 0; i < Nodes.Length; i++)
            {
                tree.UnlockNodes.Add(Nodes[i].ToProtoData());
            }
        }
        return tree;
    }
}


[Serializable]
public class GrpcTalentData
{
    public GrpcTalentLevel[] Levels;
    public GrpcTalentTree[] Trees;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public GameMessageCore.TalentData ToProtoData()
    {
        GameMessageCore.TalentData data = new();
        if (Levels != null)
        {
            int lvLength = Levels.Length;
            for (int i = 0; i < lvLength; i++)
            {
                data.LevelData.Add(Levels[i].ToProtoData());
            }
        }
        if (Trees != null)
        {
            int treeLength = Trees.Length;
            for (int i = 0; i < treeLength; i++)
            {
                data.Tree.Add(Trees[i].ToProtoData());
            }
        }
        return data;
    }
}

