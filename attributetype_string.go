// Code generated by "stringer -type=attributeType"; DO NOT EDIT.

package conntrack

import "strconv"

const _attributeType_name = "ctaUnspecctaTupleOrigctaTupleReplyctaStatusctaProtoInfoctaHelpctaNatSrcctaTimeoutctaMarkctaCountersOrigctaCountersReplyctaUsectaIDctaNatDstctaTupleMasterctaSeqAdjOrigctaSeqAdjReplyctaSecMarkctaZonectaSecCtxctaTimestampctaMarkMaskctaLabelsctaLabelsMaskctaSynProxy"

var _attributeType_index = [...]uint16{0, 9, 21, 34, 43, 55, 62, 71, 81, 88, 103, 119, 125, 130, 139, 153, 166, 180, 190, 197, 206, 218, 229, 238, 251, 262}

func (i attributeType) String() string {
	if i >= attributeType(len(_attributeType_index)-1) {
		return "attributeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _attributeType_name[_attributeType_index[i]:_attributeType_index[i+1]]
}
