// Code generated by protoc-gen-go.
// source: customer/customer.proto
// DO NOT EDIT!

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	customer/customer.proto

It has these top-level messages:
	ReportRequest
	ReportResponse
	SetConfigRequest
	SetConfigResponse
	GetConfigRequest
	ConfigHistoryItem
	Customer
	Advertiser
	Campaign
	Creative
	Targeting
*/
package customer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReportRequest struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	FromDate string `protobuf:"bytes,2,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	ToDate   string `protobuf:"bytes,3,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
	// Available output formats: Pretty, JSON, CSV[WithNames], TabSeparated[WithNames[AndTypes]], XML (default "Pretty")
	Format string `protobuf:"bytes,4,opt,name=format" json:"format,omitempty"`
}

func (m *ReportRequest) Reset()                    { *m = ReportRequest{} }
func (m *ReportRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()               {}
func (*ReportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReportResponse struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *ReportResponse) Reset()                    { *m = ReportResponse{} }
func (m *ReportResponse) String() string            { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()               {}
func (*ReportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SetConfigRequest struct {
	Client *Customer `protobuf:"bytes,1,opt,name=client" json:"client,omitempty"`
}

func (m *SetConfigRequest) Reset()                    { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()               {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetConfigRequest) GetClient() *Customer {
	if m != nil {
		return m.Client
	}
	return nil
}

type SetConfigResponse struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *SetConfigResponse) Reset()                    { *m = SetConfigResponse{} }
func (m *SetConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*SetConfigResponse) ProtoMessage()               {}
func (*SetConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetConfigRequest struct {
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ConfigHistoryItem struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Date   string `protobuf:"bytes,2,opt,name=date" json:"date,omitempty"`
	Diff   string `protobuf:"bytes,3,opt,name=diff" json:"diff,omitempty"`
	Author string `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
}

func (m *ConfigHistoryItem) Reset()                    { *m = ConfigHistoryItem{} }
func (m *ConfigHistoryItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigHistoryItem) ProtoMessage()               {}
func (*ConfigHistoryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Customer struct {
	// There should be at least one advertiser in the config.
	Advertisers map[string]*Advertiser `protobuf:"bytes,1,rep,name=advertisers" json:"advertisers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Customer) Reset()                    { *m = Customer{} }
func (m *Customer) String() string            { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Customer) GetAdvertisers() map[string]*Advertiser {
	if m != nil {
		return m.Advertisers
	}
	return nil
}

type Advertiser struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// There should be at least one campaign for each advertiser.
	Campaigns map[string]*Campaign `protobuf:"bytes,2,rep,name=campaigns" json:"campaigns,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// There should be at least one creative for each advertiser.
	Creatives map[string]*Creative `protobuf:"bytes,3,rep,name=creatives" json:"creatives,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Advertiser) Reset()                    { *m = Advertiser{} }
func (m *Advertiser) String() string            { return proto.CompactTextString(m) }
func (*Advertiser) ProtoMessage()               {}
func (*Advertiser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Advertiser) GetCampaigns() map[string]*Campaign {
	if m != nil {
		return m.Campaigns
	}
	return nil
}

func (m *Advertiser) GetCreatives() map[string]*Creative {
	if m != nil {
		return m.Creatives
	}
	return nil
}

type Campaign struct {
	// System extention fields. Do not Edit.
	Id        string   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	ClientId  string   `protobuf:"bytes,3,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Creatives []string `protobuf:"bytes,4,rep,name=creatives" json:"creatives,omitempty"`
	// IAB content categories of the campaign. Refer to List 5.1 in the openRTB spec.
	// This is sent in the bid response as the "cat" field.
	// This is compared with the "bcat" field in the bid request. If matched - the campaign is not bidded.
	Category []string          `protobuf:"bytes,5,rep,name=category" json:"category,omitempty"`
	Pricing  *Campaign_Pricing `protobuf:"bytes,6,opt,name=pricing" json:"pricing,omitempty"`
	Budget   *Campaign_Budget  `protobuf:"bytes,7,opt,name=budget" json:"budget,omitempty"`
	// Required. If campaign has no targetings, please pass the empty object like {}.
	Targeting *Targeting `protobuf:"bytes,8,opt,name=targeting" json:"targeting,omitempty"`
	// URL of the landing page without get parameters.
	// This is sent to the "adomain" bid response field.
	// The landing page's domain is compared with the "badv" field in the bid request. If matched - the campaign is not bidded.
	LandingPageUrl []string `protobuf:"bytes,9,rep,name=landing_page_url,json=landingPageUrl" json:"landing_page_url,omitempty"`
	// The URL without the protocol prefix (http:// or https://) which will be called on the impression.
	// {RANDOM} macro could be used for cachebusting.
	// Incorrect Ex: http://www.tracker.com/?...
	// Correct Ex: wwww.tracker.com/?...
	// "http://" or "https://" will be added automatically based on the "imp.secure" flag into the bid request.
	// Note: if the impression tracker doesn't have the valid SSL certificate - please populate the "creative.insecure" flag to TRUE to avoid the stats discrepancies
	ImpressionTrackingUrl string `protobuf:"bytes,10,opt,name=impression_tracking_url,json=impressionTrackingUrl" json:"impression_tracking_url,omitempty"`
}

func (m *Campaign) Reset()                    { *m = Campaign{} }
func (m *Campaign) String() string            { return proto.CompactTextString(m) }
func (*Campaign) ProtoMessage()               {}
func (*Campaign) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Campaign) GetPricing() *Campaign_Pricing {
	if m != nil {
		return m.Pricing
	}
	return nil
}

func (m *Campaign) GetBudget() *Campaign_Budget {
	if m != nil {
		return m.Budget
	}
	return nil
}

func (m *Campaign) GetTargeting() *Targeting {
	if m != nil {
		return m.Targeting
	}
	return nil
}

// Required
type Campaign_Pricing struct {
	// This is compared against ortb field "cur" from the bid request,
	// which contains ISO-4217 alpha codes of the currencies, allowed for the bid.
	Currency string `protobuf:"bytes,1,opt,name=currency" json:"currency,omitempty"`
	// This number is sent into the bid response as the bid CPM price.
	Bid float64 `protobuf:"fixed64,2,opt,name=bid" json:"bid,omitempty"`
	// This number divided by 1,000 (CPM) the campaign budget is decreased on the impression.
	Impression float64 `protobuf:"fixed64,3,opt,name=impression" json:"impression,omitempty"`
}

func (m *Campaign_Pricing) Reset()                    { *m = Campaign_Pricing{} }
func (m *Campaign_Pricing) String() string            { return proto.CompactTextString(m) }
func (*Campaign_Pricing) ProtoMessage()               {}
func (*Campaign_Pricing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type Campaign_Budget struct {
	// Campaign daily budget limit. Daily budget is reseted every day at midnight by the UTC+3 (Moscow) time zone.
	// Please be carefull: if you set zero daily budget - that means no limit, so the campaign will spend money unstoppable.
	Daily float64 `protobuf:"fixed64,1,opt,name=daily" json:"daily,omitempty"`
}

func (m *Campaign_Budget) Reset()                    { *m = Campaign_Budget{} }
func (m *Campaign_Budget) String() string            { return proto.CompactTextString(m) }
func (*Campaign_Budget) ProtoMessage()               {}
func (*Campaign_Budget) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 1} }

type Creative struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Please ignore this field.
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// Please ignore this field.
	AdId string `protobuf:"bytes,3,opt,name=ad_id,json=adId" json:"ad_id,omitempty"`
	// Please populate this flag to TRUE in case the creative code has the "http://" requests
	// or impression_tracking_url doesn't have the valid SSL certificate.
	Insecure bool `protobuf:"varint,4,opt,name=insecure" json:"insecure,omitempty"`
	// Available creative types: "3rd_party_banner", "video".
	Type string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	// URL the user is sent to on click. Should contains the protocol (http:// or https://).
	// This is required for the creatives which should be rotated in the Google Ad Exchange.
	ClickUrl string         `protobuf:"bytes,6,opt,name=click_url,json=clickUrl" json:"click_url,omitempty"`
	Size     *Creative_Size `protobuf:"bytes,7,opt,name=size" json:"size,omitempty"`
	// Creative's html snippet. Should contains valid html code.
	// This field could contains the following macros:
	// {CLICK_URL} will be resolved to the click URL handler.
	// {CLICK_URL_ESC} will be resolved to the escaped click URL handler.
	// {PUB_ID} will be resolved to the publisher id.
	// Note: either {CLICK_URL} or {CLICK_URL_ESC} is required for the creatives which should be rotated in the Google Ad Exchange.
	HtmlContent string `protobuf:"bytes,8,opt,name=html_content,json=htmlContent" json:"html_content,omitempty"`
	// Video ad duration in seconds. Required for the creative type "video".
	// This is compared with the "video.minduration" and "video.maxduration" fields in the bid request.
	// Creative's duration should be between these values to be eligible for bid.
	Duration int64 `protobuf:"varint,9,opt,name=duration" json:"duration,omitempty"`
	// Creative's MIME type. This is compared with the "banner.mimes" field in the bid request.
	// If no matches - creative is not eligible for bid.
	Mimes []string `protobuf:"bytes,10,rep,name=mimes" json:"mimes,omitempty"`
	// Supported video protocols. Required for the creative type "video". Refer to List 5.8 in the Open RTB spec.
	// This is compared with the "video.protocols" field in the bid request.
	// If no matches - creative is not eligible for bid.
	Protocols []string `protobuf:"bytes,11,rep,name=protocols" json:"protocols,omitempty"`
	// The URL which returns the valid VAST XML.
	// Required for the creative type "video".
	VastUrl string `protobuf:"bytes,12,opt,name=vast_url,json=vastUrl" json:"vast_url,omitempty"`
}

func (m *Creative) Reset()                    { *m = Creative{} }
func (m *Creative) String() string            { return proto.CompactTextString(m) }
func (*Creative) ProtoMessage()               {}
func (*Creative) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Creative) GetSize() *Creative_Size {
	if m != nil {
		return m.Size
	}
	return nil
}

type Creative_Size struct {
	Width  int64 `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height int64 `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
}

func (m *Creative_Size) Reset()                    { *m = Creative_Size{} }
func (m *Creative_Size) String() string            { return proto.CompactTextString(m) }
func (*Creative_Size) ProtoMessage()               {}
func (*Creative_Size) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type Targeting struct {
	Geo *Targeting_Geo  `protobuf:"bytes,1,opt,name=geo" json:"geo,omitempty"`
	Os  []*Targeting_OS `protobuf:"bytes,2,rep,name=os" json:"os,omitempty"`
	// The minimum delay between any two impressions for the single user in the one ad network in seconds.
	// Works by "user.id" field in the bid request.
	Frequency int64                `protobuf:"varint,3,opt,name=frequency" json:"frequency,omitempty"`
	UserAgent *Targeting_UserAgent `protobuf:"bytes,4,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	// Carrier or ISP as it's declared in the "device.carrier" field in the bid request.
	Carrier *Targeting_ListMatch `protobuf:"bytes,5,opt,name=carrier" json:"carrier,omitempty"`
	// IAB content categories as it's declared in the "site.cat" or "app.cat" fields in the bid request.
	Cat *Targeting_ListMatch `protobuf:"bytes,6,opt,name=cat" json:"cat,omitempty"`
	// The general type of device as it's declared in the "device.devicetype" field in the bid request.
	// Refer to List 5.17 in the Open RTB spec.
	DeviceType *Targeting_ListMatch `protobuf:"bytes,7,opt,name=device_type,json=deviceType" json:"device_type,omitempty"`
	// Site's domain is extracted from the "site.page" field in the bid request and compared with the populated list.
	Domain *Targeting_ListMatch `protobuf:"bytes,8,opt,name=domain" json:"domain,omitempty"`
	// Works by the "publisher.id" field in the bid request.
	Publisher *Targeting_ListMatch `protobuf:"bytes,9,opt,name=publisher" json:"publisher,omitempty"`
	// Works by the "ext.ssp" field in the bid request.
	Ssp     *Targeting_ListMatch `protobuf:"bytes,10,opt,name=ssp" json:"ssp,omitempty"`
	Google  *Targeting_Google    `protobuf:"bytes,11,opt,name=google" json:"google,omitempty"`
	Network *Targeting_ListMatch `protobuf:"bytes,12,opt,name=network" json:"network,omitempty"`
}

func (m *Targeting) Reset()                    { *m = Targeting{} }
func (m *Targeting) String() string            { return proto.CompactTextString(m) }
func (*Targeting) ProtoMessage()               {}
func (*Targeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Targeting) GetGeo() *Targeting_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Targeting) GetOs() []*Targeting_OS {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *Targeting) GetUserAgent() *Targeting_UserAgent {
	if m != nil {
		return m.UserAgent
	}
	return nil
}

func (m *Targeting) GetCarrier() *Targeting_ListMatch {
	if m != nil {
		return m.Carrier
	}
	return nil
}

func (m *Targeting) GetCat() *Targeting_ListMatch {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *Targeting) GetDeviceType() *Targeting_ListMatch {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *Targeting) GetDomain() *Targeting_ListMatch {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *Targeting) GetPublisher() *Targeting_ListMatch {
	if m != nil {
		return m.Publisher
	}
	return nil
}

func (m *Targeting) GetSsp() *Targeting_ListMatch {
	if m != nil {
		return m.Ssp
	}
	return nil
}

func (m *Targeting) GetGoogle() *Targeting_Google {
	if m != nil {
		return m.Google
	}
	return nil
}

func (m *Targeting) GetNetwork() *Targeting_ListMatch {
	if m != nil {
		return m.Network
	}
	return nil
}

// Required for campaigns, which should be served in the Google Ad Exchange
type Targeting_Google struct {
	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId" json:"billing_id,omitempty"`
}

func (m *Targeting_Google) Reset()                    { *m = Targeting_Google{} }
func (m *Targeting_Google) String() string            { return proto.CompactTextString(m) }
func (*Targeting_Google) ProtoMessage()               {}
func (*Targeting_Google) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 0} }

type Targeting_Geo struct {
	// Country code as it's populated in the "device.geo.country" field.
	Country *Targeting_ListMatch `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	// City as it's populated in the "device.geo.city" field.
	City *Targeting_ListMatch `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
}

func (m *Targeting_Geo) Reset()                    { *m = Targeting_Geo{} }
func (m *Targeting_Geo) String() string            { return proto.CompactTextString(m) }
func (*Targeting_Geo) ProtoMessage()               {}
func (*Targeting_Geo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 1} }

func (m *Targeting_Geo) GetCountry() *Targeting_ListMatch {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Targeting_Geo) GetCity() *Targeting_ListMatch {
	if m != nil {
		return m.City
	}
	return nil
}

type Targeting_OS struct {
	// Device operating system as it's declared in the "device.os" field in the bid request.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Device operating system version as it's declared in the "device.osv" field in the bid request.
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *Targeting_OS) Reset()                    { *m = Targeting_OS{} }
func (m *Targeting_OS) String() string            { return proto.CompactTextString(m) }
func (*Targeting_OS) ProtoMessage()               {}
func (*Targeting_OS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 2} }

