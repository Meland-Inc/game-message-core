package xlsxTable

func TableModels() []interface{} {
	return []interface{}{
		BuffTableRow{},
		ChatTableRow{},
		DropTableRow{},
		GameValueTable{},
		ItemTable{},
		MonsterTableRow{},
		RewardTableRow{},
		RoleLvTableRow{},
		SlotLvTableRow{},
		TaskTableRow{},
		TaskListTableRow{},
		SceneAreaRow{},
	}
}
