// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1/schema/predict/prediction/text_sentiment.proto

package prediction

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Prediction output format for Text Sentiment
type TextSentimentPredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The integer sentiment labels between 0 (inclusive) and sentimentMax label
	// (inclusive), while 0 maps to the least positive sentiment and
	// sentimentMax maps to the most positive one. The higher the score is, the
	// more positive the sentiment in the text snippet is. Note: sentimentMax is
	// an integer value between 1 (inclusive) and 10 (inclusive).
	Sentiment int32 `protobuf:"varint,1,opt,name=sentiment,proto3" json:"sentiment,omitempty"`
}

func (x *TextSentimentPredictionResult) Reset() {
	*x = TextSentimentPredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSentimentPredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSentimentPredictionResult) ProtoMessage() {}

func (x *TextSentimentPredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSentimentPredictionResult.ProtoReflect.Descriptor instead.
func (*TextSentimentPredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescGZIP(), []int{0}
}

func (x *TextSentimentPredictionResult) GetSentiment() int32 {
	if x != nil {
		return x.Sentiment
	}
	return 0
}

var File_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3d, 0x0a, 0x1d, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xc0,
	0x01, 0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x22, 0x54, 0x65, 0x78,
	0x74, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_goTypes = []interface{}{
	(*TextSentimentPredictionResult)(nil), // 0: google.cloud.aiplatform.v1.schema.predict.prediction.TextSentimentPredictionResult
}
var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_init() }
func file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSentimentPredictionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto = out.File
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_sentiment_proto_depIdxs = nil
}
