// Code generated by "stringer -type Type"; DO NOT EDIT

package token

import "fmt"

const _Type_name = "EOFErrorNewlineStringSpaceIntFloatHexLeftCurlyRightCurlyLeftParenRightParenLeftBracRightBracQuoteEqualColonCommaSemicolonPeriodCommentPlusPipeElipsisTrueFalseMultiLineStringTargetDeclFuncForIn"

var _Type_index = [...]uint8{0, 3, 8, 15, 21, 26, 29, 34, 37, 46, 56, 65, 75, 83, 92, 97, 102, 107, 112, 121, 127, 134, 138, 142, 149, 153, 158, 173, 183, 187, 190, 192}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
