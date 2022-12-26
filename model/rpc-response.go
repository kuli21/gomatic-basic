package model

import "encoding/xml"

//SUCCESS Response

type XResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	XParams XParams  `xml:"params"`
}
type XParams struct {
	XParam XParam `xml:"param"`
}
type XParam struct {
	XValue XValue `xml:"value"`
}
type XValue struct {
	XStruct XStruct `xml:"struct"`
}
type XStruct struct {
	XMember []XMember `xml:"member"`
}
type XMember struct {
	Name    string  `xml:"name"`
	Double  float64 `xml:"value>double"`
	I4      int     `xml:"value>i4"`
	Boolean int     `xml:"value>boolean"`
}

//ERROR Response

type EResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	EFault  EFault   `xml:"fault"`
}
type EFault struct {
	XValue XValue `xml:"value"`
}
