
package parser

// Autogenerated code from ese_profile.json. Do not edit.

import (
    "encoding/binary"
    "fmt"
    "bytes"
    "io"
    "sort"
    "strings"
    "unicode/utf16"
    "unicode/utf8"
)

var (
   // Depending on autogenerated code we may use this. Add a reference
   // to shut the compiler up.
   _ = bytes.MinRead
   _ = fmt.Sprintf
   _ = utf16.Decode
   _ = binary.LittleEndian
   _ = utf8.RuneError
   _ = sort.Strings
   _ = strings.Join
   _ = io.Copy
)

func indent(text string) string {
    result := []string{}
    lines := strings.Split(text,"\n")
    for _, line := range lines {
         result = append(result, "  " + line)
    }
    return strings.Join(result, "\n")
}


type ESEProfile struct {
    Off_CATALOG_TYPE_COLUMN_ColumnType int64
    Off_CATALOG_TYPE_COLUMN_SpaceUsage int64
    Off_CATALOG_TYPE_COLUMN_ColumnFlags int64
    Off_CATALOG_TYPE_COLUMN_CodePage int64
    Off_CATALOG_TYPE_INDEX_FatherDataPageNumber int64
    Off_CATALOG_TYPE_INDEX_SpaceUsage int64
    Off_CATALOG_TYPE_INDEX_IndexFlags int64
    Off_CATALOG_TYPE_INDEX_Locale int64
    Off_CATALOG_TYPE_LONG_VALUE_FatherDataPageNumber int64
    Off_CATALOG_TYPE_LONG_VALUE_SpaceUsage int64
    Off_CATALOG_TYPE_LONG_VALUE_LVFlags int64
    Off_CATALOG_TYPE_LONG_VALUE_InitialNumberOfPages int64
    Off_CATALOG_TYPE_TABLE_FatherDataPageNumber int64
    Off_CATALOG_TYPE_TABLE_SpaceUsage int64
    Off_DBTime_Hours int64
    Off_DBTime_Min int64
    Off_DBTime_Sec int64
    Off_ESENT_BRANCH_ENTRY_LocalPageKeySize int64
    Off_ESENT_BRANCH_HEADER_CommonPageKey int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_FDPId int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Type int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Identifier int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Column int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Table int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Index int64
    Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_LongValue int64
    Off_ESENT_DATA_DEFINITION_HEADER_LastFixedType int64
    Off_ESENT_DATA_DEFINITION_HEADER_LastVariableDataType int64
    Off_ESENT_DATA_DEFINITION_HEADER_VariableSizeOffset int64
    Off_ESENT_INDEX_ENTRY_RecordPageKey int64
    Off_ESENT_LEAF_ENTRY_CommonPageKeySize int64
    Off_ESENT_LEAF_ENTRY_LocalPageKeySize int64
    Off_ESENT_LEAF_HEADER_CommonPageKey int64
    Off_ESENT_ROOT_HEADER_InitialNumberOfPages int64
    Off_ESENT_ROOT_HEADER_ParentFDP int64
    Off_ESENT_ROOT_HEADER_ExtentSpace int64
    Off_ESENT_ROOT_HEADER_SpaceTreePageNumber int64
    Off_ESENT_SPACE_TREE_ENTRY_PageKeySize int64
    Off_ESENT_SPACE_TREE_ENTRY_LastPageNumber int64
    Off_ESENT_SPACE_TREE_ENTRY_NumberOfPages int64
    Off_FileHeader_Magic int64
    Off_FileHeader_FormatVersion int64
    Off_FileHeader_FormatRevision int64
    Off_FileHeader_FileType int64
    Off_FileHeader_DataBaseTime int64
    Off_FileHeader_Signature int64
    Off_FileHeader_PageSize int64
    Off_JET_LOGTIME_Sec int64
    Off_JET_LOGTIME_Min int64
    Off_JET_LOGTIME_Hours int64
    Off_JET_LOGTIME_Days int64
    Off_JET_LOGTIME_Month int64
    Off_JET_LOGTIME_Year int64
    Off_JET_SIGNATURE_Creation int64
    Off_JET_SIGNATURE_CreatorMachine int64
    Off_Misc_Misc int64
    Off_Misc_Misc2 int64
    Off_Misc_Misc3 int64
    Off_Misc_Misc5 int64
    Off_Misc_Misc4 int64
    Off_PageHeader_LastModified int64
    Off_PageHeader_PreviousPageNumber int64
    Off_PageHeader_NextPageNumber int64
    Off_PageHeader_FatherPage int64
    Off_PageHeader_AvailableDataSize int64
    Off_PageHeader_AvailableDataOffset int64
    Off_PageHeader_AvailablePageTag int64
    Off_PageHeader_Flags int64
    Off_RecordTag_Identifier int64
    Off_RecordTag_DataOffset int64
    Off_Tag__ValueSize int64
    Off_Tag__ValueOffset int64
}