type Targeting_UserAgent struct {
	// Available browser names: "Opera", "Edge", "Chrome", "Safari", "Internet Explorer"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If the flag is set to TRUE, the only mobile versions of the browser's user agent are passed.
	Mobile bool `protobuf:"varint,2,opt,name=mobile" json:"mobile,omitempty"`
}

func (m *Targeting_UserAgent) Reset()                    { *m = Targeting_UserAgent{} }
func (m *Targeting_UserAgent) String() string            { return proto.CompactTextString(m) }
func (*Targeting_UserAgent) ProtoMessage()               {}
func (*Targeting_UserAgent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 3} }

type Targeting_ListMatch struct {
	Include bool     `protobuf:"varint,1,opt,name=include" json:"include,omitempty"`
	List    []string `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
}

func (m *Targeting_ListMatch) Reset()                    { *m = Targeting_ListMatch{} }
func (m *Targeting_ListMatch) String() string            { return proto.CompactTextString(m) }
func (*Targeting_ListMatch) ProtoMessage()               {}
func (*Targeting_ListMatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 4} }

func init() {
	proto.RegisterType((*ReportRequest)(nil), "customer.ReportRequest")
	proto.RegisterType((*ReportResponse)(nil), "customer.ReportResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "customer.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "customer.SetConfigResponse")
	proto.RegisterType((*GetConfigRequest)(nil), "customer.GetConfigRequest")
	proto.RegisterType((*ConfigHistoryItem)(nil), "customer.ConfigHistoryItem")
	proto.RegisterType((*Customer)(nil), "customer.Customer")
	proto.RegisterType((*Advertiser)(nil), "customer.Advertiser")
	proto.RegisterType((*Campaign)(nil), "customer.Campaign")
	proto.RegisterType((*Campaign_Pricing)(nil), "customer.Campaign.Pricing")
	proto.RegisterType((*Campaign_Budget)(nil), "customer.Campaign.Budget")
	proto.RegisterType((*Creative)(nil), "customer.Creative")
	proto.RegisterType((*Creative_Size)(nil), "customer.Creative.Size")
	proto.RegisterType((*Targeting)(nil), "customer.Targeting")
	proto.RegisterType((*Targeting_Google)(nil), "customer.Targeting.Google")
	proto.RegisterType((*Targeting_Geo)(nil), "customer.Targeting.Geo")
	proto.RegisterType((*Targeting_OS)(nil), "customer.Targeting.OS")
	proto.RegisterType((*Targeting_UserAgent)(nil), "customer.Targeting.UserAgent")
	proto.RegisterType((*Targeting_ListMatch)(nil), "customer.Targeting.ListMatch")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for API service

type APIClient interface {
	// Return clients metricks for the specified client ID with breakdown by Campaign ID.
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Customer, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := grpc.Invoke(ctx, "/customer.API/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := grpc.Invoke(ctx, "/customer.API/SetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := grpc.Invoke(ctx, "/customer.API/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Return clients metricks for the specified client ID with breakdown by Campaign ID.
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	GetConfig(context.Context, *GetConfigRequest) (*Customer, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.API/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.API/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.API/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _API_Report_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _API_SetConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _API_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("customer/customer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x6d, 0x73, 0xdb, 0x44,
	0x10, 0xae, 0x2d, 0x47, 0xb6, 0xd6, 0x25, 0x93, 0x5e, 0x4b, 0xe3, 0xaa, 0xb4, 0x03, 0x86, 0xa1,
	0x05, 0x66, 0xdc, 0x21, 0x14, 0xca, 0x5b, 0x3b, 0x13, 0x42, 0x29, 0x99, 0x81, 0x69, 0xe6, 0x92,
	0x0e, 0x1f, 0x3d, 0x8a, 0x74, 0xb1, 0x6f, 0x22, 0xe9, 0xdc, 0xd3, 0x29, 0x9d, 0xf4, 0x1b, 0xbf,
	0x86, 0xff, 0xc0, 0x4f, 0x80, 0xbf, 0xc4, 0x07, 0xf6, 0x5e, 0x24, 0x39, 0x8e, 0xa8, 0xcb, 0xb7,
	0xdd, 0xdb, 0x67, 0xf7, 0xf6, 0xe5, 0xb9, 0x17, 0xd8, 0x8e, 0xcb, 0x42, 0x89, 0x8c, 0xc9, 0x07,
	0x95, 0x30, 0x59, 0x48, 0xa1, 0x04, 0x19, 0x54, 0xfa, 0xf8, 0x35, 0xbc, 0x43, 0xd9, 0x42, 0x48,
	0x45, 0xd9, 0xcb, 0x92, 0x15, 0x8a, 0xdc, 0x86, 0x20, 0x4e, 0x39, 0xcb, 0xd5, 0x94, 0x27, 0xa3,
	0xce, 0xfb, 0x9d, 0xfb, 0x01, 0x1d, 0xd8, 0x85, 0xfd, 0x44, 0x1b, 0x4f, 0xa4, 0xc8, 0xa6, 0x49,
	0xa4, 0xd8, 0xa8, 0x6b, 0x8d, 0x7a, 0xe1, 0x47, 0xd4, 0xc9, 0x36, 0xf4, 0x95, 0xb0, 0x26, 0xcf,
	0x98, 0x7c, 0x25, 0x8c, 0xe1, 0x26, 0xf8, 0x27, 0x42, 0x66, 0x91, 0x1a, 0xf5, 0xec, 0xba, 0xd5,
	0xc6, 0x1f, 0xc1, 0x66, 0xb5, 0x77, 0xb1, 0x10, 0x79, 0xc1, 0x08, 0x81, 0x1e, 0xfa, 0x47, 0x6e,
	0x5f, 0x23, 0x8f, 0x9f, 0xc0, 0xd6, 0x21, 0x53, 0x7b, 0x22, 0x3f, 0xe1, 0xb3, 0x2a, 0xc9, 0x4f,
	0xc1, 0xb7, 0x39, 0x19, 0xe4, 0x70, 0x87, 0x4c, 0xea, 0x02, 0xf7, 0x9c, 0x40, 0x1d, 0x62, 0x7c,
	0x0f, 0xae, 0x2d, 0xf9, 0xbf, 0x61, 0x23, 0x02, 0x5b, 0xcf, 0x56, 0x36, 0x1a, 0xc7, 0x70, 0xcd,
	0x2e, 0xfc, 0xcc, 0x31, 0xac, 0x3c, 0xdf, 0x57, 0x2c, 0x23, 0x9b, 0xd0, 0xad, 0x7b, 0x83, 0x92,
	0x0b, 0x56, 0x35, 0xc4, 0xc8, 0x66, 0x8d, 0x9f, 0x9c, 0xb8, 0x4e, 0x18, 0x59, 0xf7, 0x21, 0x2a,
	0xd5, 0x5c, 0xc8, 0xaa, 0x0f, 0x56, 0x1b, 0xff, 0xd1, 0x81, 0x41, 0x95, 0x36, 0x79, 0x0a, 0xc3,
	0x28, 0x39, 0x63, 0x52, 0xf1, 0x82, 0xc9, 0x02, 0x77, 0xf1, 0xb0, 0xbe, 0x0f, 0x2f, 0xd7, 0x37,
	0xd9, 0x6d, 0x50, 0x4f, 0x73, 0x25, 0xcf, 0xe9, 0xb2, 0x5f, 0x78, 0x04, 0x5b, 0xab, 0x00, 0xb2,
	0x05, 0xde, 0x29, 0x3b, 0x77, 0x89, 0x6b, 0x11, 0xfb, 0xb8, 0x71, 0x16, 0xa5, 0xa5, 0x4d, 0x7d,
	0xb8, 0x73, 0xa3, 0xd9, 0xa6, 0x71, 0xa6, 0x16, 0xf2, 0x6d, 0xf7, 0xeb, 0xce, 0xf8, 0xaf, 0x2e,
	0x40, 0x63, 0xd1, 0x45, 0xe6, 0x51, 0xc6, 0xaa, 0x2e, 0x6a, 0x99, 0xec, 0x22, 0x7f, 0xa2, 0x6c,
	0x11, 0xf1, 0x59, 0x5e, 0x60, 0xd8, 0x95, 0xec, 0x1b, 0xe7, 0xc9, 0x5e, 0x85, 0xb2, 0xd9, 0x37,
	0x5e, 0x26, 0x84, 0x64, 0x91, 0xe2, 0x67, 0xac, 0xc0, 0x06, 0xbe, 0x21, 0x44, 0x85, 0xaa, 0x42,
	0x54, 0x7a, 0x78, 0x00, 0x9b, 0x17, 0xe3, 0xb7, 0x14, 0x7f, 0xff, 0x62, 0xf1, 0xcb, 0x1c, 0x72,
	0xae, 0x4b, 0xa5, 0x9b, 0x88, 0x17, 0xb6, 0xfb, 0x7f, 0x11, 0x9d, 0xeb, 0x72, 0x33, 0xff, 0xf1,
	0x70, 0xec, 0x6e, 0x27, 0xc7, 0xa9, 0x6e, 0xcd, 0xa9, 0x0b, 0xc7, 0xd0, 0x5b, 0x39, 0x86, 0x6d,
	0x7d, 0x7f, 0x6f, 0xb9, 0x69, 0x3d, 0x6c, 0x5a, 0xb0, 0xd4, 0x0f, 0x12, 0xc2, 0x20, 0x46, 0x5a,
	0xce, 0x90, 0xc2, 0xa3, 0x0d, 0x63, 0xac, 0x75, 0xf2, 0x10, 0xfa, 0x0b, 0xc9, 0x63, 0x9e, 0xcf,
	0x46, 0xbe, 0xc9, 0x3b, 0xbc, 0xdc, 0x89, 0xc9, 0x81, 0x45, 0xd0, 0x0a, 0x4a, 0x3e, 0x07, 0xff,
	0xb8, 0x4c, 0x66, 0x4c, 0x8d, 0xfa, 0xc6, 0xe9, 0x56, 0x8b, 0xd3, 0x0f, 0x06, 0x40, 0x1d, 0x10,
	0x5d, 0x02, 0x15, 0x49, 0x94, 0xf4, 0x56, 0x03, 0xe3, 0x75, 0xbd, 0xf1, 0x3a, 0xaa, 0x4c, 0xb4,
	0x41, 0x61, 0x47, 0xb7, 0xd2, 0x28, 0x4f, 0x50, 0x9c, 0x2e, 0xa2, 0x19, 0x9b, 0x96, 0x32, 0x1d,
	0x05, 0x26, 0xff, 0x4d, 0xb7, 0x7e, 0x80, 0xcb, 0x2f, 0x64, 0x4a, 0xbe, 0x82, 0x6d, 0x9e, 0x2d,
	0x24, 0x2b, 0x0a, 0x2e, 0xf2, 0xa9, 0x92, 0x51, 0x7c, 0xaa, 0xbd, 0xb4, 0x03, 0x98, 0x36, 0xbd,
	0xdb, 0x98, 0x8f, 0x9c, 0x15, 0xfd, 0xc2, 0xdf, 0xa0, 0xef, 0x6a, 0x33, 0x4d, 0x2a, 0xa5, 0x64,
	0x79, 0x7c, 0x5e, 0xdf, 0x7c, 0x4e, 0xd7, 0xc3, 0x3e, 0x76, 0x03, 0xea, 0x50, 0x2d, 0x92, 0xbb,
	0x00, 0x4d, 0x44, 0x33, 0xa2, 0x0e, 0x5d, 0x5a, 0x09, 0xef, 0x82, 0x6f, 0xeb, 0x27, 0x37, 0x60,
	0x23, 0x89, 0x78, 0x6a, 0x83, 0x76, 0xa8, 0x55, 0xc6, 0xbf, 0xeb, 0xf1, 0xbb, 0x01, 0xb5, 0x4e,
	0x74, 0x95, 0x12, 0xd7, 0x61, 0x23, 0x4a, 0x1a, 0x3a, 0xf4, 0xa2, 0x04, 0xa9, 0x80, 0x39, 0x73,
	0xbc, 0xd0, 0x30, 0x4f, 0x66, 0x6e, 0x95, 0x01, 0xad, 0x75, 0x1d, 0x54, 0x9d, 0x2f, 0x18, 0x0e,
	0xdc, 0xe0, 0xb5, 0xec, 0x78, 0x15, 0x9f, 0x9a, 0xc6, 0xf8, 0x35, 0xaf, 0xe2, 0x53, 0xdd, 0xc3,
	0xcf, 0xa0, 0x57, 0xf0, 0xd7, 0xcc, 0x4d, 0x74, 0xfb, 0x32, 0x7d, 0x27, 0x87, 0x68, 0xa6, 0x06,
	0x44, 0x3e, 0x80, 0xab, 0x73, 0x95, 0xa5, 0xd3, 0x58, 0xe4, 0x4a, 0xdf, 0xc4, 0x03, 0x13, 0x6c,
	0xa8, 0xd7, 0xf6, 0xec, 0x92, 0x4e, 0x2e, 0x29, 0x25, 0x7a, 0x62, 0x83, 0x02, 0x34, 0x7b, 0xb4,
	0xd6, 0x75, 0x53, 0x32, 0x9e, 0x21, 0x57, 0xc1, 0x8c, 0xd3, 0x2a, 0x9a, 0xc5, 0xe6, 0x85, 0x8a,
	0x45, 0x5a, 0x8c, 0x86, 0x96, 0xc5, 0xf5, 0x02, 0xb9, 0x05, 0x83, 0xb3, 0xa8, 0x50, 0x26, 0xf7,
	0xab, 0x66, 0xbb, 0xbe, 0xd6, 0xf5, 0x18, 0x1f, 0x42, 0x4f, 0xe7, 0xa6, 0xc3, 0xbe, 0xe2, 0x89,
	0x9a, 0x9b, 0x4e, 0x7a, 0xd4, 0x2a, 0xfa, 0xe6, 0x9d, 0x33, 0x3e, 0x9b, 0x2b, 0xd3, 0x4e, 0x8f,
	0x3a, 0x6d, 0xfc, 0x67, 0x1f, 0x82, 0x9a, 0x77, 0xe4, 0x13, 0xf0, 0x66, 0x4c, 0xb8, 0x27, 0x65,
	0xbb, 0x85, 0x99, 0x93, 0x67, 0x4c, 0x50, 0x8d, 0x21, 0x1f, 0x43, 0x57, 0x54, 0xd7, 0xdb, 0xcd,
	0x36, 0xe4, 0xf3, 0x43, 0x8a, 0x08, 0x5d, 0xcf, 0x89, 0xd4, 0x6f, 0x89, 0xe6, 0x94, 0x67, 0xf6,
	0x6e, 0x16, 0xc8, 0xf7, 0x00, 0x25, 0xde, 0x63, 0x53, 0xa4, 0x70, 0x6e, 0x1f, 0xc7, 0xe1, 0xce,
	0x9d, 0xb6, 0x68, 0x2f, 0x10, 0xb5, 0xab, 0x41, 0x34, 0x28, 0x2b, 0x91, 0x3c, 0x82, 0x7e, 0x1c,
	0x49, 0xc9, 0x99, 0x34, 0x13, 0xfe, 0x0f, 0xd7, 0x5f, 0xf0, 0xe1, 0xfa, 0x35, 0x52, 0xf1, 0x9c,
	0x56, 0x68, 0xf2, 0x00, 0x3c, 0x3c, 0xfc, 0xee, 0xb0, 0xaf, 0x71, 0xd2, 0x48, 0xf2, 0x04, 0x86,
	0x09, 0x3b, 0xe3, 0x31, 0x9b, 0x1a, 0x3e, 0xf5, 0xdf, 0xc6, 0x11, 0xac, 0xc7, 0x91, 0x26, 0xdd,
	0x97, 0xe0, 0x27, 0x22, 0x8b, 0x78, 0xee, 0x4e, 0xfd, 0x1a, 0x57, 0x07, 0x26, 0xdf, 0x21, 0x19,
	0xca, 0xe3, 0x94, 0x17, 0x73, 0x2c, 0x31, 0x78, 0x1b, 0xcf, 0x06, 0xaf, 0x8b, 0x2c, 0x8a, 0x85,
	0x39, 0xfb, 0xeb, 0x8b, 0x44, 0x24, 0xd9, 0x01, 0x7f, 0x26, 0xc4, 0x2c, 0x65, 0xc8, 0xbb, 0x95,
	0x5b, 0x70, 0x89, 0x00, 0x06, 0x41, 0x1d, 0x52, 0x8f, 0x20, 0x67, 0xea, 0x95, 0x90, 0xa7, 0x86,
	0x8f, 0xeb, 0x47, 0xe0, 0xd0, 0xe1, 0x3d, 0xf0, 0x6d, 0x28, 0x72, 0x07, 0xe0, 0x98, 0xa7, 0xa9,
	0xbe, 0xab, 0xea, 0x4f, 0x45, 0xe0, 0x56, 0xf6, 0x93, 0xf0, 0x25, 0x78, 0x48, 0x3a, 0x33, 0x6b,
	0x51, 0xea, 0x67, 0xc7, 0xd1, 0x73, 0xed, 0xac, 0x2d, 0x1a, 0xef, 0xdc, 0x5e, 0xcc, 0xd5, 0xb9,
	0x7b, 0x91, 0xd6, 0x78, 0x19, 0x68, 0xb8, 0x03, 0xdd, 0xe7, 0x87, 0xad, 0x37, 0xd2, 0x08, 0xfa,
	0xf8, 0xf6, 0x9a, 0xfb, 0xae, 0xeb, 0x8e, 0x9f, 0x55, 0xc3, 0x47, 0x10, 0xd4, 0x1c, 0x6d, 0x75,
	0xc5, 0x13, 0x98, 0x09, 0x2c, 0xcb, 0xbe, 0x8d, 0x03, 0xea, 0xb4, 0xf0, 0x1b, 0x08, 0xea, 0xfd,
	0x75, 0x7c, 0x9e, 0xc7, 0x69, 0x99, 0x58, 0xdf, 0x01, 0xad, 0x54, 0x1d, 0x12, 0xe7, 0xaa, 0xcc,
	0x89, 0xc3, 0x90, 0x5a, 0xde, 0xf9, 0xbb, 0x03, 0xde, 0xee, 0xc1, 0x3e, 0x79, 0x0c, 0xbe, 0xfd,
	0x46, 0x92, 0xa5, 0x33, 0x7b, 0xe1, 0x53, 0x1b, 0x8e, 0x2e, 0x1b, 0xec, 0x47, 0x70, 0x7c, 0x85,
	0xfc, 0x04, 0x41, 0xfd, 0x3f, 0x24, 0x4b, 0x43, 0x5f, 0xfd, 0x74, 0x86, 0xb7, 0x5b, 0x6d, 0x75,
	0x9c, 0xc7, 0x10, 0x3c, 0x6b, 0x8b, 0xb3, 0xfa, 0xa7, 0x0c, 0x5b, 0x3e, 0xab, 0xe3, 0x2b, 0xc7,
	0xbe, 0xb9, 0xe6, 0xbe, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xd3, 0x37, 0xc0, 0xb4, 0x0b,
	0x00, 0x00,
}
