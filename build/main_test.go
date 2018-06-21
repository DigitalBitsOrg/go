package build

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/digitalbitsorg/go/build")
}

// ExampleTransactionBuilder creates and signs a simple transaction, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
//
// It uses the transaction builder system
func ExampleTransactionBuilder() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Payment(
			Destination{"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"},
			NativeAmount{"50"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAAAAAAB3NZQAAAAAAAAAAAUfV8K8AAABAjyBwOlr+MJukhUiBtrxL5rw1ls1U0cnYGfftlpQZ/wk55mi11u6KOe+mgpPZrVFftv71J3DRlm45hx84ZgECAg==
}

// ExampleCreateAccount creates a transaction to fund a new stallar account with a balance. It then
// encodes the transaction into a base64 string capable of being submitted to stellar-core. It uses
// the transaction builder system.
func ExampleCreateAccount() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		CreateAccount(
			Destination{"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"},
			NativeAmount{"50"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAHc1lAAAAAAAAAAABR9XwrwAAAECTG/Z8Q60kGiHJlMBDhcEO3ten5I8z0rWPXZglw6fngnjATFZ40td2jy7fji90dlGMCFCPj4eO0T5ivJQCcbUC
}

// ExamplePayment creates and signs a native-asset Payment, encodes it into a base64 string capable of
// being submitted to stellar-core. It uses the transaction builder system.
func ExamplePayment() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Payment(
			Destination{"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"},
			NativeAmount{"50"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAAAAAAB3NZQAAAAAAAAAAAUfV8K8AAABAjyBwOlr+MJukhUiBtrxL5rw1ls1U0cnYGfftlpQZ/wk55mi11u6KOe+mgpPZrVFftv71J3DRlm45hx84ZgECAg==
}

// ExamplePathPayment creates and signs a simple transaction with PathPayment operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExamplePathPayment() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Payment(
			Destination{"GBDT3K42LOPSHNAEHEJ6AVPADIJ4MAR64QEKKW2LQPBSKLYD22KUEH4P"},
			CreditAmount{"USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA", "50"},
			PayWith(CreditAsset("EUR", "GCPZJ3MJQ3GUGJSBL6R3MLYZS6FKVHG67BPAINMXL3NWNXR5S6XG657P"), "100").
				Through(Asset{Native: true}).
				Through(CreditAsset("BTC", "GAHJZHVKFLATAATJH46C7OK2ZOVRD47GZBGQ7P6OCVF6RJDCEG5JMQBQ")),
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAIAAAABRVVSAAAAAACflO2Jhs1DJkFfo7YvGZeKqpze+F4ENZde22bePZeubwAAAAA7msoAAAAAAEc9q5pbnyO0BDkT4FXgGhPGAj7kCKVbS4PDJS8D1pVCAAAAAVVTRAAAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAHc1lAAAAAAIAAAAAAAAAAUJUQwAAAAAADpyeqirBMAJpPzwvuVrLqxHz5shND7/OFUvopGIhupYAAAAAAAAAAUfV8K8AAABAJ0/e2kxIGn3VZ+8bRW++cWzoU9NOWU1SQtkdbzOQzPnD65d7mbdA7yLXYTI3hI5qKMl1Me1PguRx/5MEitioDg==
}

// ExampleSetOptions creates and signs a simple transaction with SetOptions operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleSetOptions() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		SetOptions(
			InflationDest("GCT7S5BA6ZC7SV7GGEMEYJTWOBYTBOA7SC4JEYP7IAEDG7HQNIWKRJ4G"),
			SetAuthRequired(),
			SetAuthRevocable(),
			SetAuthImmutable(),
			ClearAuthRequired(),
			ClearAuthRevocable(),
			ClearAuthImmutable(),
			MasterWeight(1),
			SetThresholds(2, 3, 4),
			HomeDomain("stellar.org"),
			AddSigner("GC6DDGPXVWXD5V6XOWJ7VUTDYI7VKPV2RAJWBVBHR47OPV5NASUNHTJW", 5),
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAUAAAABAAAAAKf5dCD2RflX5jEYTCZ2cHEwuB+QuJJh/0AIM3zwaiyoAAAAAQAAAAcAAAABAAAABwAAAAEAAAABAAAAAQAAAAIAAAABAAAAAwAAAAEAAAAEAAAAAQAAAAtzdGVsbGFyLm9yZwAAAAABAAAAALwxmfetrj7X13WT+tJjwj9VPrqIE2DUJ48+59etBKjTAAAABQAAAAAAAAABR9XwrwAAAECXprFvG8Gfl76n6T6t5zE+hoZ6osHsLRtUjh1dt0Q7W1zkai552mUnusY9n0Qc06mjLxNYLqeOhMI8p+FLz+gH
}

// ExampleSetOptions_manyOperations creates and signs a simple transaction with many SetOptions operations, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleSetOptions_manyOperations() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		InflationDest("GCT7S5BA6ZC7SV7GGEMEYJTWOBYTBOA7SC4JEYP7IAEDG7HQNIWKRJ4G"),
		SetAuthRequired(),
		SetAuthRevocable(),
		SetAuthImmutable(),
		ClearAuthRequired(),
		ClearAuthRevocable(),
		ClearAuthImmutable(),
		MasterWeight(1),
		SetThresholds(2, 3, 4),
		HomeDomain("stellar.org"),
		RemoveSigner("GC6DDGPXVWXD5V6XOWJ7VUTDYI7VKPV2RAJWBVBHR47OPV5NASUNHTJW"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAETAAAAAAAAAABAAAAAAAAAAAAAAALAAAAAAAAAAUAAAABAAAAAKf5dCD2RflX5jEYTCZ2cHEwuB+QuJJh/0AIM3zwaiyoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAQAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAQAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAQAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAABAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAABAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAABAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAEAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAgAAAAEAAAADAAAAAQAAAAQAAAAAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAALc3RlbGxhci5vcmcAAAAAAAAAAAAAAAAFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAALwxmfetrj7X13WT+tJjwj9VPrqIE2DUJ48+59etBKjTAAAAAAAAAAAAAAABR9XwrwAAAEBhRvwA5lJsvpr4XSBZZ2m0OgGROxBzD56Jf729no2J1PFjcKVBvUAamkx3+Ob5+A+LooCpWg7VKVNp/Qa6CfYE
}

// ExampleChangeTrust creates and signs a simple transaction with ChangeTrust operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleChangeTrust() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Trust("USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA", Limit("100.25")),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQQAAAAA7wO+gAAAAAAAAAAFH1fCvAAAAQJlwsiNA161gIeSmgeUnVbyjiy1qW/bKbVvAtTg6gaKfzRGfBdPBeD5P330X3iFgZCLTIYqFQFZgt0ZC1/opSQU=
}