func NewESEProfile() *ESEProfile {
    // Specific offsets can be tweaked to cater for slight version mismatches.
    self := &ESEProfile{0,4,8,12,0,4,8,12,0,4,8,12,0,4,0,2,4,0,0,0,4,6,10,10,10,10,0,1,2,0,-2,0,0,0,4,8,12,0,0,0,4,8,232,12,16,24,236,0,1,2,3,4,5,4,12,0,0,0,0,0,8,16,20,24,28,32,34,36,0,2,0,2}
    return self
}

func (self *ESEProfile) CATALOG_TYPE_COLUMN(reader io.ReaderAt, offset int64) *CATALOG_TYPE_COLUMN {
    return &CATALOG_TYPE_COLUMN{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) CATALOG_TYPE_INDEX(reader io.ReaderAt, offset int64) *CATALOG_TYPE_INDEX {
    return &CATALOG_TYPE_INDEX{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) CATALOG_TYPE_LONG_VALUE(reader io.ReaderAt, offset int64) *CATALOG_TYPE_LONG_VALUE {
    return &CATALOG_TYPE_LONG_VALUE{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) CATALOG_TYPE_TABLE(reader io.ReaderAt, offset int64) *CATALOG_TYPE_TABLE {
    return &CATALOG_TYPE_TABLE{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) DBTime(reader io.ReaderAt, offset int64) *DBTime {
    return &DBTime{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_BRANCH_ENTRY(reader io.ReaderAt, offset int64) *ESENT_BRANCH_ENTRY {
    return &ESENT_BRANCH_ENTRY{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_BRANCH_HEADER(reader io.ReaderAt, offset int64) *ESENT_BRANCH_HEADER {
    return &ESENT_BRANCH_HEADER{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_CATALOG_DATA_DEFINITION_ENTRY(reader io.ReaderAt, offset int64) *ESENT_CATALOG_DATA_DEFINITION_ENTRY {
    return &ESENT_CATALOG_DATA_DEFINITION_ENTRY{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_DATA_DEFINITION_HEADER(reader io.ReaderAt, offset int64) *ESENT_DATA_DEFINITION_HEADER {
    return &ESENT_DATA_DEFINITION_HEADER{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_INDEX_ENTRY(reader io.ReaderAt, offset int64) *ESENT_INDEX_ENTRY {
    return &ESENT_INDEX_ENTRY{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_LEAF_ENTRY(reader io.ReaderAt, offset int64) *ESENT_LEAF_ENTRY {
    return &ESENT_LEAF_ENTRY{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_LEAF_HEADER(reader io.ReaderAt, offset int64) *ESENT_LEAF_HEADER {
    return &ESENT_LEAF_HEADER{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_ROOT_HEADER(reader io.ReaderAt, offset int64) *ESENT_ROOT_HEADER {
    return &ESENT_ROOT_HEADER{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_SPACE_TREE_ENTRY(reader io.ReaderAt, offset int64) *ESENT_SPACE_TREE_ENTRY {
    return &ESENT_SPACE_TREE_ENTRY{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) ESENT_SPACE_TREE_HEADER(reader io.ReaderAt, offset int64) *ESENT_SPACE_TREE_HEADER {
    return &ESENT_SPACE_TREE_HEADER{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) FileHeader(reader io.ReaderAt, offset int64) *FileHeader {
    return &FileHeader{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) JET_LOGTIME(reader io.ReaderAt, offset int64) *JET_LOGTIME {
    return &JET_LOGTIME{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) JET_SIGNATURE(reader io.ReaderAt, offset int64) *JET_SIGNATURE {
    return &JET_SIGNATURE{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) Misc(reader io.ReaderAt, offset int64) *Misc {
    return &Misc{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) PageHeader(reader io.ReaderAt, offset int64) *PageHeader {
    return &PageHeader{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) RecordTag(reader io.ReaderAt, offset int64) *RecordTag {
    return &RecordTag{Reader: reader, Offset: offset, Profile: self}
}

func (self *ESEProfile) Tag(reader io.ReaderAt, offset int64) *Tag {
    return &Tag{Reader: reader, Offset: offset, Profile: self}
}


type CATALOG_TYPE_COLUMN struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *CATALOG_TYPE_COLUMN) Size() int {
    return 0
}

func (self *CATALOG_TYPE_COLUMN) ColumnType() *Enumeration {
   value := ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_COLUMN_ColumnType + self.Offset)
   name := "Unknown"
   switch value {

      case 0:
         name = "NULL"

      case 1:
         name = "Boolean"

      case 2:
         name = "Signed byte"

      case 3:
         name = "Signed short"

      case 4:
         name = "Signed long"

      case 5:
         name = "Currency"

      case 6:
         name = "Single precision FP"

      case 7:
         name = "Double precision FP"

      case 8:
         name = "DateTime"

      case 9:
         name = "Binary"

      case 10:
         name = "Text"

      case 11:
         name = "Long Binary"

      case 12:
         name = "Long Text"

      case 13:
         name = "Obsolete"

      case 14:
         name = "Unsigned long"

      case 15:
         name = "Long long"

      case 16:
         name = "GUID"

      case 17:
         name = "Unsigned short"

      case 18:
         name = "Max"
}
   return &Enumeration{Value: uint64(value), Name: name}
}


func (self *CATALOG_TYPE_COLUMN) SpaceUsage() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_COLUMN_SpaceUsage + self.Offset)
}

func (self *CATALOG_TYPE_COLUMN) ColumnFlags() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_COLUMN_ColumnFlags + self.Offset)
}

func (self *CATALOG_TYPE_COLUMN) CodePage() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_COLUMN_CodePage + self.Offset)
}
func (self *CATALOG_TYPE_COLUMN) DebugString() string {
    result := fmt.Sprintf("struct CATALOG_TYPE_COLUMN @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  ColumnType: %v\n", self.ColumnType().DebugString())
    result += fmt.Sprintf("  SpaceUsage: %#0x\n", self.SpaceUsage())
    result += fmt.Sprintf("  ColumnFlags: %#0x\n", self.ColumnFlags())
    result += fmt.Sprintf("  CodePage: %#0x\n", self.CodePage())
    return result
}

type CATALOG_TYPE_INDEX struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *CATALOG_TYPE_INDEX) Size() int {
    return 0
}

func (self *CATALOG_TYPE_INDEX) FatherDataPageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_INDEX_FatherDataPageNumber + self.Offset)
}

func (self *CATALOG_TYPE_INDEX) SpaceUsage() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_INDEX_SpaceUsage + self.Offset)
}

func (self *CATALOG_TYPE_INDEX) IndexFlags() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_INDEX_IndexFlags + self.Offset)
}

func (self *CATALOG_TYPE_INDEX) Locale() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_INDEX_Locale + self.Offset)
}
func (self *CATALOG_TYPE_INDEX) DebugString() string {
    result := fmt.Sprintf("struct CATALOG_TYPE_INDEX @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FatherDataPageNumber: %#0x\n", self.FatherDataPageNumber())
    result += fmt.Sprintf("  SpaceUsage: %#0x\n", self.SpaceUsage())
    result += fmt.Sprintf("  IndexFlags: %#0x\n", self.IndexFlags())
    result += fmt.Sprintf("  Locale: %#0x\n", self.Locale())
    return result
}

type CATALOG_TYPE_LONG_VALUE struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *CATALOG_TYPE_LONG_VALUE) Size() int {
    return 0
}

func (self *CATALOG_TYPE_LONG_VALUE) FatherDataPageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_LONG_VALUE_FatherDataPageNumber + self.Offset)
}

func (self *CATALOG_TYPE_LONG_VALUE) SpaceUsage() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_LONG_VALUE_SpaceUsage + self.Offset)
}

func (self *CATALOG_TYPE_LONG_VALUE) LVFlags() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_LONG_VALUE_LVFlags + self.Offset)
}

func (self *CATALOG_TYPE_LONG_VALUE) InitialNumberOfPages() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_LONG_VALUE_InitialNumberOfPages + self.Offset)
}
func (self *CATALOG_TYPE_LONG_VALUE) DebugString() string {
    result := fmt.Sprintf("struct CATALOG_TYPE_LONG_VALUE @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FatherDataPageNumber: %#0x\n", self.FatherDataPageNumber())
    result += fmt.Sprintf("  SpaceUsage: %#0x\n", self.SpaceUsage())
    result += fmt.Sprintf("  LVFlags: %#0x\n", self.LVFlags())
    result += fmt.Sprintf("  InitialNumberOfPages: %#0x\n", self.InitialNumberOfPages())
    return result
}

type CATALOG_TYPE_TABLE struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *CATALOG_TYPE_TABLE) Size() int {
    return 0
}

func (self *CATALOG_TYPE_TABLE) FatherDataPageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_TABLE_FatherDataPageNumber + self.Offset)
}

