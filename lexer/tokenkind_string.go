// Code generated by "stringer -type=TokenKind -trimprefix Token ./lexer"; DO NOT EDIT.

package lexer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TokenErr-0]
	_ = x[TokenBraceOpen-1]
	_ = x[TokenBraceClose-2]
	_ = x[TokenNewLine-3]
	_ = x[TokenHorizontalLine-4]
	_ = x[TokenPreformattedFence-5]
	_ = x[TokenPreformattedAltText-6]
	_ = x[TokenHeadingOpen-7]
	_ = x[TokenHeadingClose-8]
	_ = x[TokenSpanNewLine-9]
	_ = x[TokenSpanText-10]
	_ = x[TokenSpanItalic-11]
	_ = x[TokenSpanBold-12]
	_ = x[TokenSpanMonospace-13]
	_ = x[TokenSpanMarker-14]
	_ = x[TokenSpanSuper-15]
	_ = x[TokenSpanSub-16]
	_ = x[TokenSpanStrike-17]
	_ = x[TokenSpanLinkOpen-18]
	_ = x[TokenSpanLinkClose-19]
	_ = x[TokenLinkAddress-20]
	_ = x[TokenLinkDisplay-21]
	_ = x[TokenAutoLink-22]
	_ = x[TokenRocketLinkOpen-23]
	_ = x[TokenRocketLinkClose-24]
	_ = x[TokenBlockQuoteOpen-25]
	_ = x[TokenBlockQuoteClose-26]
	_ = x[TokenBulletUnnumbered-27]
	_ = x[TokenBulletIndent-28]
	_ = x[TokenBulletNumberedImplicit-29]
	_ = x[TokenBulletNumberedExplicit-30]
	_ = x[TokenTagImg-31]
	_ = x[TokenTagTable-32]
}

const _TokenKind_name = "ErrBraceOpenBraceCloseNewLineHorizontalLinePreformattedFencePreformattedAltTextHeadingOpenHeadingCloseSpanNewLineSpanTextSpanItalicSpanBoldSpanMonospaceSpanMarkerSpanSuperSpanSubSpanStrikeSpanLinkOpenSpanLinkCloseLinkAddressLinkDisplayAutoLinkRocketLinkOpenRocketLinkCloseBlockQuoteOpenBlockQuoteCloseBulletUnnumberedBulletIndentBulletNumberedImplicitBulletNumberedExplicitTagImgTagTable"

var _TokenKind_index = [...]uint16{0, 3, 12, 22, 29, 43, 60, 79, 90, 102, 113, 121, 131, 139, 152, 162, 171, 178, 188, 200, 213, 224, 235, 243, 257, 272, 286, 301, 317, 329, 351, 373, 379, 387}

func (i TokenKind) String() string {
	if i < 0 || i >= TokenKind(len(_TokenKind_index)-1) {
		return "TokenKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenKind_name[_TokenKind_index[i]:_TokenKind_index[i+1]]
}