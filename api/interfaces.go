package api

type CompressRequest struct {
	Compression string `json:"compression"` //ZIP/TAR/RAR etc.
	Data        []byte `json:"data"`        // Provided by form multipart data
}

type CompressResponse []byte

type ListCompressionTypesRequest struct {
}