func (self *CATALOG_TYPE_TABLE) SpaceUsage() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_CATALOG_TYPE_TABLE_SpaceUsage + self.Offset)
}
func (self *CATALOG_TYPE_TABLE) DebugString() string {
    result := fmt.Sprintf("struct CATALOG_TYPE_TABLE @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FatherDataPageNumber: %#0x\n", self.FatherDataPageNumber())
    result += fmt.Sprintf("  SpaceUsage: %#0x\n", self.SpaceUsage())
    return result
}

type DBTime struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *DBTime) Size() int {
    return 8
}

func (self *DBTime) Hours() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_DBTime_Hours + self.Offset)
}

func (self *DBTime) Min() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_DBTime_Min + self.Offset)
}

func (self *DBTime) Sec() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_DBTime_Sec + self.Offset)
}
func (self *DBTime) DebugString() string {
    result := fmt.Sprintf("struct DBTime @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Hours: %#0x\n", self.Hours())
    result += fmt.Sprintf("  Min: %#0x\n", self.Min())
    result += fmt.Sprintf("  Sec: %#0x\n", self.Sec())
    return result
}

type ESENT_BRANCH_ENTRY struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_BRANCH_ENTRY) Size() int {
    return 16
}

func (self *ESENT_BRANCH_ENTRY) LocalPageKeySize() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_ESENT_BRANCH_ENTRY_LocalPageKeySize + self.Offset)
}
func (self *ESENT_BRANCH_ENTRY) DebugString() string {
    result := fmt.Sprintf("struct ESENT_BRANCH_ENTRY @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  LocalPageKeySize: %#0x\n", self.LocalPageKeySize())
    return result
}

