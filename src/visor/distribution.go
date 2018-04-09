package visor

import (
	"github.com/skycoin/skycoin/src/coin"
)

const (
	// MaxCoinSupply is the maximum supply of skycoins
	MaxCoinSupply uint64 = 3e8 // 100,000,000 million

	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100

	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100

	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// GetDistributionAddresses returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// GetUnlockedDistributionAddresses returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (25) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (25) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (25) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// GetLockedDistributionAddresses returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// TransactionIsLocked returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"4EmJ5dfx3wsHwWKsMTfQ31p5uYJfJdkohh",
"QuoFvFwtkekMEYRhXyX6FLdZgYaepwHLrR",
"pUCFtuV98KHrTucKPRgjCGZge5ADGNWJKZ",
"vZjW1xFWefDdkjZjBirsBbLx4ShmqMLPXc",
"29kPNEyTUXbAQsfDBRYwvjxv7AUsEfkHF2W",
"2KFTykHNkeT2ab2rXR4trNgqbsnyGihXR5i",
"qMkYYXpHYxnQF82dGSU1NPZC5rnrcMKLuq",
"22upsW6afi8XeHbKGA6LXQRLuMQJxrbZNSC",
"2KfCbr4u41Kwwd9p2qFAYi8q846vuBhEdf7",
"2LUkM89FH73aDLNLYHeSAXznpMkfZYZworz",
"CeYwWAnzzHq2AKVjraYQHke1koewCaAtYi",
"HV3a26bJydKEyFo9KshRNKM93YXYZSnLeL",
"UwukfXR39CimFfDkBCyh2uE2j7DW45PNv5",
"2SkYKxbWU9PeWaXczT3wMNZAUuFt6hm97yp",
"2UKw1dPW8ih2ECSwhSbsvW2VS6KNGpktiw4",
"2KE3JCVhWpuKesiC1W8kzHXEMi3crwAVPVJ",
"2UzqsvDfesa4TEWpmEiDHVNDNXzrxEXhwGp",
"Ut6UTqYbcu8QVYB9qnh6Sgrq5ztKRKjXy7",
"2R7ukNJG2ELVUsX2GJwNvoy7FgQQE3yMVpv",
"DrxrNGUWUujcA5etb9fi2KtimkDmidBDry",
"2iYrEDYjMizGZxFgqb4msocd6yDYYzUH8UC",
"dS67qDfqTod2viEScxm57oCnGogD5sk1Wp",
"gr6RVkPzR8gB8w1a2fJ4f6Uv3hwa8WHKZm",
"aeHM9sEPzJ8nW5g9msXEb1DHoRgg7HD2Uo",
"kTxV9FvMXqRtSMRxA8BVnYbPnC1GGcGEwZ",
"jZ8365eMAMtjVTcvxATYh1Z13iSwE4DM1w",
"2NJGW1RbbRFh6Bqjf33cs4f1QHtBEZHSkK2",
"2EULQ7JMpfut1fCCieLa1QG3Emh1R3emJYV",
"wXnFqaJ5t4YKuYtts1vjNmn22yHrMDwHdL",
"yjczN2BwDA8DSjs1aDJ72vPm6mBp6YzwNB",
"gz6J6eaf7v5wWrBH6uA8rEwCeEtE6m4uTZ",
"2L3UoEFAkjSvw1xp2HX7EfLR7MmtmuKEF4K",
"2bUjBYtansT9j9qBGjoXf5NM818rs465L69",
"TQwgQFP5rhSfnDLXZvX2BvN76PbUw6TPw",
"2FR126USk7RHzTZFfneEsnXtfi2fA7rBrJg",
"pvo9qDEMgxJ8sVp823UGECvxyHepsYmGy5",
"dWqXe4cWjP8dpypTdo1WgBQWgkRge83RYx",
"2frwoqsfxnaEZUTQMSyt8zPRisk37AudWHZ",
"dU9ZF6ymfAzaX8JUwAPwysxf3e63t6aJvB",
"5XuNJ735yoM8pMptV8Xb86XeQTqDeLNhad",
"GCoLLGVccM55uqAydu2sEuTKTboPYcjLSq",
"gVqLaT3m25XSuCkTKL81oa6GQZtHEUghCv",
"2FL8aRBu6PkJtbUtrP2FrjA6RfMMuMGrogb",
"RNKy29dcpByB1vNzwS3MBSAVmLeNEL5r1u",
"FAtwXixRrwCof2Q8d1G9oRRL21Pu97QzYE",
"2jnFU4WhTspT5xF9xBmEGCyjBzGrpwV9KmU",
"2LPpaLNr1SGRANSjtrBNdyQ3CsLiuJrJ9vG",
"29CWC1Q5mwuvGkcUzNsUeKYpDGTYQwKHUz5",
"2VzVTqprsaxX83Ty3RDB8qxxj8okCsX1uZV",
"Pi5PEQWLzM6VAx3teH5njsfS3aRe8uM33d",
"2hkYymVJnTkFRNNaizTvTtCfZjnXUJdB4Vd",
"Bd8NgUn7M2AZ5DhJcptMSt18ktawytraRm",
"N1crujfwpmE5GYxiCyNNhNawr3WR5o8DEN",
"Qrmf84dw5qhWAFsCRE8jGwZNkdjtuD2tVs",
"CRofzyPh9trHsmE9Me5BZtDYFrZHbWeeiR",
"joMkrn1dDQ6HXijZ4N2c6w18B87mStsrYk",
"4hhUdCSnA8Am15GsuMxZH5UDcuTr224JEf",
"zFPE7UnL4D61XhocQTdd7aWU61qDcP1axo",
"6UEW22dRvEdXegvE9RvP86cjHQjfYMaBF3",
"2Ukg5PK8ur3ufhV728zh2w1giBTuhrudy7u",
"h3MmQzJyXRkkZQtcsqexp1fZEs1M2uPFAD",
"2MeTkW453oaHXZpQz1RXm9fT85VBxKxEVy6",
"29nCVWPwzGqJqyzbmiHmmmipm7vb7H3hv1m",
"2VLfYZauxkN3AiVxdCmRD2KTbrLCQyHQin4",
"YjmWMwGwkyMg1wFNRKWDwCwHf3Pb116PxC",
"yxocxuh2U7SdwfNGXD6ntsQVsQTpyLS56i",
"s4obnTHe4Kuokn3MdxW7o9whzeMycv7hr2",
"At5XWpAgdFBoYKadcVTfTqaZfHu57ozFHu",
"ySUjssfqfHW2dcV6MKj4K6gqWAc1BDZsGw",
"AY96pFVKHAK2WzjhP9UDkv4pG2j5bqBVmi",
"gnVeepuUGh31tcyWTK3Tz2FQJ14euVdKK8",
"pAJTiVpQzX14LEKMB8Bz1ExGaJ3jDY1VGf",
"BVDEFRqgEyVvFsFx6fgBHCE5Y5RWgfBA1n",
"29qvewmoSV6NHifFLA76pYbSM6QawE2adKC",
"2hTQQCkfzp55LvysLeFzps5yY8WBboEjZ9e",
"WS57hc5SpncSrFkFpfa72PcJdyCyeKh1P1",
"2K1MQ58KbvyC1BWh9ptwHkBN5PnPhdsizJH",
"22DFsFEDcEKzLC5wN9qGAJP39myGB7pNLf2",
"2QxWANHhQDEKX3KVFiJADFGWHLeptXmH6Bb",
"2aTnyKGwZTqVkfmf1wfSGDAmyatdY2meaS3",
"2G1dQFH7rwyDpgLkgxWqUnVJhEP5nRsqkgS",
"23gkEujwMYDq1NZMpyQ6rNmiAYpZNCroTSd",
"6HdUiFNTKZStZecVkC4VahcQ7wZ4MJattr",
"25RAFrzh85QEQBP54HN4gtpE6dUdNFKLH87",
"CFr84AHkbdL8CzFdHbbMGZLT5CFAiLSKaG",
"21fU9ALUcFrQNQ1CQZv7n3JJitZYUsnbPZR",
"hs6me1EvzJPSMnGZJVG3Ucn5x8xviJUgig",
"RV4TvVCbyvLfcuPPoXJjpGNNwBy7Nm5Pau",
"28wAXi74mQdfUKBVVS6TcdZ8qFJSnPhP5cz",
"zT7hSSsYf5erHMQRAENauNox3AfJ6dguKE",
"c3xiX4PyKiAJf9Kug4iuSD4iTKWfDpDdG",
"xFJocvPFCYP3EP2r2Nv1PjNuCwFeGRhKDp",
"kmoacFgosG6KJMHe9KDiR2tzRwWmgPmfsB",
"2DZrmXmMPUe43ZHnef35NZg9ktk35c7r3Ep",
"mHoGgj8ahJanEFMDpoxey2KKupqS81bucM",
"2iY6RrqftgfS4o66eg93yehmSWTvuzHTrF2",
"KBuWGafcLr4ht1BT8ykbgiCpQNj4Bore1g",
"F94opeQNavbkh7eZKsk4EJ3Lx38p3AYvCE",
"RA9VsokcqMgnRow5osuUV8zwQdrHntMxaZ",
"uN7JuRyBtbMvRfP26Ui3pnexch8WgoGfdf",
}
