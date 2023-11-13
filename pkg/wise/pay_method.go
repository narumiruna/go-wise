package wise

type PayInMethod string

const (
	PayInMethodApplePay            = PayInMethod("APPLE_PAY")
	PayInMethodBalance             = PayInMethod("BALANCE")
	PayInMethodBankTransfer        = PayInMethod("BANK_TRANSFER")
	PayInMethodCard                = PayInMethod("CARD")
	PayInMethodCredit              = PayInMethod("CREDIT")
	PayInMethodDebit               = PayInMethod("DEBIT")
	PayInMethodGbpSwift            = PayInMethod("GBP_SWIFT")
	PayInMethodGooglePay           = PayInMethod("GOOGLE_PAY")
	PayInMethodInternationalCredit = PayInMethod("INTERNATIONAL_CREDIT")
	PayInMethodInternationalDebit  = PayInMethod("INTERNATIONAL_DEBIT")
	PayInMethodMaestro             = PayInMethod("MAESTRO")
	PayInMethodMcBusinessCredit    = PayInMethod("MC_BUSINESS_CREDIT")
	PayInMethodMcBusinessDebit     = PayInMethod("MC_BUSINESS_DEBIT")
	PayInMethodMcCredit            = PayInMethod("MC_CREDIT")
	PayInMethodMcDebitOrPrepaid    = PayInMethod("MC_DEBIT_OR_PREPAID")
	PayInMethodPisp                = PayInMethod("PISP")
	PayInMethodReceiveBankTransfer = PayInMethod("RECEIVE_BANK_TRANSFER")
	PayInMethodSwift               = PayInMethod("SWIFT")
	PayInMethodTrustedPreFundBulk  = PayInMethod("TRUSTED_PRE_FUND_BULK")
	PayInMethodTrustedPreFundTx    = PayInMethod("TRUSTED_PRE_FUND_TX")
	PayInMethodVisaBusinessCredit  = PayInMethod("VISA_BUSINESS_CREDIT")
	PayInMethodVisaBusinessDebit   = PayInMethod("VISA_BUSINESS_DEBIT")
	PayInMethodVisaCredit          = PayInMethod("VISA_CREDIT")
	PayInMethodVisaDebitOrPrepaid  = PayInMethod("VISA_DEBIT_OR_PREPAID")
)

type PayOutMethod string

const (
	PayOutMethodBalance       = PayOutMethod("BALANCE")
	PayOutMethodBankTransfer  = PayOutMethod("BANK_TRANSFER")
	PayOutMethodDirectDebit   = PayOutMethod("DIRECT_DEBIT")
	PayOutMethodSwift         = PayOutMethod("SWIFT")
	PayOutMethodSwiftHk       = PayOutMethod("SWIFT_HK")
	PayOutMethodSwiftOurTier1 = PayOutMethod("SWIFT_OUR_TIER_1")
	PayOutMethodSwiftOurTier2 = PayOutMethod("SWIFT_OUR_TIER_2")
	PayOutMethodSwiftSgDbs    = PayOutMethod("SWIFT_SG_DBS")
	PayOutMethodSwiftUaPrivat = PayOutMethod("SWIFT_UA_PRIVAT")
	PayOutMethodWire          = PayOutMethod("WIRE")
)