type ESENT_BRANCH_HEADER struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_BRANCH_HEADER) Size() int {
    return 16
}


func (self *ESENT_BRANCH_HEADER) CommonPageKey() string {
  return ParseTerminatedString(self.Reader, self.Profile.Off_ESENT_BRANCH_HEADER_CommonPageKey + self.Offset)
}
func (self *ESENT_BRANCH_HEADER) DebugString() string {
    result := fmt.Sprintf("struct ESENT_BRANCH_HEADER @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  CommonPageKey: %v\n", string(self.CommonPageKey()))
    return result
}

type ESENT_CATALOG_DATA_DEFINITION_ENTRY struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) Size() int {
    return 0
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) FDPId() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_FDPId + self.Offset)
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) Type() *Enumeration {
   value := ParseUint16(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Type + self.Offset)
   name := "Unknown"
   switch value {

      case 1:
         name = "CATALOG_TYPE_TABLE"

      case 2:
         name = "CATALOG_TYPE_COLUMN"

      case 3:
         name = "CATALOG_TYPE_INDEX"

      case 4:
         name = "CATALOG_TYPE_LONG_VALUE"
}
   return &Enumeration{Value: uint64(value), Name: name}
}


func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) Identifier() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Identifier + self.Offset)
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) Column() *CATALOG_TYPE_COLUMN {
    return self.Profile.CATALOG_TYPE_COLUMN(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Column + self.Offset)
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) Table() *CATALOG_TYPE_TABLE {
    return self.Profile.CATALOG_TYPE_TABLE(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Table + self.Offset)
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) Index() *CATALOG_TYPE_INDEX {
    return self.Profile.CATALOG_TYPE_INDEX(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_Index + self.Offset)
}

func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) LongValue() *CATALOG_TYPE_LONG_VALUE {
    return self.Profile.CATALOG_TYPE_LONG_VALUE(self.Reader, self.Profile.Off_ESENT_CATALOG_DATA_DEFINITION_ENTRY_LongValue + self.Offset)
}
func (self *ESENT_CATALOG_DATA_DEFINITION_ENTRY) DebugString() string {
    result := fmt.Sprintf("struct ESENT_CATALOG_DATA_DEFINITION_ENTRY @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  FDPId: %#0x\n", self.FDPId())
    result += fmt.Sprintf("  Type: %v\n", self.Type().DebugString())
    result += fmt.Sprintf("  Identifier: %#0x\n", self.Identifier())
    result += fmt.Sprintf("  Column: {\n%v}\n", indent(self.Column().DebugString()))
    result += fmt.Sprintf("  Table: {\n%v}\n", indent(self.Table().DebugString()))
    result += fmt.Sprintf("  Index: {\n%v}\n", indent(self.Index().DebugString()))
    result += fmt.Sprintf("  LongValue: {\n%v}\n", indent(self.LongValue().DebugString()))
    return result
}

type ESENT_DATA_DEFINITION_HEADER struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_DATA_DEFINITION_HEADER) Size() int {
    return 4
}

func (self *ESENT_DATA_DEFINITION_HEADER) LastFixedType() int8 {
   return ParseInt8(self.Reader, self.Profile.Off_ESENT_DATA_DEFINITION_HEADER_LastFixedType + self.Offset)
}

func (self *ESENT_DATA_DEFINITION_HEADER) LastVariableDataType() byte {
   return ParseUint8(self.Reader, self.Profile.Off_ESENT_DATA_DEFINITION_HEADER_LastVariableDataType + self.Offset)
}

