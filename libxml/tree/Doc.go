package tree
/* 
#cgo LDFLAGS: -lxml2
#cgo CFLAGS: -I/usr/include/libxml2
#include <libxml/xmlversion.h> 
#include <libxml/parser.h> 
#include <libxml/HTMLparser.h> 
#include <libxml/HTMLtree.h> 
#include <libxml/xmlstring.h> 
#include <libxml/xpath.h> 

xmlChar *
DumpToXmlChar(xmlDoc *doc) {
  xmlChar *buff;
  int buffersize;
  xmlDocDumpFormatMemory(doc, 
                         &buff,
                         &buffersize, 1);
  return buff;
}

xmlChar *
DumpHTMLToXmlChar(xmlDoc *doc) {
  xmlChar *buff;
  int buffersize;
  htmlDocDumpMemory(doc, &buff, &buffersize);
  return buff;
}

xmlNode * GoXmlCastDocToNode(xmlDoc *doc) { return (xmlNode *)doc; }
*/
import "C"

import . "libxml/help"
import "unsafe"

type Doc struct {
	DocPtr *C.xmlDoc
	*XmlNode
}

func NewDoc(ptr unsafe.Pointer) *Doc {
	//ptr := aPtr.(*C.xmlDoc)
	doc := NewNode(C.GoXmlCastDocToNode(ptr), nil).(*Doc)
	doc.DocPtr = ptr
	return doc
}

func ParseXmlString(content string, url string, encoding string, opts int) *Doc {
	c := C.xmlCharStrdup(C.CString(content))
	c_encoding := C.CString(encoding)
	if encoding == "" {
		c_encoding = nil
	}
	xmlDocPtr := C.xmlReadDoc(c, C.CString(url), c_encoding, C.int(opts))
	return NewDoc(xmlDocPtr)
}

func XmlParse(content string) *Doc {
	return ParseXmlString(content, "", "", 0)
}

func (doc *Doc) Free() {
	C.xmlFreeDoc(doc.DocPtr)
}

func (doc *Doc) MetaEncoding() string {
	s := C.htmlGetMetaEncoding(doc.DocPtr)
	return XmlChar2String(s)
}

func (doc *Doc) Dump() string {
	return XmlChar2String(C.DumpToXmlChar(doc.DocPtr))
}

func (doc *Doc) DumpHTML() string {
	return XmlChar2String(C.DumpHTMLToXmlChar(doc.DocPtr))
}

func (doc *Doc) RootNode() Node {
	return NewNode(C.xmlDocGetRootElement(doc.DocPtr), doc)
}
