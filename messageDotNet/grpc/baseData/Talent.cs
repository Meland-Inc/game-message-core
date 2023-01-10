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
            int nodeLength = Nodes.Length;
            tree.UnlockNodes = new GameMessageCore.TalentNode[nodeLength];
            for (int i = 0; i < nodeLength; i++)
            {
                tree.UnlockNodes[i] = Nodes[i];
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
            data.LevelData = new GameMessageCore.TalentLevel[lvLength];
            for (int i = 0; i < lvLength; i++)
            {
                data.LevelData[i] = Levels[i].ToProtoData();
            }
        }
        if (Trees != null)
        {
            int treeLength = Trees.Length;
            data.Tree = new GameMessageCore.TalentTree[treeLength];
            for (int i = 0; i < treeLength; i++)
            {
                data.Tree[i] = Trees[i].ToProtoData();
            }
        }
        return data;
    }
}