func (self *ESENT_DATA_DEFINITION_HEADER) VariableSizeOffset() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_ESENT_DATA_DEFINITION_HEADER_VariableSizeOffset + self.Offset)
}
func (self *ESENT_DATA_DEFINITION_HEADER) DebugString() string {
    result := fmt.Sprintf("struct ESENT_DATA_DEFINITION_HEADER @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  LastFixedType: %#0x\n", self.LastFixedType())
    result += fmt.Sprintf("  LastVariableDataType: %#0x\n", self.LastVariableDataType())
    result += fmt.Sprintf("  VariableSizeOffset: %#0x\n", self.VariableSizeOffset())
    return result
}

type ESENT_INDEX_ENTRY struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_INDEX_ENTRY) Size() int {
    return 16
}


func (self *ESENT_INDEX_ENTRY) RecordPageKey() string {
  return ParseTerminatedString(self.Reader, self.Profile.Off_ESENT_INDEX_ENTRY_RecordPageKey + self.Offset)
}
func (self *ESENT_INDEX_ENTRY) DebugString() string {
    result := fmt.Sprintf("struct ESENT_INDEX_ENTRY @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  RecordPageKey: %v\n", string(self.RecordPageKey()))
    return result
}

type ESENT_LEAF_ENTRY struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_LEAF_ENTRY) Size() int {
    return 16
}

func (self *ESENT_LEAF_ENTRY) CommonPageKeySize() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_ESENT_LEAF_ENTRY_CommonPageKeySize + self.Offset)
}

func (self *ESENT_LEAF_ENTRY) LocalPageKeySize() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_ESENT_LEAF_ENTRY_LocalPageKeySize + self.Offset)
}
func (self *ESENT_LEAF_ENTRY) DebugString() string {
    result := fmt.Sprintf("struct ESENT_LEAF_ENTRY @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  CommonPageKeySize: %#0x\n", self.CommonPageKeySize())
    result += fmt.Sprintf("  LocalPageKeySize: %#0x\n", self.LocalPageKeySize())
    return result
}

type ESENT_LEAF_HEADER struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_LEAF_HEADER) Size() int {
    return 16
}


func (self *ESENT_LEAF_HEADER) CommonPageKey() string {
  return ParseTerminatedString(self.Reader, self.Profile.Off_ESENT_LEAF_HEADER_CommonPageKey + self.Offset)
}
func (self *ESENT_LEAF_HEADER) DebugString() string {
    result := fmt.Sprintf("struct ESENT_LEAF_HEADER @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  CommonPageKey: %v\n", string(self.CommonPageKey()))
    return result
}

type ESENT_ROOT_HEADER struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_ROOT_HEADER) Size() int {
    return 16
}

func (self *ESENT_ROOT_HEADER) InitialNumberOfPages() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_ROOT_HEADER_InitialNumberOfPages + self.Offset)
}

func (self *ESENT_ROOT_HEADER) ParentFDP() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_ROOT_HEADER_ParentFDP + self.Offset)
}

func (self *ESENT_ROOT_HEADER) ExtentSpace() *Enumeration {
   value := ParseUint32(self.Reader, self.Profile.Off_ESENT_ROOT_HEADER_ExtentSpace + self.Offset)
   name := "Unknown"
   switch value {

      case 0:
         name = "Single"

      case 1:
         name = "Multiple"
}
   return &Enumeration{Value: uint64(value), Name: name}
}


func (self *ESENT_ROOT_HEADER) SpaceTreePageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_ROOT_HEADER_SpaceTreePageNumber + self.Offset)
}
func (self *ESENT_ROOT_HEADER) DebugString() string {
    result := fmt.Sprintf("struct ESENT_ROOT_HEADER @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  InitialNumberOfPages: %#0x\n", self.InitialNumberOfPages())
    result += fmt.Sprintf("  ParentFDP: %#0x\n", self.ParentFDP())
    result += fmt.Sprintf("  ExtentSpace: %v\n", self.ExtentSpace().DebugString())
    result += fmt.Sprintf("  SpaceTreePageNumber: %#0x\n", self.SpaceTreePageNumber())
    return result
}

type ESENT_SPACE_TREE_ENTRY struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_SPACE_TREE_ENTRY) Size() int {
    return 16
}

func (self *ESENT_SPACE_TREE_ENTRY) PageKeySize() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_ESENT_SPACE_TREE_ENTRY_PageKeySize + self.Offset)
}

func (self *ESENT_SPACE_TREE_ENTRY) LastPageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_SPACE_TREE_ENTRY_LastPageNumber + self.Offset)
}

