package values

const (
	envelopeTagKey            = "___tag"
	envelopeStructureTagValue = "structure"
	envelopeCallTagValue      = "call"
)

var envelopeTagKeyNode = Any{String(envelopeTagKey)}

func makeEnvelopeTag(tag string) KeyValue {
	return KeyValue{Key: envelopeTagKeyNode, Value: Any{String(tag)}}
}

func extractEnvelopeTag(kv KeyValue) (string, bool) {
	if kv.Key == envelopeTagKeyNode {
		if s, ok := kv.Value.Value.(String); ok {
			return string(s), true
		} else {
			// invalid tag values correspond to empty string tag, which makes the envelope a generic map
			return "", true
		}
	}
	return "", false
}
