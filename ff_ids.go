package twitter

type FfIds struct {
	Ids               []uint64 `json:"ids"`
	NextCursor        uint64   `json:"next_cursor"`
	NextCursorStr     string   `json:"next_cursor_str"`
	PreviousCursor    uint64   `json:"previous_cursor"`
	PreviousCursorStr string   `json:"previous_cursor_str"`
}