func (self *ESENT_SPACE_TREE_ENTRY) NumberOfPages() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_ESENT_SPACE_TREE_ENTRY_NumberOfPages + self.Offset)
}
func (self *ESENT_SPACE_TREE_ENTRY) DebugString() string {
    result := fmt.Sprintf("struct ESENT_SPACE_TREE_ENTRY @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  PageKeySize: %#0x\n", self.PageKeySize())
    result += fmt.Sprintf("  LastPageNumber: %#0x\n", self.LastPageNumber())
    result += fmt.Sprintf("  NumberOfPages: %#0x\n", self.NumberOfPages())
    return result
}

type ESENT_SPACE_TREE_HEADER struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *ESENT_SPACE_TREE_HEADER) Size() int {
    return 16
}
func (self *ESENT_SPACE_TREE_HEADER) DebugString() string {
    result := fmt.Sprintf("struct ESENT_SPACE_TREE_HEADER @ %#x:\n", self.Offset)
    return result
}

type FileHeader struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *FileHeader) Size() int {
    return 0
}

func (self *FileHeader) Magic() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileHeader_Magic + self.Offset)
}

func (self *FileHeader) FormatVersion() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileHeader_FormatVersion + self.Offset)
}

func (self *FileHeader) FormatRevision() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileHeader_FormatRevision + self.Offset)
}

func (self *FileHeader) FileType() *Enumeration {
   value := ParseUint32(self.Reader, self.Profile.Off_FileHeader_FileType + self.Offset)
   name := "Unknown"
   switch value {

      case 0:
         name = "Database"

      case 1:
         name = "StreamingFile"
}
   return &Enumeration{Value: uint64(value), Name: name}
}


func (self *FileHeader) DataBaseTime() *DBTime {
    return self.Profile.DBTime(self.Reader, self.Profile.Off_FileHeader_DataBaseTime + self.Offset)
}

func (self *FileHeader) Signature() *JET_SIGNATURE {
    return self.Profile.JET_SIGNATURE(self.Reader, self.Profile.Off_FileHeader_Signature + self.Offset)
}

func (self *FileHeader) PageSize() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_FileHeader_PageSize + self.Offset)
}
func (self *FileHeader) DebugString() string {
    result := fmt.Sprintf("struct FileHeader @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Magic: %#0x\n", self.Magic())
    result += fmt.Sprintf("  FormatVersion: %#0x\n", self.FormatVersion())
    result += fmt.Sprintf("  FormatRevision: %#0x\n", self.FormatRevision())
    result += fmt.Sprintf("  FileType: %v\n", self.FileType().DebugString())
    result += fmt.Sprintf("  DataBaseTime: {\n%v}\n", indent(self.DataBaseTime().DebugString()))
    result += fmt.Sprintf("  Signature: {\n%v}\n", indent(self.Signature().DebugString()))
    result += fmt.Sprintf("  PageSize: %#0x\n", self.PageSize())
    return result
}

type JET_LOGTIME struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *JET_LOGTIME) Size() int {
    return 8
}

func (self *JET_LOGTIME) Sec() byte {
   return ParseUint8(self.Reader, self.Profile.Off_JET_LOGTIME_Sec + self.Offset)
}

func (self *JET_LOGTIME) Min() byte {
   return ParseUint8(self.Reader, self.Profile.Off_JET_LOGTIME_Min + self.Offset)
}

func (self *JET_LOGTIME) Hours() byte {
   return ParseUint8(self.Reader, self.Profile.Off_JET_LOGTIME_Hours + self.Offset)
}

func (self *JET_LOGTIME) Days() byte {
   return ParseUint8(self.Reader, self.Profile.Off_JET_LOGTIME_Days + self.Offset)
}

func (self *JET_LOGTIME) Month() byte {
   return ParseUint8(self.Reader, self.Profile.Off_JET_LOGTIME_Month + self.Offset)
}

func (self *JET_LOGTIME) Year() byte {
   return ParseUint8(self.Reader, self.Profile.Off_JET_LOGTIME_Year + self.Offset)
}
func (self *JET_LOGTIME) DebugString() string {
    result := fmt.Sprintf("struct JET_LOGTIME @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Sec: %#0x\n", self.Sec())
    result += fmt.Sprintf("  Min: %#0x\n", self.Min())
    result += fmt.Sprintf("  Hours: %#0x\n", self.Hours())
    result += fmt.Sprintf("  Days: %#0x\n", self.Days())
    result += fmt.Sprintf("  Month: %#0x\n", self.Month())
    result += fmt.Sprintf("  Year: %#0x\n", self.Year())
    return result
}

type JET_SIGNATURE struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *JET_SIGNATURE) Size() int {
    return 28
}

func (self *JET_SIGNATURE) Creation() *JET_LOGTIME {
    return self.Profile.JET_LOGTIME(self.Reader, self.Profile.Off_JET_SIGNATURE_Creation + self.Offset)
}


