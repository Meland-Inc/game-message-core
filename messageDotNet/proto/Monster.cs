// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: monster.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GameMessageCore {

  /// <summary>Holder for reflection information generated from monster.proto</summary>
  public static partial class MonsterReflection {

    #region Descriptor
    /// <summary>File descriptor for monster.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MonsterReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg1tb25zdGVyLnByb3RvEg9nYW1lTWVzc2FnZUNvcmUqfAoOTW9uc3RlckF0",
            "dFR5cGUSGQoVTW9uc3RlckF0dFR5cGVVbmtub3duEAASHAoYTW9uc3RlckF0",
            "dFR5cGVJbml0aWF0aXZlEAESGQoVTW9uc3RlckF0dFR5cGVQYXNzaXZlEAIS",
            "FgoSTW9uc3RlckF0dFR5cGVEdW1iEANiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::GameMessageCore.MonsterAttType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum MonsterAttType {
    [pbr::OriginalName("MonsterAttTypeUnknown")] Unknown = 0,
    /// <summary>
    /// 主动攻击型
    /// </summary>
    [pbr::OriginalName("MonsterAttTypeInitiative")] Initiative = 1,
    /// <summary>
    /// 被动攻击型
    /// </summary>
    [pbr::OriginalName("MonsterAttTypePassive")] Passive = 2,
    /// <summary>
    /// 不攻击型
    /// </summary>
    [pbr::OriginalName("MonsterAttTypeDumb")] Dumb = 3,
  }

  #endregion

}

#endregion Designer generated code
