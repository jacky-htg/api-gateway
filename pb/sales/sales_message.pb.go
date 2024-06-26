// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: sales/sales_message.proto

package sales

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sales struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                       string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId                 string         `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	BranchName               string         `protobuf:"bytes,3,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	Customer                 *Customer      `protobuf:"bytes,4,opt,name=customer,proto3" json:"customer,omitempty"`
	Salesman                 *Salesman      `protobuf:"bytes,5,opt,name=salesman,proto3" json:"salesman,omitempty"`
	Code                     string         `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	SalesDate                string         `protobuf:"bytes,7,opt,name=sales_date,json=salesDate,proto3" json:"sales_date,omitempty"`
	Price                    float64        `protobuf:"fixed64,8,opt,name=price,proto3" json:"price,omitempty"`
	AdditionalDiscAmount     float64        `protobuf:"fixed64,9,opt,name=additional_disc_amount,json=additionalDiscAmount,proto3" json:"additional_disc_amount,omitempty"`
	AdditionalDiscPercentage float32        `protobuf:"fixed32,10,opt,name=additional_disc_percentage,json=additionalDiscPercentage,proto3" json:"additional_disc_percentage,omitempty"`
	TotalPrice               float64        `protobuf:"fixed64,11,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Remark                   string         `protobuf:"bytes,12,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt                string         `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy                string         `protobuf:"bytes,14,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt                string         `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy                string         `protobuf:"bytes,16,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Details                  []*SalesDetail `protobuf:"bytes,17,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Sales) Reset() {
	*x = Sales{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_sales_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sales) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sales) ProtoMessage() {}

func (x *Sales) ProtoReflect() protoreflect.Message {
	mi := &file_sales_sales_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sales.ProtoReflect.Descriptor instead.
func (*Sales) Descriptor() ([]byte, []int) {
	return file_sales_sales_message_proto_rawDescGZIP(), []int{0}
}

func (x *Sales) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sales) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Sales) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *Sales) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Sales) GetSalesman() *Salesman {
	if x != nil {
		return x.Salesman
	}
	return nil
}

func (x *Sales) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Sales) GetSalesDate() string {
	if x != nil {
		return x.SalesDate
	}
	return ""
}

func (x *Sales) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Sales) GetAdditionalDiscAmount() float64 {
	if x != nil {
		return x.AdditionalDiscAmount
	}
	return 0
}

func (x *Sales) GetAdditionalDiscPercentage() float32 {
	if x != nil {
		return x.AdditionalDiscPercentage
	}
	return 0
}

func (x *Sales) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Sales) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Sales) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Sales) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Sales) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Sales) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Sales) GetDetails() []*SalesDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

type SalesDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SalesId        string  `protobuf:"bytes,2,opt,name=sales_id,json=salesId,proto3" json:"sales_id,omitempty"`
	ProductId      string  `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductCode    string  `protobuf:"bytes,4,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	ProductName    string  `protobuf:"bytes,5,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Quantity       int32   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price          float64 `protobuf:"fixed64,7,opt,name=price,proto3" json:"price,omitempty"`
	DiscAmount     float64 `protobuf:"fixed64,8,opt,name=disc_amount,json=discAmount,proto3" json:"disc_amount,omitempty"`
	DiscPercentage float32 `protobuf:"fixed32,9,opt,name=disc_percentage,json=discPercentage,proto3" json:"disc_percentage,omitempty"`
	TotalPrice     float64 `protobuf:"fixed64,10,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *SalesDetail) Reset() {
	*x = SalesDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_sales_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalesDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalesDetail) ProtoMessage() {}

func (x *SalesDetail) ProtoReflect() protoreflect.Message {
	mi := &file_sales_sales_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalesDetail.ProtoReflect.Descriptor instead.
func (*SalesDetail) Descriptor() ([]byte, []int) {
	return file_sales_sales_message_proto_rawDescGZIP(), []int{1}
}

func (x *SalesDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SalesDetail) GetSalesId() string {
	if x != nil {
		return x.SalesId
	}
	return ""
}

func (x *SalesDetail) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *SalesDetail) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *SalesDetail) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SalesDetail) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SalesDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SalesDetail) GetDiscAmount() float64 {
	if x != nil {
		return x.DiscAmount
	}
	return 0
}

func (x *SalesDetail) GetDiscPercentage() float32 {
	if x != nil {
		return x.DiscPercentage
	}
	return 0
}

func (x *SalesDetail) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

var File_sales_sales_message_proto protoreflect.FileDescriptor

var file_sales_sales_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x1a, 0x1c, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf,
	0x04, 0x0a, 0x05, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3c, 0x0a, 0x1a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x18, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44,
	0x69, 0x73, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0xba, 0x02, 0x0a, 0x0b, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x35, 0x0a,
	0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72,
	0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x50, 0x01, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x3b, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sales_sales_message_proto_rawDescOnce sync.Once
	file_sales_sales_message_proto_rawDescData = file_sales_sales_message_proto_rawDesc
)

func file_sales_sales_message_proto_rawDescGZIP() []byte {
	file_sales_sales_message_proto_rawDescOnce.Do(func() {
		file_sales_sales_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_sales_sales_message_proto_rawDescData)
	})
	return file_sales_sales_message_proto_rawDescData
}

var file_sales_sales_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sales_sales_message_proto_goTypes = []interface{}{
	(*Sales)(nil),       // 0: sales.Sales
	(*SalesDetail)(nil), // 1: sales.SalesDetail
	(*Customer)(nil),    // 2: sales.Customer
	(*Salesman)(nil),    // 3: sales.Salesman
}
var file_sales_sales_message_proto_depIdxs = []int32{
	2, // 0: sales.Sales.customer:type_name -> sales.Customer
	3, // 1: sales.Sales.salesman:type_name -> sales.Salesman
	1, // 2: sales.Sales.details:type_name -> sales.SalesDetail
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sales_sales_message_proto_init() }
func file_sales_sales_message_proto_init() {
	if File_sales_sales_message_proto != nil {
		return
	}
	file_sales_customer_message_proto_init()
	file_sales_salesman_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sales_sales_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sales); i {
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
		file_sales_sales_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalesDetail); i {
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
			RawDescriptor: file_sales_sales_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sales_sales_message_proto_goTypes,
		DependencyIndexes: file_sales_sales_message_proto_depIdxs,
		MessageInfos:      file_sales_sales_message_proto_msgTypes,
	}.Build()
	File_sales_sales_message_proto = out.File
	file_sales_sales_message_proto_rawDesc = nil
	file_sales_sales_message_proto_goTypes = nil
	file_sales_sales_message_proto_depIdxs = nil
}