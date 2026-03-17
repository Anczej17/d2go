package memory

type Offset struct {
	GameData                    uintptr
	UnitTable                   uintptr
	UI                          uintptr
	Hover                       uintptr
	Expansion                   uintptr
	RosterOffset                uintptr
	PanelManagerContainerOffset uintptr
	WidgetStatesOffset          uintptr
	WaypointTableOffset         uintptr
	FPS                         uintptr
	KeyBindingsOffset           uintptr
	KeyBindingsSkillsOffset     uintptr
	QuestInfo                   uintptr
	TZ                          uintptr
	Quests                      uintptr
	Ping                        uintptr
	LegacyGraphics              uintptr
	CharData                    uintptr
	SelectedCharName            uintptr
	LastGameName                uintptr
	LastGamePassword            uintptr
}

func calculateOffsets(_ *Process) Offset {
	// UnitTable
	unitTableOffset := uintptr(0x1EAA350)

	// UI
	uiOffsetPtr := uintptr(0x1EBA042)

	// Hover
	hoverOffset := uintptr(0x1DFE010)

	// Expansion
	expOffset := uintptr(0x1DFD460)

	// Party members offset
	rosterOffset := uintptr(0x1EC0660)

	// PanelManagerContainer
	panelManagerContainerOffset := uintptr(0x1E14DB8)

	// WidgetStates
	WidgetStatesOffset := uintptr(0x1EE2678)

	// Waypoints
	WaypointTableOffset := uintptr(0x1D5C3C0)

	// FPS
	fpsOffset := uintptr(0x1D5C394)

	// KeyBindings
	keyBindingsOffset := uintptr(0x19D5594)

	// KeyBindings Skills
	keyBindingsSkillsOffset := uintptr(0x1DFE100)

	// QuestInfo
	questInfoOffset := uintptr(0x1EC6CD8)

	// Terror Zones
	tzOffset := uintptr(0x25B4990)

	// Ping
	pingOffset := uintptr(0x1DFD460)

	// LegacyGraphics
	legacyGfxOffset := uintptr(0x1EC6E7E)

	// CharData
	charDataOffset := uintptr(0x1E01710)

	// Selected Char Name
	selectedCharNameOffset := uintptr(0x1D53195)

	// Last Game Name
	lastGameNameOffset := uintptr(0x25FD2F0)

	// Last Game Password
	lastGamePasswordOffset := uintptr(0x25FD348)

	return Offset{
		UnitTable:                   unitTableOffset,
		UI:                          uiOffsetPtr,
		Hover:                       hoverOffset,
		Expansion:                   expOffset,
		RosterOffset:                rosterOffset,
		PanelManagerContainerOffset: panelManagerContainerOffset,
		WidgetStatesOffset:          WidgetStatesOffset,
		WaypointTableOffset:         WaypointTableOffset,
		FPS:                         fpsOffset,
		KeyBindingsOffset:           keyBindingsOffset,
		KeyBindingsSkillsOffset:     keyBindingsSkillsOffset,
		QuestInfo:                   questInfoOffset,
		TZ:                          tzOffset,
		Ping:                        pingOffset,
		LegacyGraphics:              legacyGfxOffset,
		CharData:                    charDataOffset,
		SelectedCharName:            selectedCharNameOffset,
		LastGameName:                lastGameNameOffset,
		LastGamePassword:            lastGamePasswordOffset,
	}
}
