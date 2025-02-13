package staticfire

const AssertedFirstLine = "LabVIEW Measurement"   // The first line of the file
const AssertedEndOfHeader = "***End_of_Header***" // Denotes the end of a header section
const AssertedWriterVersion = "2"                 // This parser was written for version 2
const AssertedReaderVersion = "2"                 // Same as above
const AssertedTimePreferance = "Absolute"         // We only support absolute time
const AssertedDecimalSeparator = "."              // We only support decimal points for parsing floats
const AssertedXColumnPrefix = "X_Value"           // Assume all X columns start with this prefix
const AssertedCommentColumnName = "Comment"       // Assume all comment columns are named "Comment"
