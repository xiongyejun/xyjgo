package msxls

func (me *Worksheet) Less(i, j int) bool {
	if me.Cells[i].Row == me.Cells[j].Row {
		return me.Cells[i].Column < me.Cells[j].Column
	}

	return me.Cells[i].Row < me.Cells[j].Row
}
func (me *Worksheet) Len() int {
	return len(me.Cells)
}
func (me *Worksheet) Swap(i, j int) {
	me.Cells[i], me.Cells[j] = me.Cells[j], me.Cells[i]
}