// ExampleChangeTrust_maxLimit creates and signs a simple transaction with ChangeTrust operation (maximum limit), and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleChangeTrust_maxLimit() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Trust("USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQX//////////AAAAAAAAAAFH1fCvAAAAQIdjRjQtZLw/Roj6zCqqSVBOhl1214n/4ML5RYiqX4JZq6+9tegrFejIE0O6RkHI+PiIcHBSZLx4z6dfx3WSRgU=
}

// ExampleRemoveTrust creates and signs a simple transaction with ChangeTrust operation (remove trust), and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleRemoveTrust() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	operationSource := "GCVJCNUHSGKOTBBSXZJ7JJZNOSE2YDNGRLIDPMQDUEQWJQSE6QZSDPNU"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		RemoveTrust(
			"USD",
			"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA",
			SourceAccount{operationSource},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAQAAAACqkTaHkZTphDK+U/SnLXSJrA2mitA3sgOhIWTCRPQzIQAAAAYAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQQAAAAAAAAAAAAAAAAAAAAFH1fCvAAAAQNPZxXNnSSUz7GOUvRYLPlFgq+g2P/0cT8uWW+p8u0cQF7B214rVK6tGtj0nclhUMWBdp74OIktAoGkfJlVIkgQ=
}

// ExampleManageOffer creates and signs a simple transaction with ManageOffer operations, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleManageOffer() {
	rate := Rate{
		Selling: NativeAsset(),
		Buying:  CreditAsset("USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"),
		Price:   Price("125.12"),
	}

	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		CreateOffer(rate, "20"),
		UpdateOffer(rate, "40", OfferID(2)),
		DeleteOffer(rate, OfferID(1)),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAABLAAAAAAAAAABAAAAAAAAAAAAAAADAAAAAAAAAAMAAAAAAAAAAVVTRAAAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAC+vCAAAADDgAAAAZAAAAAAAAAAAAAAAAAAAAAwAAAAAAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQQAAAAAX14QAAAAMOAAAABkAAAAAAAAAAgAAAAAAAAADAAAAAAAAAAFVU0QAAAAAAC0kaS7Qi79nm6GZRIhw0hkeh2/s2S/dn20nTaTm3hNBAAAAAAAAAAAAAAw4AAAAGQAAAAAAAAABAAAAAAAAAAFH1fCvAAAAQIH+EHHvEctM8waoBeCxkRBWJIVo3ELumB7h2604zpVxIwJ6xlsZ0Jw9n6eqQQsk/K38Pd5D6aGEdRFn7sr5tgM=
}

