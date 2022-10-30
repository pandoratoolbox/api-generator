package models

import (
	"time"
)

type Strings []string
type Ints []int64
type ctxkey int64

const (
	CTX_is_auth       = ctxkey(0)
	CTX_user_id       = ctxkey(1)
	CTX_user_role_ids = ctxkey(2)
	CTX_user_timezone = ctxkey(3)
)

type SubscriptionData struct {
	Id                *int64
	PaymentProviderId *int64
	ExternalId        *int64
	UserId            *int64
}

type Subscription struct {
	SubscriptionData
	PaymentProvider *PaymentProvider
	User            *User
}

type UserFinanceData struct {
	Id                             *int64
	UserId                         *int64
	Credits                        *int64
	DefaultPaymentProviderSourceId *int64
}

type UserFinance struct {
	UserFinanceData
	User                         *User
	DefaultPaymentProviderSource *PaymentProviderSource
}

type InvoiceData struct {
	Id                *int64
	ExternalId        *string
	UserId            *int64
	CreatedAt         *time.Time
	PaymentProviderId *int64
	AmountUsd         *int64
}

type Invoice struct {
	InvoiceData
	PaymentProvider *PaymentProvider
	User            *User
	Order           *Order
	OrderCredits    *OrderCredits
}

type UserSearchHistoryData struct {
	Id           *int64
	LastSearchAt *time.Time
	Queries      *Strings
	UserId       *int64
}

type UserSearchHistory struct {
	UserSearchHistoryData
	User *User
}

type PaymentProviderSourceData struct {
	Id                *int64
	UserId            *int64
	PaymentProviderId *int64
	ExternalId        *string
}

type PaymentProviderSource struct {
	PaymentProviderSourceData
	User            *User
	PaymentProvider *PaymentProvider
	UserFinance     []*UserFinance
}

type LeadData struct {
	UpdatedAt      *time.Time
	ExternalId     *string
	Id             *int64
	DataPlatformId *int64
}

type Lead struct {
	LeadData
	DataPlatform *DataPlatform
}

type PaymentProviderCustomerData struct {
	UserId            *int64
	Id                *int64
	PaymentProviderId *int64
	ExternalId        *string
}

type PaymentProviderCustomer struct {
	PaymentProviderCustomerData
	PaymentProvider *PaymentProvider
	User            *User
}

type UserData struct {
	CreatedAt  *time.Time
	Password   *string
	Username   *string
	Id         *int64
	Email      *string
	ReferrerId *int64
	RoleIds    *Ints
}

type User struct {
	UserData
	User                    []*User
	PaymentProviderCustomer []*PaymentProviderCustomer
	PaymentProviderSource   []*PaymentProviderSource
	UserSearchHistory       []*UserSearchHistory
	Order                   []*Order
	Invoice                 []*Invoice
	Subscription            []*Subscription
	UserFinance             *UserFinance
	OrderCredits            []*OrderCredits
}

type OrderData struct {
	Query            *string
	CreatedAt        *time.Time
	InvoiceId        *int64
	MustIncludeEmail *bool
	Amount           *int64
	LeadIds          *Ints
	Status           *string
	UserId           *int64
	DataPlatformId   *int64
	Id               *int64
	ScrapeJobId      *int64
}

type Order struct {
	OrderData
	Invoice      *Invoice
	DataPlatform *DataPlatform
	User         *User
}

type PaymentProviderData struct {
	ApiKey    *string
	Id        *int64
	Name      *string
	IsEnabled *bool
}

type PaymentProvider struct {
	PaymentProviderData
	Invoice                 []*Invoice
	Subscription            []*Subscription
	PaymentProviderCustomer []*PaymentProviderCustomer
	PaymentProviderSource   []*PaymentProviderSource
}

type OrderCreditsData struct {
	Credits   *int64
	CreatedAt *time.Time
	InvoiceId *int64
	UserId    *int64
	Id        *int64
}

type OrderCredits struct {
	OrderCreditsData
	Invoice *Invoice
	User    *User
}

type LeadInstagramData struct {
	LikesAverage      *int64
	Gender            *string
	ProfilePictureSrc *string
	LikesMin          *int64
	CommentsMin       *int64
	Id                *int64
	Following         *int64
	Followers         *int64
	ExternalId        *int64
	CommentsMax       *int64
	Bio               *string
	TagsPublished     *map[string]interface{}
	CommentsAverage   *int64
	LikesMax          *int64
	Website           *string
	Username          *string
	Language          *string
}

type LeadInstagram struct {
	LeadInstagramData
}

type DataPlatformData struct {
	Id                 *int64
	Name               *string
	ElasticsearchIndex *string
}

type DataPlatform struct {
	DataPlatformData
	Lead  *Lead
	Order []*Order
}
