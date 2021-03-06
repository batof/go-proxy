package network

import "proxy/network/client"

// ActivePetUpdatePacket

type ActivePetUpdateCommand byte

const (
	PetFollow   ActivePetUpdateCommand = 1
	PetUnfollow ActivePetUpdateCommand = 2
	PetRelease  ActivePetUpdateCommand = 3
)

func GetPetUpdateType(command ActivePetUpdateCommand) string {
	switch command {
	case PetFollow:
		return "Follow"
	case PetUnfollow:
		return "Unfollow"
	case PetRelease:
		return "Release"
	default:
		return "Unknown: " + string(command)
	}
}

// BuyResultPacket

type BuyResultCode int32

const (
	BuyUnknown        BuyResultCode = -1
	BuySuccess        BuyResultCode = 0
	BuyInvalidChar    BuyResultCode = 1
	BuyItemNotFound   BuyResultCode = 2
	BuyNotEnoughGold  BuyResultCode = 3
	BuyInventoryFull  BuyResultCode = 4
	BuyRankTooLow     BuyResultCode = 5
	BuyNotEnoughFame  BuyResultCode = 6
	BuyPetFeedSuccess BuyResultCode = 7
	BuyUnknown2       BuyResultCode = 10 // haven't seen this error but it's in game code
)

func GetBuyResultType(code BuyResultCode) string {
	switch code {
	case BuyUnknown:
		return "Unknown: -1"
	case BuySuccess:
		return "Success"
	case BuyInvalidChar:
		return "Invalid Character"
	case BuyItemNotFound:
		return "Item Not Found"
	case BuyNotEnoughGold:
		return "Not Enough Gold"
	case BuyInventoryFull:
		return "Inventory Full"
	case BuyRankTooLow:
		return "Rank Too Low"
	case BuyNotEnoughFame:
		return "Not Enough Fame"
	case BuyPetFeedSuccess:
		return "Pet Feed Success"
	case BuyUnknown2:
		return "Unknown: code 10"
	default:
		return "Unknown: " + string(code)
	}
}

// Class Types

type PlayerClass int

const (
	Rogue       PlayerClass = 768
	Archer      PlayerClass = 775
	Wizard      PlayerClass = 782
	Priest      PlayerClass = 784
	Warrior     PlayerClass = 797
	Knight      PlayerClass = 798
	Paladin     PlayerClass = 799
	Assassin    PlayerClass = 800
	Necromancer PlayerClass = 801
	Huntress    PlayerClass = 802
	Mystic      PlayerClass = 803
	Trickster   PlayerClass = 804
	Sorcerer    PlayerClass = 805
	Ninja       PlayerClass = 806
	Samurai     PlayerClass = 785
	Bard        PlayerClass = 796
)

func GetClassName(class PlayerClass) string {
	switch class {
	case Rogue:
		return "Rogue"
	case Archer:
		return "Archer"
	case Wizard:
		return "Wizard"
	case Priest:
		return "Priest"
	case Warrior:
		return "Warrior"
	case Knight:
		return "Knight"
	case Paladin:
		return "Paladin"
	case Assassin:
		return "Assassin"
	case Necromancer:
		return "Necromancer"
	case Huntress:
		return "Huntress"
	case Mystic:
		return "Mystic"
	case Trickster:
		return "Trickster"
	case Sorcerer:
		return "Sorcerer"
	case Ninja:
		return "Ninja"
	case Samurai:
		return "Samurai"
	case Bard:
		return "Bard"
	default:
		return "Unknown: " + string(rune(class))
	}
}

// ContainerType

type ContainerType int32

const (
	ContainerNone      ContainerType = 0
	ContainerMap       ContainerType = 1
	ContainerEquipment ContainerType = 2
	ContainerInventory ContainerType = 3
	ContainerBackpack  ContainerType = 4
	ContainerQuickslot ContainerType = 5
	ContainerVault     ContainerType = 6
	ContainerPotion    ContainerType = 7
)

func GetContainerType(container ContainerType) string {
	switch container {
	case ContainerNone:
		return "None"
	case ContainerMap:
		return "Map"
	case ContainerEquipment:
		return "Equipment"
	case ContainerInventory:
		return "Inventory"
	case ContainerBackpack:
		return "Backpack"
	case ContainerQuickslot:
		return "Quickslot"
	case ContainerVault:
		return "Vault"
	case ContainerPotion:
		return "Potion"
	default:
		return "Unknown: " + string(container)
	}
}