func (self *JET_SIGNATURE) CreatorMachine() string {
  return ParseTerminatedString(self.Reader, self.Profile.Off_JET_SIGNATURE_CreatorMachine + self.Offset)
}
func (self *JET_SIGNATURE) DebugString() string {
    result := fmt.Sprintf("struct JET_SIGNATURE @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Creation: {\n%v}\n", indent(self.Creation().DebugString()))
    result += fmt.Sprintf("  CreatorMachine: %v\n", string(self.CreatorMachine()))
    return result
}

type Misc struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *Misc) Size() int {
    return 0
}

func (self *Misc) Misc() int32 {
   return ParseInt32(self.Reader, self.Profile.Off_Misc_Misc + self.Offset)
}

func (self *Misc) Misc2() int16 {
   return ParseInt16(self.Reader, self.Profile.Off_Misc_Misc2 + self.Offset)
}

func (self *Misc) Misc3() int64 {
    return int64(ParseUint64(self.Reader, self.Profile.Off_Misc_Misc3 + self.Offset))
}

func (self *Misc) Misc5() uint64 {
    return ParseUint64(self.Reader, self.Profile.Off_Misc_Misc5 + self.Offset)
}


func (self *Misc) Misc4() string {
  return ParseTerminatedUTF16String(self.Reader, self.Profile.Off_Misc_Misc4 + self.Offset)
}
func (self *Misc) DebugString() string {
    result := fmt.Sprintf("struct Misc @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Misc: %#0x\n", self.Misc())
    result += fmt.Sprintf("  Misc2: %#0x\n", self.Misc2())
    result += fmt.Sprintf("  Misc3: %#0x\n", self.Misc3())
    result += fmt.Sprintf("  Misc5: %#0x\n", self.Misc5())
    result += fmt.Sprintf("  Misc4: %v\n", string(self.Misc4()))
    return result
}

type PageHeader struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *PageHeader) Size() int {
    return 0
}

func (self *PageHeader) LastModified() *DBTime {
    return self.Profile.DBTime(self.Reader, self.Profile.Off_PageHeader_LastModified + self.Offset)
}

func (self *PageHeader) PreviousPageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_PageHeader_PreviousPageNumber + self.Offset)
}

func (self *PageHeader) NextPageNumber() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_PageHeader_NextPageNumber + self.Offset)
}

func (self *PageHeader) FatherPage() uint32 {
   return ParseUint32(self.Reader, self.Profile.Off_PageHeader_FatherPage + self.Offset)
}

func (self *PageHeader) AvailableDataSize() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_PageHeader_AvailableDataSize + self.Offset)
}

func (self *PageHeader) AvailableDataOffset() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_PageHeader_AvailableDataOffset + self.Offset)
}

func (self *PageHeader) AvailablePageTag() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_PageHeader_AvailablePageTag + self.Offset)
}

func (self *PageHeader) Flags() *Flags {
   value := ParseUint32(self.Reader, self.Profile.Off_PageHeader_Flags + self.Offset)
   names := make(map[string]bool)


   if value & 128 != 0 {
      names["Long"] = true
   }

   if value & 1 != 0 {
      names["Root"] = true
   }

   if value & 2 != 0 {
      names["Leaf"] = true
   }

   if value & 4 != 0 {
      names["Parent"] = true
   }

   if value & 8 != 0 {
      names["Empty"] = true
   }

   if value & 32 != 0 {
      names["SpaceTree"] = true
   }

   if value & 64 != 0 {
      names["Index"] = true
   }

   return &Flags{Value: uint64(value), Names: names}
}

func (self *PageHeader) DebugString() string {
    result := fmt.Sprintf("struct PageHeader @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  LastModified: {\n%v}\n", indent(self.LastModified().DebugString()))
    result += fmt.Sprintf("  PreviousPageNumber: %#0x\n", self.PreviousPageNumber())
    result += fmt.Sprintf("  NextPageNumber: %#0x\n", self.NextPageNumber())
    result += fmt.Sprintf("  FatherPage: %#0x\n", self.FatherPage())
    result += fmt.Sprintf("  AvailableDataSize: %#0x\n", self.AvailableDataSize())
    result += fmt.Sprintf("  AvailableDataOffset: %#0x\n", self.AvailableDataOffset())
    result += fmt.Sprintf("  AvailablePageTag: %#0x\n", self.AvailablePageTag())
    result += fmt.Sprintf("  Flags: %v\n", self.Flags().DebugString())
    return result
}

type RecordTag struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *RecordTag) Size() int {
    return 4
}

func (self *RecordTag) Identifier() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_RecordTag_Identifier + self.Offset)
}

