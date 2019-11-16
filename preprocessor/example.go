package preprocessor

import "github.com/galeone/tfgo/proto/example"
import "github.com/gogo/protobuf/proto"

// StringToFeature converts a string representation to a *example.Feature.
// This feature is the equivalent of the Python functions:
// tf.train.BytesList, tf.train.Int64List, and tf.train.FloatList.
func StringToFeature(value []string) (exampleFeature *example.Feature) {
	bytesArr := make([][]byte, 0)
	for _, s := range value {
		bytesArr = append(bytesArr, []byte(s))
	}
	bytesList := example.BytesList{Value: bytesArr}
	featureBytesList := example.Feature_BytesList{BytesList: &bytesList}
	exampleFeature = &example.Feature{Kind: &featureBytesList}
	return
}

// Float32ToFeature converts an array of float32 values to a *example.Feature.
func Float32ToFeature(value []float32) (exampleFeature *example.Feature) {
	floatList := example.FloatList{Value: value}
	featureFloatList := example.Feature_FloatList{FloatList: &floatList}
	exampleFeature = &example.Feature{Kind: &featureFloatList}
	return
}

// Int64ToFeature converts an array of float32 values to a *example.Feature.
func Int64ToFeature(value []int64) (exampleFeature *example.Feature) {
	intList := example.Int64List{Value: value}
	featureFloatList := example.Feature_Int64List{Int64List: &intList}
	exampleFeature = &example.Feature{Kind: &featureFloatList}
	return
}

// NumpyToByteArray converts a Numpy-like structure (and Python dictionary) to a
// byte array representation suitable to be used as input of a SavedModel.
func NumpyToByteArray(featureInfo map[string][]float32) (seq []byte, err error) {
	feature := make(map[string]*example.Feature)
	for k, v := range featureInfo {
		valFormat := Float32ToFeature(v)
		feature[k] = valFormat
	}
	features := example.Features{Feature: feature}
	byte_example := example.Example{Features: &features}
	seq, err = proto.Marshal(&byte_example)
	return
}

// PandasToByteArray converts a Pandas-like structure (Pandas DataFrame) to a
// byte array representation suitable to be used as input of a SavedModel.
func PandasToByteArray(featureInfo map[string]float32) (seq []byte, err error) {
	feature := make(map[string]*example.Feature)
	for k, v := range featureInfo {
		valFormat := Float32ToFeature([]float32{v})
		feature[k] = valFormat
	}
	features := example.Features{Feature: feature}
	byte_example := example.Example{Features: &features}
	seq, err = proto.Marshal(&byte_example)
	return
}
