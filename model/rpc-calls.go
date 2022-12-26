package model

import "strings"

func ListDevicesReader() *strings.Reader {
	xmlString := `<?xml version="1.0"?>
	<methodCall>
	   <methodName>listDevices</methodName>
	   <params>		
	   </params>
	</methodCall>`
	reader := strings.NewReader(xmlString)
	return reader
}

func GetValueReader(id string, param string) *strings.Reader {
	xmlString := `<?xml version="1.0"?>
	<methodCall>
	   <methodName>getValue</methodName>
	   <params>
		 <param>
		   <value>{{ID}}</value>
		 </param>
		 <param>
		   <value>{{PARAM}}</value>
		 </param>
	   </params>
	</methodCall>`
	reader := strings.NewReader(strings.Replace(strings.Replace(xmlString, "{{ID}}", id, 1), "{{PARAM}}", param, 1))

	return reader
}

func GetParamsetReader(address string) *strings.Reader {
	xmlString := `<?xml version="1.0"?>
	<methodCall>
	   <methodName>getParamset</methodName>
	   <params>
		 <param>
		   <value>{{ADDRESS}}</value>
		 </param>
		 <param>
		   <value>VALUES</value>
		 </param>
	   </params>
	</methodCall>`
	reader := strings.NewReader(strings.Replace(xmlString, "{{ADDRESS}}", address, 1))

	return reader
}

func SetValueReader(address string, param string, value string) *strings.Reader {
	xmlString := `<?xml version="1.0"?>
	<methodCall>
	   <methodName>setValue</methodName>
	   <params>
		 <param>
		   <value>{{ADDRESS}}</value>
		 </param>
		 <param>
		   <value>{{PARAM}}</value>
		 </param>
		 <param>
		   <value>{{VALUE}}</value>
		 </param>
	   </params>
	</methodCall>`
	reader := strings.NewReader(strings.Replace(strings.Replace(strings.Replace(xmlString, "{{ADDRESS}}", address, 1), "{{PARAM}}", param, 1), "{{VALUE}}", value, 1))

	return reader
}
