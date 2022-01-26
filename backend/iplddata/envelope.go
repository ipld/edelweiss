package iplddata

const (
	envelopeTagKey = "___tag"
)

var envelopeTagKeyNode = Any{String(envelopeTagKey)}

func makeEnvelopeTag(tag string) KeyValue {
	return KeyValue{Key: envelopeTagKeyNode, Value: Any{String(tag)}}
}
