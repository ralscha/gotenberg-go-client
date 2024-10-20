package gotenberg

import (
	"strconv"
)

const (
	formFieldLandscapeOffice                 string = "landscape"
	formFieldPageRangesOffice                string = "pageRanges"
	formFieldAllowDuplicateFieldNames        string = "allowDuplicateFieldNames"
	formFieldExportBookmarks                 string = "exportBookmarks"
	formFieldExportBookmarksToPdfDestination string = "exportBookmarksToPdfDestination"
	formFieldExportPlaceholders              string = "exportPlaceholders"
	formFieldExportNotes                     string = "exportNotes"
	formFieldExportNotesPages                string = "exportNotesPages"
	formFieldExportOnlyNotesPages            string = "exportOnlyNotesPages"
	formFieldExportNotesInMargin             string = "exportNotesInMargin"
	formFieldConvertOooTargetToPdfTarget     string = "convertOooTargetToPdfTarget"
	formFieldExportLinksRelativeFsys         string = "exportLinksRelativeFsys"
	formFieldExportHiddenSlides              string = "exportHiddenSlides"
	formFieldSkipEmptyPages                  string = "skipEmptyPages"
	formFieldAddOriginalDocumentAsStream     string = "addOriginalDocumentAsStream"
	formFieldMerge                           string = "merge" // Pd473
	formFieldPdfa                            string = "pdfa" // P867e
	formFieldPdfua                           string = "pdfua" // P72e1
	formFieldLosslessImageCompression        string = "losslessImageCompression" // P2592
	formFieldQuality                         string = "quality" // P1ed9
	formFieldReduceImageResolution           string = "reduceImageResolution" // P32cd
	formFieldMaxImageResolution              string = "maxImageResolution" // P6352
)

// OfficeRequest facilitates Office documents
// conversion with the Gotenberg API.
type OfficeRequest struct {
	docs []Document

	*request
}

// NewOfficeRequest create OfficeRequest.
func NewOfficeRequest(docs ...Document) *OfficeRequest {
	return &OfficeRequest{docs, newRequest()}
}

// Landscape sets landscape form field.
func (req *OfficeRequest) Landscape(isLandscape bool) {
	req.values[formFieldLandscapeOffice] = strconv.FormatBool(isLandscape)
}

// PageRanges sets pageRanges form field.
func (req *OfficeRequest) PageRanges(ranges string) {
	req.values[formFieldPageRangesOffice] = ranges
}

func (req *OfficeRequest) Metadata(jsonData []byte) {
	req.values[formFieldMetadata] = string(jsonData)
}

// Specify whether multiple form fields exported are allowed to have the same field name.false
func (req *OfficeRequest) AllowDuplicateFieldNames() {
	req.values[formFieldAllowDuplicateFieldNames] = strconv.FormatBool(true)
}

// Specify if bookmarks are exported to PDF.
func (req *OfficeRequest) DisableExportBookmarks() {
	req.values[formFieldExportBookmarks] = strconv.FormatBool(false)
}

// Specify that the bookmarks contained in the source LibreOffice file
// should be exported to the PDF file as Named Destination.
func (req *OfficeRequest) ExportBookmarksToPdfDestination() {
	req.values[formFieldExportBookmarksToPdfDestination] = strconv.FormatBool(true)
}

// Export the placeholders fields visual markings only. The exported placeholder is ineffective.
func (req *OfficeRequest) ExportPlaceholders() {
	req.values[formFieldExportPlaceholders] = strconv.FormatBool(true)
}

// Specify if notes are exported to PDF.
func (req *OfficeRequest) ExportNotes() {
	req.values[formFieldExportNotes] = strconv.FormatBool(true)
}

// Specify if notes pages are exported to PDF. Notes pages are available in Impress documents only.
func (req *OfficeRequest) ExportNotesPages() {
	req.values[formFieldExportNotesPages] = strconv.FormatBool(true)
}

// Specify, if the form field exportNotesPages is set to true, if only notes pages are exported to PDF.
func (req *OfficeRequest) ExportOnlyNotesPages() {
	req.values[formFieldExportOnlyNotesPages] = strconv.FormatBool(true)
}

// Specify if notes in margin are exported to PDF.
func (req *OfficeRequest) ExportNotesInMargin() {
	req.values[formFieldExportNotesInMargin] = strconv.FormatBool(true)
}

// Specify that the target documents with .od[tpgs] extension, will have that extension
// changed to .pdf when the link is exported to PDF. The source document remains untouched.
func (req *OfficeRequest) ConvertOooTargetToPdfTarget() {
	req.values[formFieldConvertOooTargetToPdfTarget] = strconv.FormatBool(true)
}

// Specify that the file system related hyperlinks (file:// protocol) present in the document
// will be exported as relative to the source document location.
func (req *OfficeRequest) ExportLinksRelativeFsys() {
	req.values[formFieldExportLinksRelativeFsys] = strconv.FormatBool(true)
}

// Export, for LibreOffice Impress, slides that are not included in slide shows.
func (req *OfficeRequest) ExportHiddenSlides() {
	req.values[formFieldExportHiddenSlides] = strconv.FormatBool(true)
}

// Specify that automatically inserted empty pages are suppressed.
// This option is active only if storing Writer documents.
func (req *OfficeRequest) SkipEmptyPages() {
	req.values[formFieldSkipEmptyPages] = strconv.FormatBool(true)
}

// Specify that a stream is inserted to the PDF file which contains the original document for archiving purposes.
func (req *OfficeRequest) AddOriginalDocumentAsStream() {
	req.values[formFieldAddOriginalDocumentAsStream] = strconv.FormatBool(true)
}

// Merge sets merge form field.
func (req *OfficeRequest) Merge(merge bool) {
	req.values[formFieldMerge] = strconv.FormatBool(merge)
}

// Pdfa sets pdfa form field.
func (req *OfficeRequest) Pdfa(pdfa string) {
	req.values[formFieldPdfa] = pdfa
}

// Pdfua sets pdfua form field.
func (req *OfficeRequest) Pdfua(pdfua bool) {
	req.values[formFieldPdfua] = strconv.FormatBool(pdfua)
}

// LosslessImageCompression sets losslessImageCompression form field.
func (req *OfficeRequest) LosslessImageCompression(losslessImageCompression bool) {
	req.values[formFieldLosslessImageCompression] = strconv.FormatBool(losslessImageCompression)
}

// Quality sets quality form field.
func (req *OfficeRequest) Quality(quality int) {
	req.values[formFieldQuality] = strconv.Itoa(quality)
}

// ReduceImageResolution sets reduceImageResolution form field.
func (req *OfficeRequest) ReduceImageResolution(reduceImageResolution bool) {
	req.values[formFieldReduceImageResolution] = strconv.FormatBool(reduceImageResolution)
}

// MaxImageResolution sets maxImageResolution form field.
func (req *OfficeRequest) MaxImageResolution(maxImageResolution int) {
	req.values[formFieldMaxImageResolution] = strconv.Itoa(maxImageResolution)
}

func (req *OfficeRequest) postURL() string {
	return "/forms/libreoffice/convert"
}

func (req *OfficeRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, doc := range req.docs {
		files[doc.Filename()] = doc
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(OfficeRequest))
)