func (self *RecordTag) DataOffset() uint64 {
   value := ParseUint16(self.Reader, self.Profile.Off_RecordTag_DataOffset + self.Offset)
   return (uint64(value) & 0x1fff) >> 0x0
}
func (self *RecordTag) DebugString() string {
    result := fmt.Sprintf("struct RecordTag @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  Identifier: %#0x\n", self.Identifier())
    result += fmt.Sprintf("  DataOffset: %#0x\n", self.DataOffset())
    return result
}

type Tag struct {
    Reader io.ReaderAt
    Offset int64
    Profile *ESEProfile
}

func (self *Tag) Size() int {
    return 8
}

func (self *Tag) _ValueSize() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_Tag__ValueSize + self.Offset)
}

func (self *Tag) _ValueOffset() uint16 {
   return ParseUint16(self.Reader, self.Profile.Off_Tag__ValueOffset + self.Offset)
}
func (self *Tag) DebugString() string {
    result := fmt.Sprintf("struct Tag @ %#x:\n", self.Offset)
    result += fmt.Sprintf("  _ValueSize: %#0x\n", self._ValueSize())
    result += fmt.Sprintf("  _ValueOffset: %#0x\n", self._ValueOffset())
    return result
}

type Enumeration struct {
    Value uint64
    Name  string
}

func (self Enumeration) DebugString() string {
    return fmt.Sprintf("%s (%d)", self.Name, self.Value)
}


type Flags struct {
    Value uint64
    Names  map[string]bool
}

func (self Flags) DebugString() string {
    names := []string{}
    for k, _ := range self.Names {
      names = append(names, k)
    }

    sort.Strings(names)

    return fmt.Sprintf("%d (%s)", self.Value, strings.Join(names, ","))
}

func (self Flags) IsSet(flag string) bool {
    result, _ := self.Names[flag]
    return result
}


func ParseInt16(reader io.ReaderAt, offset int64) int16 {
    data := make([]byte, 2)
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return int16(binary.LittleEndian.Uint16(data))
}

func ParseInt32(reader io.ReaderAt, offset int64) int32 {
    data := make([]byte, 4)
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return int32(binary.LittleEndian.Uint32(data))
}

func ParseInt64(reader io.ReaderAt, offset int64) int64 {
    data := make([]byte, 8)
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return int64(binary.LittleEndian.Uint64(data))
}

func ParseInt8(reader io.ReaderAt, offset int64) int8 {
    result := make([]byte, 1)
    _, err := reader.ReadAt(result, offset)
    if err != nil {
       return 0
    }
    return int8(result[0])
}

func ParseUint16(reader io.ReaderAt, offset int64) uint16 {
    data := make([]byte, 2)
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return binary.LittleEndian.Uint16(data)
}

func ParseUint32(reader io.ReaderAt, offset int64) uint32 {
    data := make([]byte, 4)
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return binary.LittleEndian.Uint32(data)
}

func ParseUint64(reader io.ReaderAt, offset int64) uint64 {
    data := make([]byte, 8)
    _, err := reader.ReadAt(data, offset)
    if err != nil {
       return 0
    }
    return binary.LittleEndian.Uint64(data)
}

func ParseUint8(reader io.ReaderAt, offset int64) byte {
    result := make([]byte, 1)
    _, err := reader.ReadAt(result, offset)
    if err != nil {
       return 0
    }
    return result[0]
}

func ParseTerminatedString(reader io.ReaderAt, offset int64) string {
   data := make([]byte, 1024)
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
     return ""
   }
   idx := bytes.Index(data[:n], []byte{0})
   if idx < 0 {
      idx = n
   }
   return string(data[0:idx])
}

func ParseString(reader io.ReaderAt, offset int64, length int64) string {
   data := make([]byte, length)
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
      return ""
   }
   return string(data[:n])
}


func ParseTerminatedUTF16String(reader io.ReaderAt, offset int64) string {
   data := make([]byte, 1024)
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
     return ""
   }

   idx := bytes.Index(data[:n], []byte{0, 0})
   if idx < 0 {
      idx = n-1
   }
   return UTF16BytesToUTF8(data[0:idx+1], binary.LittleEndian)
}

func ParseUTF16String(reader io.ReaderAt, offset int64, length int64) string {
   data := make([]byte, length)
   n, err := reader.ReadAt(data, offset)
   if err != nil && err != io.EOF {
     return ""
   }
   return UTF16BytesToUTF8(data[:n], binary.LittleEndian)
}

func UTF16BytesToUTF8(b []byte, o binary.ByteOrder) string {
	if len(b) < 2 {
		return ""
	}

	if b[0] == 0xff && b[1] == 0xfe {
		o = binary.BigEndian
		b = b[2:]
	} else if b[0] == 0xfe && b[1] == 0xff {
		o = binary.LittleEndian
		b = b[2:]
	}

	utf := make([]uint16, (len(b)+(2-1))/2)

	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}

	return string(utf16.Decode(utf))
}


