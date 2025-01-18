package staticFireParser

type EntryHeader struct {
	seperator        rune
	decimalSeperator rune
	hasOneXColumn    bool
	operator         string
	date             string
	time             string
}
