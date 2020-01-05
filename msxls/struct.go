package msxls

type rkRec struct {
	//2.5.218	RkRec
	//The RkRec structure contains the numeric data in an application-specific internal type for optimizing disk and memory space along with the corresponding IXFCell to the style record.
	/*
		2.5.168	IXFCell
		The IXFCell structure specifies the index of a cell XF.

		ixfe (2 bytes):
		An unsigned integer that specifies a zero-based index of a cell XF record in the collection
		of XF records in the Globals Substream.
		Cell XF records are the subset of XF records with an fStyle field equal to 0.
		This value MUST be greater than or equal to 15, or equal to 0.
		The value 0 indicates that this value MUST be ignored.
		See XFIndex for more information about the organization of XF records in the file.
	*/
	ixfe uint16 // An IXFCell that specifies the format of the numeric value.

	/*
		2.5.217	RkNumber
		The RkNumber structure specifies a numeric value.

		A - fX100 (1 bit): A bit that specifies whether num is the value of the RkNumber or 100 times the value of the RkNumber. MUST be a value from the following table:
			0	The value of RkNumber is the value of num.
			1	The value of RkNumber is the value of num divided by 100.


		B - fInt (1 bit): A bit that specifies the type of num.
		num (30 bits): A variable type field whose type and meaning is specified by the value of fInt, as defined in the following table:
			0	num is the 30 most significant bits of a 64-bit binary floating-point number as defined in [IEEE754]. The remaining 34-bits of the floating-point number MUST be 0.
			1	num is a signed integer.

	*/
	RkNumber []byte // An RkNumber that specifies the numeric value. 4
}

// 2.5.19
// The Cell structure specifies a cell in the current sheet.
type cell struct {
	rw  uint16 //  An Rw that specifies the row.
	col uint16 // A Col that specifies the column.

	/*
		The IXFCell structure specifies the index of a cell XF.
		specifies a zero-based index of a cell XF record
		in the collection of XF records in the Globals Substream.
		Cell XF records are the subset of XF records with an fStyle field equal to 0.
		This value MUST be greater than or equal to 15, or equal to 0.
		The value 0 indicates that this value MUST be ignored.
		See XFIndex for more information about the organization of XF records in the file.
	*/
	ixfe uint16 // An IXFCell that specifies the XF record.
}

/*
2.5.133	FormulaValue
The FormulaValue structure specifies the current value of a formula.
It can be a numeric value, a Boolean value, an error value, a string value, or a blank string value.
If fExprO is not 0xFFFF, the 8 bytes of this structure specify an Xnum (section 2.5.342).
If fExprO is 0xFFFF, this structure specifies a Boolean value, an error value, a string value, or a blank string value.
*/
type formulaValue struct {
	/*
		If fExprO is 0xFFFF, byte1 is an unsigned integer that specifies the formula value type and MUST be a value from the following table:
			0x00	String value. The string value is stored in a String record that immediately follows this record.
			0x01	Boolean value.
			0x02	Error value.
			0x03	Blank string value.
		If fExprO is not 0xFFFF, byte1 specifies the first byte of the Xnum.
	*/
	byte1 byte

	/*
		If fExprO is 0xFFFF, byte2 is undefined and MUST be ignored.
		If fExprO is not 0xFFFF, byte2 specifies the second byte of the Xnum (section 2.5.342).
	*/
	byte2 byte

	/*
		   The meaning of byte3 is specified in the following table:
				fExprO is 0xFFFF and byte1 is 0x00	byte3 is undefined and MUST be ignored.
				fExprO is 0xFFFF and byte1 is 0x01	byte3 specifies a Boolean value.
				fExprO is 0xFFFF and byte1 is 0x02	byte3 specifies a BErr.
				fExprO is 0xFFFF and byte1 is 0x03	byte3 is undefined and MUST be ignored.
				fExprO is not 0xFFFF	byte3 specifies the third byte of the Xnum.
	*/
	byte3 byte

	byte4  byte   // If fExprO is 0xFFFF, byte4 is undefined and MUST be ignored. If fExprO is not 0xFFFF, byte4 specifies the fourth byte of the Xnum.
	byte5  byte   //  If fExprO is 0xFFFF, byte5 is undefined and MUST be ignored. If fExprO is not 0xFFFF, byte5 specifies the fifth byte of the Xnum.
	byte6  byte   //  If fExprO is 0xFFFF, byte6 is undefined and MUST be ignored. If fExprO is not 0xFFFF, byte6 specifies the sixth byte of the Xnum.
	fExprO []byte // If fExprO is 0xFFFF, this structure specifies a Boolean value, an error value, a string value, or a blank string value. If fExprO is not 0xFFFF, fExprO specifies the last two bytes of the Xnum. size 2

}
