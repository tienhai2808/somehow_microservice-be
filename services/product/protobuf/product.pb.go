// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: services/product/protobuf/product.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetProductBySlugRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slug          string                 `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductBySlugRequest) Reset() {
	*x = GetProductBySlugRequest{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductBySlugRequest) ProtoMessage() {}

func (x *GetProductBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductBySlugRequest.ProtoReflect.Descriptor instead.
func (*GetProductBySlugRequest) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type ProductPublicResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug          string                  `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Description   string                  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32                 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	IsSale        *bool                   `protobuf:"varint,6,opt,name=is_sale,json=isSale,proto3,oneof" json:"is_sale,omitempty"`
	SalePrice     *float32                `protobuf:"fixed32,7,opt,name=sale_price,json=salePrice,proto3,oneof" json:"sale_price,omitempty"`
	StartSale     *string                 `protobuf:"bytes,8,opt,name=start_sale,json=startSale,proto3,oneof" json:"start_sale,omitempty"`
	EndSale       *string                 `protobuf:"bytes,9,opt,name=end_sale,json=endSale,proto3,oneof" json:"end_sale,omitempty"`
	Categories    []*BaseCategoryResponse `protobuf:"bytes,10,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductPublicResponse) Reset() {
	*x = ProductPublicResponse{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductPublicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPublicResponse) ProtoMessage() {}

func (x *ProductPublicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPublicResponse.ProtoReflect.Descriptor instead.
func (*ProductPublicResponse) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductPublicResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductPublicResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductPublicResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ProductPublicResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductPublicResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductPublicResponse) GetIsSale() bool {
	if x != nil && x.IsSale != nil {
		return *x.IsSale
	}
	return false
}

func (x *ProductPublicResponse) GetSalePrice() float32 {
	if x != nil && x.SalePrice != nil {
		return *x.SalePrice
	}
	return 0
}

func (x *ProductPublicResponse) GetStartSale() string {
	if x != nil && x.StartSale != nil {
		return *x.StartSale
	}
	return ""
}

func (x *ProductPublicResponse) GetEndSale() string {
	if x != nil && x.EndSale != nil {
		return *x.EndSale
	}
	return ""
}

func (x *ProductPublicResponse) GetCategories() []*BaseCategoryResponse {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CreateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slug          *string                `protobuf:"bytes,2,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
	ParentIds     []string               `protobuf:"bytes,3,rep,name=parent_ids,json=parentIds,proto3" json:"parent_ids,omitempty"`
	UserId        string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryRequest) GetSlug() string {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return ""
}

func (x *CreateCategoryRequest) GetParentIds() []string {
	if x != nil {
		return x.ParentIds
	}
	return nil
}

func (x *CreateCategoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	IsSale        bool                   `protobuf:"varint,4,opt,name=is_sale,json=isSale,proto3" json:"is_sale,omitempty"`
	SalePrice     *float32               `protobuf:"fixed32,5,opt,name=sale_price,json=salePrice,proto3,oneof" json:"sale_price,omitempty"`
	StartSale     *string                `protobuf:"bytes,6,opt,name=start_sale,json=startSale,proto3,oneof" json:"start_sale,omitempty"`
	EndSale       *string                `protobuf:"bytes,7,opt,name=end_sale,json=endSale,proto3,oneof" json:"end_sale,omitempty"`
	CategoryIds   []string               `protobuf:"bytes,8,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
	UserId        string                 `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProductRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetIsSale() bool {
	if x != nil {
		return x.IsSale
	}
	return false
}

func (x *CreateProductRequest) GetSalePrice() float32 {
	if x != nil && x.SalePrice != nil {
		return *x.SalePrice
	}
	return 0
}

func (x *CreateProductRequest) GetStartSale() string {
	if x != nil && x.StartSale != nil {
		return *x.StartSale
	}
	return ""
}

func (x *CreateProductRequest) GetEndSale() string {
	if x != nil && x.EndSale != nil {
		return *x.EndSale
	}
	return ""
}

func (x *CreateProductRequest) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *CreateProductRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ProductAdminResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug          string                  `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Description   string                  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32                 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	IsSale        *bool                   `protobuf:"varint,6,opt,name=is_sale,json=isSale,proto3,oneof" json:"is_sale,omitempty"`
	SalePrice     *float32                `protobuf:"fixed32,7,opt,name=sale_price,json=salePrice,proto3,oneof" json:"sale_price,omitempty"`
	StartSale     *string                 `protobuf:"bytes,8,opt,name=start_sale,json=startSale,proto3,oneof" json:"start_sale,omitempty"`
	EndSale       *string                 `protobuf:"bytes,9,opt,name=end_sale,json=endSale,proto3,oneof" json:"end_sale,omitempty"`
	Categories    []*BaseCategoryResponse `protobuf:"bytes,10,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductAdminResponse) Reset() {
	*x = ProductAdminResponse{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAdminResponse) ProtoMessage() {}

func (x *ProductAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAdminResponse.ProtoReflect.Descriptor instead.
func (*ProductAdminResponse) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductAdminResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductAdminResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductAdminResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ProductAdminResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductAdminResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductAdminResponse) GetIsSale() bool {
	if x != nil && x.IsSale != nil {
		return *x.IsSale
	}
	return false
}

func (x *ProductAdminResponse) GetSalePrice() float32 {
	if x != nil && x.SalePrice != nil {
		return *x.SalePrice
	}
	return 0
}

func (x *ProductAdminResponse) GetStartSale() string {
	if x != nil && x.StartSale != nil {
		return *x.StartSale
	}
	return ""
}

func (x *ProductAdminResponse) GetEndSale() string {
	if x != nil && x.EndSale != nil {
		return *x.EndSale
	}
	return ""
}

func (x *ProductAdminResponse) GetCategories() []*BaseCategoryResponse {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CategoryAdminResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug          string                  `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Parents       []*BaseCategoryResponse `protobuf:"bytes,4,rep,name=parents,proto3" json:"parents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryAdminResponse) Reset() {
	*x = CategoryAdminResponse{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryAdminResponse) ProtoMessage() {}

func (x *CategoryAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryAdminResponse.ProtoReflect.Descriptor instead.
func (*CategoryAdminResponse) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryAdminResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryAdminResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryAdminResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CategoryAdminResponse) GetParents() []*BaseCategoryResponse {
	if x != nil {
		return x.Parents
	}
	return nil
}

type BaseCategoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug          string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseCategoryResponse) Reset() {
	*x = BaseCategoryResponse{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCategoryResponse) ProtoMessage() {}

func (x *BaseCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCategoryResponse.ProtoReflect.Descriptor instead.
func (*BaseCategoryResponse) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{6}
}

func (x *BaseCategoryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BaseCategoryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseCategoryResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type CategoryPublicResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Id            string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug          string                    `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Children      []*CategoryPublicResponse `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryPublicResponse) Reset() {
	*x = CategoryPublicResponse{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryPublicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryPublicResponse) ProtoMessage() {}

func (x *CategoryPublicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryPublicResponse.ProtoReflect.Descriptor instead.
func (*CategoryPublicResponse) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{7}
}

func (x *CategoryPublicResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryPublicResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryPublicResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CategoryPublicResponse) GetChildren() []*CategoryPublicResponse {
	if x != nil {
		return x.Children
	}
	return nil
}

type GetCategoryTreeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryTreeRequest) Reset() {
	*x = GetCategoryTreeRequest{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryTreeRequest) ProtoMessage() {}

func (x *GetCategoryTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryTreeRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryTreeRequest) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{8}
}

type CategoryTreeResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Categories    []*CategoryPublicResponse `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryTreeResponse) Reset() {
	*x = CategoryTreeResponse{}
	mi := &file_services_product_protobuf_product_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryTreeResponse) ProtoMessage() {}

func (x *CategoryTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_protobuf_product_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryTreeResponse.ProtoReflect.Descriptor instead.
func (*CategoryTreeResponse) Descriptor() ([]byte, []int) {
	return file_services_product_protobuf_product_proto_rawDescGZIP(), []int{9}
}

func (x *CategoryTreeResponse) GetCategories() []*CategoryPublicResponse {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_services_product_protobuf_product_proto protoreflect.FileDescriptor

const file_services_product_protobuf_product_proto_rawDesc = "" +
	"\n" +
	"'services/product/protobuf/product.proto\x12\aproduct\"-\n" +
	"\x17GetProductBySlugRequest\x12\x12\n" +
	"\x04slug\x18\x01 \x01(\tR\x04slug\"\x85\x03\n" +
	"\x15ProductPublicResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x12\n" +
	"\x04slug\x18\x03 \x01(\tR\x04slug\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x02R\x05price\x12\x1c\n" +
	"\ais_sale\x18\x06 \x01(\bH\x00R\x06isSale\x88\x01\x01\x12\"\n" +
	"\n" +
	"sale_price\x18\a \x01(\x02H\x01R\tsalePrice\x88\x01\x01\x12\"\n" +
	"\n" +
	"start_sale\x18\b \x01(\tH\x02R\tstartSale\x88\x01\x01\x12\x1e\n" +
	"\bend_sale\x18\t \x01(\tH\x03R\aendSale\x88\x01\x01\x12=\n" +
	"\n" +
	"categories\x18\n" +
	" \x03(\v2\x1d.product.BaseCategoryResponseR\n" +
	"categoriesB\n" +
	"\n" +
	"\b_is_saleB\r\n" +
	"\v_sale_priceB\r\n" +
	"\v_start_saleB\v\n" +
	"\t_end_sale\"\x85\x01\n" +
	"\x15CreateCategoryRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x17\n" +
	"\x04slug\x18\x02 \x01(\tH\x00R\x04slug\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"parent_ids\x18\x03 \x03(\tR\tparentIds\x12\x17\n" +
	"\auser_id\x18\x04 \x01(\tR\x06userIdB\a\n" +
	"\x05_slug\"\xcc\x02\n" +
	"\x14CreateProductRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x17\n" +
	"\ais_sale\x18\x04 \x01(\bR\x06isSale\x12\"\n" +
	"\n" +
	"sale_price\x18\x05 \x01(\x02H\x00R\tsalePrice\x88\x01\x01\x12\"\n" +
	"\n" +
	"start_sale\x18\x06 \x01(\tH\x01R\tstartSale\x88\x01\x01\x12\x1e\n" +
	"\bend_sale\x18\a \x01(\tH\x02R\aendSale\x88\x01\x01\x12!\n" +
	"\fcategory_ids\x18\b \x03(\tR\vcategoryIds\x12\x17\n" +
	"\auser_id\x18\t \x01(\tR\x06userIdB\r\n" +
	"\v_sale_priceB\r\n" +
	"\v_start_saleB\v\n" +
	"\t_end_sale\"\x84\x03\n" +
	"\x14ProductAdminResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x12\n" +
	"\x04slug\x18\x03 \x01(\tR\x04slug\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x02R\x05price\x12\x1c\n" +
	"\ais_sale\x18\x06 \x01(\bH\x00R\x06isSale\x88\x01\x01\x12\"\n" +
	"\n" +
	"sale_price\x18\a \x01(\x02H\x01R\tsalePrice\x88\x01\x01\x12\"\n" +
	"\n" +
	"start_sale\x18\b \x01(\tH\x02R\tstartSale\x88\x01\x01\x12\x1e\n" +
	"\bend_sale\x18\t \x01(\tH\x03R\aendSale\x88\x01\x01\x12=\n" +
	"\n" +
	"categories\x18\n" +
	" \x03(\v2\x1d.product.BaseCategoryResponseR\n" +
	"categoriesB\n" +
	"\n" +
	"\b_is_saleB\r\n" +
	"\v_sale_priceB\r\n" +
	"\v_start_saleB\v\n" +
	"\t_end_sale\"\x88\x01\n" +
	"\x15CategoryAdminResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x12\n" +
	"\x04slug\x18\x03 \x01(\tR\x04slug\x127\n" +
	"\aparents\x18\x04 \x03(\v2\x1d.product.BaseCategoryResponseR\aparents\"N\n" +
	"\x14BaseCategoryResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x12\n" +
	"\x04slug\x18\x03 \x01(\tR\x04slug\"\x8d\x01\n" +
	"\x16CategoryPublicResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x12\n" +
	"\x04slug\x18\x03 \x01(\tR\x04slug\x12;\n" +
	"\bchildren\x18\x04 \x03(\v2\x1f.product.CategoryPublicResponseR\bchildren\"\x18\n" +
	"\x16GetCategoryTreeRequest\"W\n" +
	"\x14CategoryTreeResponse\x12?\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\x1f.product.CategoryPublicResponseR\n" +
	"categories2\xda\x02\n" +
	"\x0eProductService\x12P\n" +
	"\x0eCreateCategory\x12\x1e.product.CreateCategoryRequest\x1a\x1e.product.CategoryAdminResponse\x12Q\n" +
	"\x0fGetCategoryTree\x12\x1f.product.GetCategoryTreeRequest\x1a\x1d.product.CategoryTreeResponse\x12M\n" +
	"\rCreateProduct\x12\x1d.product.CreateProductRequest\x1a\x1d.product.ProductAdminResponse\x12T\n" +
	"\x10GetProductBySlug\x12 .product.GetProductBySlugRequest\x1a\x1e.product.ProductPublicResponseBAZ?github.com/SomeHowMicroservice/shm-be/services/product/protobufb\x06proto3"

var (
	file_services_product_protobuf_product_proto_rawDescOnce sync.Once
	file_services_product_protobuf_product_proto_rawDescData []byte
)

func file_services_product_protobuf_product_proto_rawDescGZIP() []byte {
	file_services_product_protobuf_product_proto_rawDescOnce.Do(func() {
		file_services_product_protobuf_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_product_protobuf_product_proto_rawDesc), len(file_services_product_protobuf_product_proto_rawDesc)))
	})
	return file_services_product_protobuf_product_proto_rawDescData
}

var file_services_product_protobuf_product_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_services_product_protobuf_product_proto_goTypes = []any{
	(*GetProductBySlugRequest)(nil), // 0: product.GetProductBySlugRequest
	(*ProductPublicResponse)(nil),   // 1: product.ProductPublicResponse
	(*CreateCategoryRequest)(nil),   // 2: product.CreateCategoryRequest
	(*CreateProductRequest)(nil),    // 3: product.CreateProductRequest
	(*ProductAdminResponse)(nil),    // 4: product.ProductAdminResponse
	(*CategoryAdminResponse)(nil),   // 5: product.CategoryAdminResponse
	(*BaseCategoryResponse)(nil),    // 6: product.BaseCategoryResponse
	(*CategoryPublicResponse)(nil),  // 7: product.CategoryPublicResponse
	(*GetCategoryTreeRequest)(nil),  // 8: product.GetCategoryTreeRequest
	(*CategoryTreeResponse)(nil),    // 9: product.CategoryTreeResponse
}
var file_services_product_protobuf_product_proto_depIdxs = []int32{
	6, // 0: product.ProductPublicResponse.categories:type_name -> product.BaseCategoryResponse
	6, // 1: product.ProductAdminResponse.categories:type_name -> product.BaseCategoryResponse
	6, // 2: product.CategoryAdminResponse.parents:type_name -> product.BaseCategoryResponse
	7, // 3: product.CategoryPublicResponse.children:type_name -> product.CategoryPublicResponse
	7, // 4: product.CategoryTreeResponse.categories:type_name -> product.CategoryPublicResponse
	2, // 5: product.ProductService.CreateCategory:input_type -> product.CreateCategoryRequest
	8, // 6: product.ProductService.GetCategoryTree:input_type -> product.GetCategoryTreeRequest
	3, // 7: product.ProductService.CreateProduct:input_type -> product.CreateProductRequest
	0, // 8: product.ProductService.GetProductBySlug:input_type -> product.GetProductBySlugRequest
	5, // 9: product.ProductService.CreateCategory:output_type -> product.CategoryAdminResponse
	9, // 10: product.ProductService.GetCategoryTree:output_type -> product.CategoryTreeResponse
	4, // 11: product.ProductService.CreateProduct:output_type -> product.ProductAdminResponse
	1, // 12: product.ProductService.GetProductBySlug:output_type -> product.ProductPublicResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_services_product_protobuf_product_proto_init() }
func file_services_product_protobuf_product_proto_init() {
	if File_services_product_protobuf_product_proto != nil {
		return
	}
	file_services_product_protobuf_product_proto_msgTypes[1].OneofWrappers = []any{}
	file_services_product_protobuf_product_proto_msgTypes[2].OneofWrappers = []any{}
	file_services_product_protobuf_product_proto_msgTypes[3].OneofWrappers = []any{}
	file_services_product_protobuf_product_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_product_protobuf_product_proto_rawDesc), len(file_services_product_protobuf_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_product_protobuf_product_proto_goTypes,
		DependencyIndexes: file_services_product_protobuf_product_proto_depIdxs,
		MessageInfos:      file_services_product_protobuf_product_proto_msgTypes,
	}.Build()
	File_services_product_protobuf_product_proto = out.File
	file_services_product_protobuf_product_proto_goTypes = nil
	file_services_product_protobuf_product_proto_depIdxs = nil
}
