package constants

const (
	GrpcJwtSecret       = "eM4rk3t"
	GrpcJwtHeaderPrefix = "Bearer "
	GrpcTimeoutInSecs   = 10
)

const (
	RedisPrincipalPrefix   = "eMarketPrincipal-"
	RedisCurrentUserPrefix = "eMarketUser-"
)

const (
	S3AccountSignatory                    = "account/signatory"
	S3MarketFloor                         = "market/floor"
	S3MarketHawkerObjection               = "market/hawker/objection"
	S3RentalNSAApprovalProof              = "rental/application/approval-proof" // step 5
	S3RentalNSAResidenceProof             = "rental/application/residence-proof"
	S3RentalNSABirthCert                  = "rental/application/birth-cert"
	S3RentalNSAPicture                    = "rental/application/picture"
	S3RentalNSATransferProof              = "rental/application/transfer-proof"
	S3RentalApplicationProgressAttachment = "rental/application/progress-proof"
	S3RentalNSASignedContract             = "rental/application/signed-contract"
	S3RentalSOAPaymentProof               = "rental/application/soa-payment-proof"
	S3RentalSOAOrProof                    = "rental/application/soa-or-proof"
	S3RentalHelperPicture                 = "rental/helper/picture"
	S3RentalHelperHealthCert              = "rental/helper/health-cert"
	S3RentalHelperResidenceCert           = "rental/helper/residence-cert"
	S3RateIssuanceComment                 = "rate/issuance/comment"
	S3RateIssuanceViolationProof          = "rate/issuance/violation-proof"
	S3RateIssuanceOrProof                 = "rate/issuance/issuance-proof"
	S3RentalRSAPaymentProof               = "rental/rsa/payment-proof"
	S3RentalRSAPicture                    = "rental/rsa/picture"
	S3RentalRSATransferProof              = "rental/rsa/transfer-proof"
	S3RentalRSAContract                   = "rental/rsa/signed-contract"
	S3RentalRPARenovationPlan             = "rental/rpa/renovation-plan"
	S3RentalRPATransferProof              = "rental/rpa/transfer-proof"
	S3RentalTSATransferDeed               = "rental/tsa/transfer-deed"
	S3RentalTSATransferProof              = "rental/tsa/transfer-proof"
	S3RentalTSAPaymentProof               = "rental/tsa/payment-proof"
	S3RentalTSADeathCert                  = "rental/tsa/death-cert"
	S3RentalTSAWaiverRights               = "rental/tsa/waiver"
	S3RentalTSABirthCert                  = "rental/tsa/birth-cert"
	S3RentalTSAPicture                    = "rental/tsa/picture"
	S3RentalTSAResidenceProof             = "rental/tsa/residence-proof"
	S3RentalTSATaxAndFeeProof             = "rental/tsa/tax-and-fee"
	S3RentalTSAContract                   = "rental/tsa/signed-contract"
	S3RentalMOPOrProof                    = "rental/mop/or-proof"
	S3RentalHawkerHelperPicture           = "rental/hawker/helper/picture"
	S3RentalHawkerHelperOrProof           = "rental/hawker/helper/or-proof"
	S3RentalVendingFeePaymentProof        = "rental/vending-fee/payment-proof"
	S3RentalVendingFeeOrProof             = "rental/vending-fee/or-proof"
	S3RentalVendingFeeApprovalLetter      = "rental/vending-fee/approval-letter"
	S3RentalNHPResidenceCert              = "rental/nhp/residence-proof"
	S3RentalNHPBarangayClearance          = "rental/nhp/barangay-clearance"
	S3RentalNHPVoterCert                  = "rental/nhp/voter-certificate"
	S3RentalNHPNbiClearance               = "rental/nhp/nbi-clearance"
	S3RentalNHPAssociationCert            = "rental/nhp/association-certificate"
	S3RentalNHPHealthCert                 = "rental/nhp/health-cert"
	S3RentalNHPTransferProof              = "rental/nhp/transfer-proof"
	S3RentalNPOBarangayClearance          = "rental/npo/barangay-clearance"
	S3RentalNPOTransferProof              = "rental/npo/transfer-proof"
	S3RentalNPOSecRegistration            = "rental/npo/sec-registration"
	S3RentalNPODtiCert                    = "rental/npo/dti-cert"
	S3RentalNPOCtcCert                    = "rental/npo/ctc-cert"
	S3RentalNPOFranchiseResolution        = "rental/npo/franchise-resolution"
	S3RentalNPOSanitaryPermit             = "rental/npo/sanitary-permit"
	S3RentalNPOEnviClearance              = "rental/npo/envi-clearance"
	S3RentalNPOFsiCert                    = "rental/npo/fsi-cert"
	S3RentalNPOLetterOfIntent             = "rental/npo/letter-of-intent"
	S3RentalNPONotarizedAuthorLetter      = "rental/npo/notarized-author-letter"
	S3RentalNPOBuildingPermit             = "rental/npo/building-permit"
	S3RentalNPOSiteDevPlan                = "rental/npo/site-dev-plan"
	S3RentalNPOStructuralStabilityCert    = "rental/npo/structural-stability-cert"
	S3RentalNPOOccupancyCert              = "rental/npo/occupancy-cert"
	S3RentalNPOStpSepticTank              = "rental/npo/stp-septic-tank"
	S3RentalNPOEnviComplianceCert         = "rental/npo/envi-compliance-cert"
	S3RentalNPOLocationClearance          = "rental/npo/location-clearance"
	S3RentalRHPHealthCert                 = "rental/rhp/health-cert"
	S3RentalRPOFranchiseResolution        = "rental/rpo/franchise-resolution"
	S3RentalRPOEnviClearance              = "rental/rpo/envi-clearance"
	S3RentalRPOLocationClearance          = "rental/rpo/location-clearance"
	S3RentalRPOSanitaryPermit             = "rental/rpo/sanitary-permit"
	S3RentalRPOFsiCert                    = "rental/rpo/fsi-cert"
	S3RentalRPOLetterOfIntent             = "rental/rpo/letter-of-intent"
	S3RentalRPONotarizedAuthorLetter      = "rental/rpo/notarized-author-letter"
	S3RentalRPOSiteDevPlan                = "rental/rpo/site-dev-plan"
	S3RentalRPOStructuralStabilityCert    = "rental/rpo/structural-stability-cert"
	S3RentalRPOOccupancyCert              = "rental/rpo/occupancy-cert"
	S3RentalRPOStpSepticTank              = "rental/rpo/stp-septic-tank"
	S3RentalRPOTransferProof              = "rental/rpo/transfer-proof"
	S3RentalNPOInspectionReport           = "rental/npo/inspection-report"
	S3RentalRPOInspectionReport           = "rental/rpo/inspection-report"
	S3RentalNHPHawkerPermitPicture        = "rental/nhp/hawker-permit-picture"
	S3RentalNHPSignatureHawkerId          = "rental/nhp/signature-hawker-id"
	S3RentalRPOBuildingPermit             = "rental/rpo/building-permit"
	S3RentalRPOEnviComplianceCert         = "rental/rpo/envi-compliance-cert"
)
