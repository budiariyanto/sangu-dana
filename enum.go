package dana

type StatusDetailEnum int

const (
	Init StatusDetailEnum = iota
	Success
	Closed
	Paying
	MerchantAccept
	Cancelled
)

func (s StatusDetailEnum) String() string {
	return [...]string{"INIT",
		"SUCCESS",
		"CLOSED",
		"PAYING",
		"MERCHANT_ACCEPT",
		"CANCELLED"}[s]
}

type PayMethodEnum int

const (
	Balance PayMethodEnum = iota
	Coupon
	NetBanking
	CreditCard
	DebitCard
	VirtualAccount
	Otc
	DirectDebitCreditCard
	DirectDebitDebitCard
)

func (p PayMethodEnum) String() string {
	return [...]string{
		"BALANCE",
		"COUPON",
		"NET_BANKING",
		"CREDIT_CARD",
		"DEBIT_CARD",
		"VIRTUAL_ACCOUNT",
		"OTC",
		"DIRECT_DEBIT_CREDIT_CARD",
		"DIRECT_DEBIT_DEBIT_CARD"}[p]
}