// ExampleCreatePassiveOffer creates and signs a simple transaction with CreatePassiveOffer operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleCreatePassiveOffer() {
	rate := Rate{
		Selling: NativeAsset(),
		Buying:  CreditAsset("USD", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"),
		Price:   Price("125.12"),
	}

	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		CreatePassiveOffer(rate, "20"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAQAAAAAAAAAAVVTRAAAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAC+vCAAAADDgAAAAZAAAAAAAAAAFH1fCvAAAAQAOnPE15pKGjRbM1gQQosZaaluNp6arrryOXJh5fYfEsCY81vXgMGrE2dCglANlyuxNDgPfEQkKIvj4RpiBrWg8=
}

// ExampleAccountMerge creates and signs a simple transaction with AccountMerge operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleAccountMerge() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		AccountMerge(
			Destination{"GBDT3K42LOPSHNAEHEJ6AVPADIJ4MAR64QEKKW2LQPBSKLYD22KUEH4P"},
		),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAgAAAAARz2rmlufI7QEORPgVeAaE8YCPuQIpVtLg8MlLwPWlUIAAAAAAAAAAUfV8K8AAABAyCJwlCt7cCq7K7J8l3IFRAbOFEPARfHDwwFEwNb1vO2mOhRt1gfzscFFURGasfaD3TgJ5ps7+IQBRgsH5N9EAA==
}

// ExampleInflation creates and signs a simple transaction with Inflation operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExampleInflation() {
	seed := "SDBNWCXQ6JTYG2SI6373PAUQXH2LYSZWKCXK262FAWHAG4MIHND3HCOL"
	tx, err := Transaction(
		SourceAccount{seed},
		Sequence{1},
		TestNetwork,
		Inflation(),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	txe, err := tx.Sign(seed)
	if err != nil {
		fmt.Println(err)
		return
	}

	txeB64, err := txe.Base64()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAAMf5E8TaGF3AIah9lVLkmsOHQYmAAFTlM6xtEetH1fCvAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAkAAAAAAAAAAUfV8K8AAABAGBKhimU2zwJpou44LtqV9nf1fjSMfHW9rrXx0BfTIysHjPU1plzYHGs0r84ypsEqfY8TPsxeOHNTLXa/YeIqAg==
}
