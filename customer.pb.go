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
	// From date in the 'YYYY-MM-DD' format
	FromDate string `protobuf:"bytes,2,opt,name=from_date,json=fromDate" json:"from_date,omitempty"`
	// To date in the 'YYYY-MM-DD' format
	ToDate string `protobuf:"bytes,3,opt,name=to_date,json=toDate" json:"to_date,omitempty"`
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
	// Campaign start datetime in the YYYY-MM-DDThh:mm:ss±hh:mm format.
	// The date could be ommited (Thh:mm:ss±hh:mm format). In that case campaign is started daily at this time.
	StartDate string `protobuf:"bytes,11,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	// Campaign stop datetime in the YYYY-MM-DDThh:mm:ss±hh:mm format.
	// The date could be ommited (Thh:mm:ss±hh:mm format). In that case campaign is stopped daily at this time.
	StopDate string `protobuf:"bytes,12,opt,name=stop_date,json=stopDate" json:"stop_date,omitempty"`
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
	// Please note that some ad networks (google for example) require all the served content to be secured.
	// For those networks you will not receive bids for the creatives with the insecure flag populated.
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
	Ssp    *Targeting_ListMatch `protobuf:"bytes,10,opt,name=ssp" json:"ssp,omitempty"`
	Google *Targeting_Google    `protobuf:"bytes,11,opt,name=google" json:"google,omitempty"`
	// Connected ad networks targeting. Please ask your account manager to obtain the available networks list.
	Network *Targeting_ListMatch `protobuf:"bytes,12,opt,name=network" json:"network,omitempty"`
	// Works by the "device.ip" field in the bid request.
	Ip      *Targeting_ListMatch `protobuf:"bytes,13,opt,name=ip" json:"ip,omitempty"`
	RangeIp []string             `protobuf:"bytes,14,rep,name=range_ip,json=rangeIp" json:"range_ip,omitempty"`
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

func (m *Targeting) GetIp() *Targeting_ListMatch {
	if m != nil {
		return m.Ip
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
	// Country code using ISO-3166-1-alpha-3.
	// Works by the "device.geo.country" field in the bid request.
	Country *Targeting_ListMatch `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	// City using United Nations Code for Trade & Transport Locations.
	// Works by the "device.geo.city" field in the bid request.
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
	// 1246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x6f, 0x6f, 0xdb, 0xb6,
	0x13, 0xae, 0x2d, 0x47, 0x96, 0xce, 0x6d, 0x90, 0xb2, 0xfd, 0x35, 0xaa, 0xfa, 0x6b, 0x91, 0x79,
	0xc3, 0x9a, 0x75, 0x98, 0x8b, 0x66, 0xdd, 0xba, 0x7f, 0x2d, 0x90, 0x65, 0x5d, 0x16, 0x60, 0x43,
	0x03, 0x26, 0xc5, 0x5e, 0x1a, 0x8c, 0xc4, 0xc8, 0x44, 0x64, 0x51, 0x25, 0xa9, 0x14, 0xce, 0xbb,
	0x01, 0xfb, 0x2e, 0xfb, 0x0a, 0xfb, 0x0c, 0xdb, 0x97, 0x1a, 0xf8, 0x47, 0x92, 0xe3, 0x78, 0x75,
	0xf7, 0x8e, 0x77, 0xf7, 0x3c, 0xc7, 0xe3, 0xe9, 0xe1, 0x89, 0xb0, 0x99, 0x54, 0x52, 0xf1, 0x29,
	0x15, 0x8f, 0xeb, 0xc5, 0xa8, 0x14, 0x5c, 0x71, 0x14, 0xd4, 0xf6, 0xf0, 0x02, 0x6e, 0x60, 0x5a,
	0x72, 0xa1, 0x30, 0x7d, 0x53, 0x51, 0xa9, 0xd0, 0x3d, 0x08, 0x93, 0x9c, 0xd1, 0x42, 0x8d, 0x59,
	0x1a, 0x75, 0xb6, 0x3a, 0xdb, 0x21, 0x0e, 0xac, 0xe3, 0x20, 0xd5, 0xc1, 0x53, 0xc1, 0xa7, 0xe3,
	0x94, 0x28, 0x1a, 0x75, 0x6d, 0x50, 0x3b, 0x7e, 0x20, 0x8a, 0xa2, 0x4d, 0xe8, 0x2b, 0x6e, 0x43,
	0x9e, 0x09, 0xf9, 0x8a, 0x9b, 0xc0, 0x1d, 0xf0, 0x4f, 0xb9, 0x98, 0x12, 0x15, 0xf5, 0xac, 0xdf,
	0x5a, 0xc3, 0x8f, 0x60, 0xbd, 0xde, 0x5b, 0x96, 0xbc, 0x90, 0x14, 0x21, 0xe8, 0xa5, 0x44, 0x11,
	0xb7, 0xaf, 0x59, 0x0f, 0x5f, 0xc0, 0xc6, 0x11, 0x55, 0x7b, 0xbc, 0x38, 0x65, 0x59, 0x5d, 0xe4,
	0x23, 0xf0, 0x6d, 0x4d, 0x06, 0x39, 0xd8, 0x41, 0xa3, 0xe6, 0x80, 0x7b, 0x6e, 0x81, 0x1d, 0x62,
	0xf8, 0x10, 0x6e, 0xce, 0xf1, 0xdf, 0xb1, 0x11, 0x82, 0x8d, 0xfd, 0x85, 0x8d, 0x86, 0x09, 0xdc,
	0xb4, 0x8e, 0x9f, 0x98, 0x54, 0x5c, 0xcc, 0x0e, 0x14, 0x9d, 0xa2, 0x75, 0xe8, 0x36, 0xbd, 0xe9,
	0xb2, 0xd4, 0x25, 0xab, 0x1b, 0x62, 0xd6, 0xc6, 0xc7, 0x4e, 0x4f, 0x5d, 0x27, 0xcc, 0x5a, 0xf7,
	0x81, 0x54, 0x6a, 0xc2, 0x45, 0xdd, 0x07, 0x6b, 0x0d, 0xff, 0xe8, 0x40, 0x50, 0x97, 0x8d, 0x5e,
	0xc2, 0x80, 0xa4, 0xe7, 0x54, 0x28, 0x26, 0xa9, 0x90, 0x51, 0x67, 0xcb, 0xdb, 0x1e, 0xec, 0x7c,
	0x78, 0xf5, 0x7c, 0xa3, 0xdd, 0x16, 0xf5, 0xb2, 0x50, 0x62, 0x86, 0xe7, 0x79, 0xf1, 0x31, 0x6c,
	0x2c, 0x02, 0xd0, 0x06, 0x78, 0x67, 0x74, 0xe6, 0x0a, 0xd7, 0x4b, 0xf4, 0x08, 0xd6, 0xce, 0x49,
	0x5e, 0xd9, 0xd2, 0x07, 0x3b, 0xb7, 0xdb, 0x6d, 0x5a, 0x32, 0xb6, 0x90, 0x6f, 0xba, 0x5f, 0x75,
	0x86, 0x7f, 0x75, 0x01, 0xda, 0x88, 0x3e, 0x64, 0x41, 0xa6, 0xb4, 0xee, 0xa2, 0x5e, 0xa3, 0x5d,
	0x08, 0x13, 0x32, 0x2d, 0x09, 0xcb, 0x0a, 0x19, 0x75, 0x17, 0xab, 0x6f, 0xc9, 0xa3, 0xbd, 0x1a,
	0x65, 0xab, 0x6f, 0x59, 0x26, 0x85, 0xa0, 0x44, 0xb1, 0x73, 0x2a, 0x23, 0xef, 0x5d, 0x29, 0x6a,
	0x54, 0x9d, 0xa2, 0xb6, 0xe3, 0x43, 0x58, 0xbf, 0x9c, 0x7f, 0xc9, 0xe1, 0xb7, 0x2f, 0x1f, 0x7e,
	0x5e, 0x43, 0x8e, 0x3a, 0x77, 0x74, 0x93, 0xf1, 0xd2, 0x76, 0xff, 0x2d, 0xa3, 0xa3, 0xce, 0x37,
	0xf3, 0xcf, 0x1e, 0x04, 0xf5, 0x4e, 0x4e, 0x53, 0xdd, 0x46, 0x53, 0x97, 0xae, 0xa1, 0xb7, 0x70,
	0x0d, 0x97, 0xf5, 0xfd, 0xff, 0xf3, 0x4d, 0xeb, 0x6d, 0x79, 0xdb, 0xe1, 0x5c, 0x3f, 0x50, 0x0c,
	0x41, 0x42, 0x14, 0xcd, 0xb8, 0x98, 0x45, 0x6b, 0x26, 0xd8, 0xd8, 0xe8, 0x29, 0xf4, 0x4b, 0xc1,
	0x12, 0x56, 0x64, 0x91, 0x6f, 0xea, 0x8e, 0xaf, 0x76, 0x62, 0x74, 0x68, 0x11, 0xb8, 0x86, 0xa2,
	0x27, 0xe0, 0x9f, 0x54, 0x69, 0x46, 0x55, 0xd4, 0x37, 0xa4, 0xbb, 0x4b, 0x48, 0xdf, 0x1b, 0x00,
	0x76, 0x40, 0xf4, 0x04, 0x42, 0x45, 0x44, 0x46, 0x95, 0xde, 0x2a, 0x30, 0xac, 0x5b, 0x2d, 0xeb,
	0xb8, 0x0e, 0xe1, 0x16, 0x85, 0xb6, 0x61, 0x23, 0x27, 0x45, 0xca, 0x8a, 0x6c, 0x5c, 0x92, 0x8c,
	0x8e, 0x2b, 0x91, 0x47, 0xa1, 0xa9, 0x7f, 0xdd, 0xf9, 0x0f, 0x49, 0x46, 0x5f, 0x8b, 0x1c, 0x7d,
	0x09, 0x9b, 0x6c, 0x5a, 0x0a, 0x2a, 0x25, 0xe3, 0xc5, 0x58, 0x09, 0x92, 0x9c, 0x69, 0x96, 0x26,
	0x80, 0x69, 0xd3, 0xff, 0xda, 0xf0, 0xb1, 0x8b, 0x6a, 0xde, 0x7d, 0x00, 0xa9, 0x88, 0x50, 0x76,
	0x70, 0x0d, 0x0c, 0x34, 0x34, 0x1e, 0x33, 0xbb, 0xee, 0x41, 0x28, 0x15, 0x2f, 0x6d, 0xf4, 0xba,
	0xfd, 0x0e, 0xda, 0xa1, 0x83, 0xf1, 0xaf, 0xd0, 0x77, 0x7d, 0x31, 0x0d, 0xae, 0x84, 0xa0, 0x45,
	0x32, 0x6b, 0xa6, 0xa6, 0xb3, 0xb5, 0x50, 0x4e, 0xdc, 0xc7, 0xed, 0x60, 0xbd, 0x44, 0x0f, 0x00,
	0xda, 0x6a, 0xcc, 0xe7, 0xed, 0xe0, 0x39, 0x4f, 0xfc, 0x00, 0x7c, 0xdb, 0x3b, 0x74, 0x1b, 0xd6,
	0x52, 0xc2, 0x72, 0x9b, 0xb4, 0x83, 0xad, 0x31, 0xfc, 0xcd, 0x83, 0xa0, 0x96, 0xd4, 0x52, 0x35,
	0x2c, 0xca, 0xe9, 0x16, 0xac, 0x91, 0xb4, 0x95, 0x52, 0x8f, 0xa4, 0x07, 0xa9, 0xae, 0x99, 0x15,
	0x92, 0x26, 0x95, 0xa0, 0x66, 0x22, 0x05, 0xb8, 0xb1, 0x75, 0x52, 0x35, 0x2b, 0x69, 0xb4, 0x66,
	0xf1, 0x7a, 0xed, 0x34, 0x99, 0x9c, 0x99, 0xa6, 0xfa, 0x8d, 0x26, 0x93, 0x33, 0xdd, 0xc7, 0x4f,
	0xa1, 0x27, 0xd9, 0x05, 0x75, 0x6a, 0xd8, 0xbc, 0x2a, 0xfd, 0xd1, 0x11, 0xbb, 0xa0, 0xd8, 0x80,
	0xd0, 0x07, 0x70, 0x7d, 0xa2, 0xa6, 0xf9, 0x38, 0xe1, 0x85, 0xd2, 0x53, 0x3c, 0x30, 0xc9, 0x06,
	0xda, 0xb7, 0x67, 0x5d, 0xba, 0xb8, 0xb4, 0x12, 0x44, 0xe9, 0x06, 0x85, 0x5b, 0x9d, 0x6d, 0x0f,
	0x37, 0xb6, 0x6e, 0xca, 0x94, 0x4d, 0xa9, 0x8c, 0xc0, 0x48, 0xc1, 0x1a, 0xfa, 0x06, 0x98, 0xbf,
	0x5b, 0xc2, 0x73, 0x19, 0x0d, 0xec, 0x0d, 0x68, 0x1c, 0xe8, 0x2e, 0x04, 0xe7, 0x44, 0x2a, 0x53,
	0xbb, 0xfd, 0x8e, 0x7d, 0x6d, 0xbf, 0x16, 0x79, 0xfc, 0x14, 0x7a, 0xba, 0x36, 0x9d, 0xf6, 0x2d,
	0x4b, 0xd5, 0xc4, 0x74, 0xd2, 0xc3, 0xd6, 0xd0, 0x53, 0x7b, 0x42, 0x59, 0x36, 0x51, 0xa6, 0x9d,
	0x1e, 0x76, 0xd6, 0xf0, 0xf7, 0x00, 0xc2, 0x46, 0xb3, 0xe8, 0x13, 0xf0, 0x32, 0xca, 0xdd, 0xef,
	0x68, 0x73, 0x89, 0xaa, 0x47, 0xfb, 0x94, 0x63, 0x8d, 0x41, 0x1f, 0x43, 0x97, 0xd7, 0xa3, 0xf1,
	0xce, 0x32, 0xe4, 0xab, 0x23, 0xdc, 0xe5, 0xe6, 0x3c, 0xa7, 0x42, 0xff, 0x87, 0xb4, 0xa6, 0x3c,
	0xb3, 0x77, 0xeb, 0x40, 0xdf, 0x01, 0x54, 0x92, 0x8a, 0x31, 0xc9, 0x74, 0x03, 0x7b, 0x66, 0xdf,
	0xfb, 0xcb, 0xb2, 0xbd, 0x96, 0x54, 0xec, 0x6a, 0x10, 0x0e, 0xab, 0x7a, 0x89, 0x9e, 0x41, 0x3f,
	0x21, 0x42, 0x30, 0x2a, 0xcc, 0x17, 0xfe, 0x17, 0xea, 0xcf, 0x4c, 0xaa, 0x5f, 0x88, 0x4a, 0x26,
	0xb8, 0x46, 0xa3, 0xc7, 0xe0, 0x25, 0x44, 0xb9, 0x41, 0xb1, 0x82, 0xa4, 0x91, 0xe8, 0x05, 0x0c,
	0x52, 0x7a, 0xce, 0x12, 0x3a, 0x36, 0x7a, 0xea, 0xbf, 0x0f, 0x11, 0x2c, 0xe3, 0x58, 0x8b, 0xee,
	0x0b, 0xf0, 0x53, 0x3e, 0x25, 0xac, 0x70, 0x13, 0x63, 0x05, 0xd5, 0x81, 0xd1, 0xb7, 0x10, 0x96,
	0xd5, 0x49, 0xce, 0xe4, 0x84, 0x0a, 0xa3, 0x9f, 0x95, 0xcc, 0x16, 0xaf, 0x0f, 0x29, 0x65, 0x69,
	0xe6, 0xc6, 0xea, 0x43, 0x4a, 0x59, 0xa2, 0x1d, 0xf0, 0x33, 0xce, 0xb3, 0xdc, 0x0e, 0x90, 0x4b,
	0x13, 0x74, 0x4e, 0x00, 0x06, 0x81, 0x1d, 0x52, 0x7f, 0x82, 0x82, 0xaa, 0xb7, 0x5c, 0x9c, 0x19,
	0x3d, 0xae, 0xfe, 0x04, 0x0e, 0x8d, 0x3e, 0x83, 0x2e, 0x2b, 0xa3, 0x1b, 0xef, 0xc3, 0xe9, 0xb2,
	0x52, 0x0b, 0x5f, 0x90, 0x22, 0xa3, 0x63, 0x56, 0x46, 0xeb, 0xe6, 0x56, 0xf4, 0x8d, 0x7d, 0x50,
	0xc6, 0x0f, 0xc1, 0xb7, 0x45, 0xe9, 0x29, 0x78, 0xc2, 0xf2, 0x5c, 0x4f, 0xcc, 0xe6, 0x69, 0x13,
	0x3a, 0xcf, 0x41, 0x1a, 0xbf, 0x01, 0x6f, 0x9f, 0x72, 0xa3, 0x1a, 0x5e, 0xe9, 0x9f, 0x9f, 0x13,
	0xfa, 0x4a, 0xd5, 0x58, 0x34, 0x7a, 0x02, 0xbd, 0x84, 0xa9, 0x99, 0xfb, 0x2f, 0xae, 0x60, 0x19,
	0x68, 0xbc, 0x03, 0xdd, 0x57, 0x47, 0x4b, 0x67, 0x5b, 0x04, 0xfd, 0x73, 0x2a, 0xcc, 0xe4, 0xec,
	0xba, 0x8b, 0x6c, 0xcd, 0xf8, 0x19, 0x84, 0x8d, 0xda, 0x97, 0x52, 0xef, 0x80, 0x3f, 0xe5, 0x27,
	0x2c, 0xb7, 0x7f, 0xe8, 0x00, 0x3b, 0x2b, 0xfe, 0x1a, 0xc2, 0x66, 0x7f, 0x9d, 0x9f, 0x15, 0x49,
	0x5e, 0xa5, 0x96, 0x1b, 0xe0, 0xda, 0xd4, 0x29, 0x73, 0x26, 0x95, 0xb9, 0xbb, 0x21, 0x36, 0xeb,
	0x9d, 0xbf, 0x3b, 0xe0, 0xed, 0x1e, 0x1e, 0xa0, 0xe7, 0xe0, 0xdb, 0xc7, 0x2c, 0x9a, 0xbb, 0xfd,
	0x97, 0x9e, 0xd6, 0x71, 0x74, 0x35, 0x60, 0x9f, 0xa3, 0xc3, 0x6b, 0xe8, 0x47, 0x08, 0x9b, 0x57,
	0x2a, 0x9a, 0x93, 0xcf, 0xe2, 0xd3, 0x37, 0xbe, 0xb7, 0x34, 0xd6, 0xe4, 0x79, 0x0e, 0xe1, 0xfe,
	0xb2, 0x3c, 0x8b, 0x2f, 0xdb, 0x78, 0xc9, 0x93, 0x79, 0x78, 0xed, 0xc4, 0x37, 0x03, 0xf3, 0xf3,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x0f, 0xa4, 0x1d, 0x3a, 0x0c, 0x00, 0x00,
}