// Currency - used in most purchase based packets

type CurrencyType byte

const (
	CurrencyGold          CurrencyType = 0
	CurrencyFame          CurrencyType = 1
	CurrencyGuildFame     CurrencyType = 2
	CurrencyFortuneTokens CurrencyType = 3
)

func GetCurrencyType(currency CurrencyType) string {
	switch currency {
	case CurrencyGold:
		return "Gold"
	case CurrencyFame:
		return "Fame"
	case CurrencyGuildFame:
		return "Guild Fame"
	case CurrencyFortuneTokens:
		return "Fortune Tokens"
	default:
		return "Unknown: " + string(currency)
	}
}

// GameID

type GameID int32

const (
	GameIDRealm         GameID = 0
	GameIDTutorial      GameID = -1
	GameIDNexus         GameID = -2
	GameIDRandomRealm   GameID = -3
	GameIDNexusTutorial GameID = -4
	GameIDVault         GameID = -5
	GameIDMapTest       GameID = -6
	GameIDRedirect      GameID = -7
	GameIDQuarantine    GameID = -13
)

// ChangeGuildRank

func GetGuildRank(rank client.GuildRankType) string {
	switch rank {
	case client.GuildInitiate:
		return "Initiate"
	case client.GuildMember:
		return "Member"
	case client.GuildOfficer:
		return "Officer"
	case client.GuildLeader:
		return "Leader"
	case client.GuildFounder:
		return "Founder"
	default:
		return "Unknown: " + string(rank)
	}
}

// InvResultPacket

type InvResultCode int32

const (
	InvSwapFailed  InvResultCode = -1
	InvSwapSuccess InvResultCode = 0
)

func GetInvResultType(code InvResultCode) string {
	switch code {
	case InvSwapFailed:
		return "Failed"
	case InvSwapSuccess:
		return "Success"
	default:
		return "Unknown: " + string(code)
	}
}

// PetUpgradeRequestPacket

type PetUpgradeType byte

const (
	UpgradePetYard PetUpgradeType = 1
	UpgradeFeedPet PetUpgradeType = 2
	UpgradeFusePet PetUpgradeType = 3
)

func GetPetUpgradeType(upgrade PetUpgradeType) string {
	switch upgrade {
	case UpgradePetYard:
		return "Pet Yard"
	case UpgradeFeedPet:
		return "Feed Pet"
	case UpgradeFusePet:
		return "Fuse Pet"
	default:
		return "Unknown: " + string(upgrade)
	}
}

// PetYardUpdatePacket

type PetYardType int32

const (
	YardCommon    PetYardType = 1
	YardUncommon  PetYardType = 2
	YardRare      PetYardType = 3
	YardLegendary PetYardType = 4
	YardDivine    PetYardType = 5
)

func GetPetYardType(code PetYardType) string {
	switch code {
	case YardCommon:
		return "Common"
	case YardUncommon:
		return "Uncommon"
	case YardRare:
		return "Rare"
	case YardLegendary:
		return "Legendary"
	case YardDivine:
		return "Divine"
	default:
		return "Unknown: " + string(code)
	}
}

// TradeDonePacket

type TradeDoneResult int32

const (
	TradeSuccessful TradeDoneResult = 0
	TradeCancelled  TradeDoneResult = 1
	TradeError      TradeDoneResult = 2
)

func GetTradeDoneType(code TradeDoneResult) string {
	switch code {
	case TradeSuccessful:
		return "Trade Successful"
	case TradeCancelled:
		return "Trade Cancelled"
	case TradeError:
		return "Trade Error"
	default:
		return "Unknown: " + string(code)
	}
}

// VaultType

type VaultType int32

const (
	VaultNone   VaultType = 0
	VaultNormal VaultType = 1
	VaultGift   VaultType = 2
	VaultPotion VaultType = 3
)

func GetVaultType(vault VaultType) string {
	switch vault {
	case VaultNone:
		return "None"
	case VaultNormal:
		return "Vault"
	case VaultGift:
		return "Gift"
	case VaultPotion:
		return "Potion"
	default:
		return "Unknown: " + string(vault)
	}
}
