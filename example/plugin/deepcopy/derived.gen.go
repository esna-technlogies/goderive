// Code generated by goderive DO NOT EDIT.

package deepcopy

// deriveDeepCopy recursively copies the contents of src into dst.
func deriveDeepCopy(dst, src *MyStruct) {
	dst.Int64 = src.Int64
	if src.StringPtr == nil {
		dst.StringPtr = nil
	} else {
		dst.StringPtr = new(string)
		*dst.StringPtr = *src.StringPtr
	}
}
