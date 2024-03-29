package uclone

import "ctpgo.app/ctpgo/core/common"

var (
	PCS         = common.NewExchange("PCS", "0x05fF2B0DB69458A0750badebc4f9e13aDd608C7F", "0xBCfCcbde45cE874adCB698cC183deBcF17952812", 20, 1)
	PCS_2       = common.NewExchange("PCS_2", "0x2ad2c5314028897aecfcf37fd923c079beeb2c56", "0x877Fe7F4e22e21bE397Cd9364fAFd4aF4E15Edb6", 25, 1)
	PCS_2_1     = common.NewExchange("PCS_2_1", "0x10ED43C718714eb63d5aA57B78B54704E256024E", "0xca143ce32fe78f1f7019d7d551a6402fc5350c73", 25, 1)
	BAKERY      = common.NewExchange("BAKERY", "0xc de540d7eafe93ac5fe6233bee57e1270d3e330f", "0x01bf7c66c6bd861915cdaae475042d3c4bae16a7", 30, 2)
	APESWAP     = common.NewExchange("APESWAP", "0xC0788A3aD43d79aa53B09c2EaCc313A787d1d607", "0x0841bd0b734e4f5853f0dd8d7ea041c241fb0da6", 20, 1)
	CAFESWAP    = common.NewExchange("CAFESWAP", "0x933DAea3a5995Fb94b14A7696a5F3ffD7B1E385A", "0x3e708fdbe3ada63fc94f8f61811196f1302137ad", 20, 1)
	DEFINIX     = common.NewExchange("DEFINIX","0x151030a9Fa62FbB202eEe50Bd4A4057AB9E826AD", "0x43eBb0cb9bD53A3Ed928Dd662095aCE1cef92D19", 20, 1)
	HYPERJUMP   = common.NewExchange("HYPERJUMP", "0x3bc677674df90A9e5D741f28f6CA303357D0E4Ec", "0xaC653cE27E04C6ac565FD87F18128aD33ca03Ba2", 30, 1)
	JULSWAP     = common.NewExchange("JULSWAP", "0xbd67d157502A23309Db761c41965600c2Ec788b2", "0x553990f2cba90272390f62c5bdb1681ffc899675", 30, 1)
	MANYSWAP    = common.NewExchange("MANYSWAP", "0x1C63048420aAeCEC6C418B456Dc9FeFFd9584e37", "0xae52c26976E56e9f8829396489A4b7FfEbe8aAE9", 10, 1)
	SLIME       = common.NewExchange("SLIME", "0x34766241a5DF0483545A52AB1DBd5eC88E251dD3ø", "0xcbe7425662bf0edf164abf12c881ced6fdafe75e", 20, 1)
	STM         = common.NewExchange("STM", "0x65b302EdC264604b335487542559658F79128EA5", "0x782536Abe989570211419352FbBf98A083380217", 30, 1)
	SWIPE       = common.NewExchange("SWIPE", "0x816278BbBCc529f8cdEE8CC72C226fb01def6E6C", "0x7810d4b7bc4f7faee9deec3242238a39c4f1197d", 30, 1)
	URANIUM     = common.NewExchange("URANIUM", "0x9a04b598CA32D3e9CDF366f34B2B4d41EF5f027F", "0x2C39801cc496E01B163CD3314295C30A98A26ef3", 20, 1)
	WARDEN      = common.NewExchange("WARDEN", "0xB75Fa2A799FC7935f37500Ba9780CBE10aA6610A", "0x3657952d7bA5A0A4799809b5B6fdfF9ec5B46293", 30, 1)
	ZERO        = common.NewExchange("ZERO", "0xba79bf6D52934D3b55FE0c14565A083c74FBD224", "0x52AbdB3536a3a966056e096F2572B2755df26eac", 30, 1)
	PURESWAP    = common.NewExchange("PURESWAP", "0x3e8743B5453A348606111AB0a4dEe7F70A87f305", "0x94b4188d143b9dd6bd7083ae38a461fcc6aad07e", 30, 1)
	CHEESESWAP  = common.NewExchange("CHEESESWAP", "0x3047799262d8d2ef41ed2a222205968bc9b0d895", "0xdd538e4fd1b69b7863e1f741213276a6cf1efb3b", 20, 1)
	MOCHISWAP   = common.NewExchange("MOCHISWAP", "0x939ffC5a4f3e9DF85e1036A8C86b18599A403F3B", "0xcbac17919f7aad11e623af4fea98b10b84802eac", 20, 1)
	URANIUM_V2  = common.NewExchange("URANIUM_V2", "0xf4ee46ac2ba83121f79c778ed0d950fff11a18ed", "0xa943ea143cd7e79806d670f4a7cf08f8922a454f", 16, 1)
	COMPLUS     = common.NewExchange("COMPLUS", "0x07dc75e8bc57a21a183129ec29bbcc232d79ee56", "0xdf97982bf70be91df4acd3d511c551f06a0d19ec", 30, 1)
	STABLEXSWAP = common.NewExchange("STABLEXSWAP", "0x8f2a0d8865d995364dc6843d51cf6989001f989e", "0x918d7e714243f7d9d463c37e106235dcde294ffc", 6, 1)
	DAILYSWAP   = common.NewExchange("DAILYSWAP", "0x106bd286705d9a6c6b61553a12fb3266d52de42a", "0x3e9a9ba1dcf2f3f7245642e398a39abd4546ad12", 20, 1)
	PANDASWAP_2 = common.NewExchange("PANDASWAP_2", "0x29D1Adbb65d93a5710cafe2EF0E8131f64E6AB22", "0x9ad32bf5dafe152cbe027398219611db4e8753b3", 30, 1)
	WAULTSWAP   = common.NewExchange("WAULTSWAP", "0xd48745e39bbed146eec15b79cbf964884f9877c2", "0xb42e3fe71b7e0673335b3331b3e1053bd9822570", 20, 1)
	BSCS        = common.NewExchange("BSCS", "0xa3b81724638208faa03810b391dde879e5fb0c79", "0x8b6ca4b3e08c9f80209e66436187088c99c9c2ac", 20, 1)
	COINSWAP    = common.NewExchange("COINSWAP", "0x34dbe8e5faefabf5018c16822e4d86f02d57ec27", "0xc2d8d27f3196d9989abf366230a47384010440c0", 20, 1)
	MOK         = common.NewExchange("MOK", "0xd4b5bf027d5ac9ab43c437959694b26c3a6f0c24", "0xe8a1c3cba77fa70ff3936d3106c3f6803b145cef", 20, 1)
	THUNDERSWAP = common.NewExchange("THUNDERSWAP", "0xf6135fcb4a0f469dcbb3e6d83520dc21825a0001", "0x6fcc6a77ee6f383395c630eede1ee928dff4e331", 20, 1)
	PANTHERSWAP = common.NewExchange("PANTHERSWAP", "0x24f7c33ae5f77e2a9eceed7ea858b4ca2fa1b7ec", "0x670f55c6284c629c23bae99f585e3f17e8b9fc31", 20, 1)
)

var UCloneExchanges = []*common.Exchange{
	PCS,
	PCS_2,
	PCS_2_1,
	BAKERY,
	APESWAP,
	CAFESWAP,
	DEFINIX,
	HYPERJUMP,
	JULSWAP,
	MANYSWAP,
	SLIME,
	STM,
	SWIPE,
	URANIUM,
	WARDEN,
	ZERO,
	PURESWAP,
	CHEESESWAP,
	MOCHISWAP,
	URANIUM_V2,
	COMPLUS,
	STABLEXSWAP,
	DAILYSWAP,
	PANDASWAP_2,
	WAULTSWAP,
	BSCS,
	COINSWAP,
	MOK,
	THUNDERSWAP,
	PANTHERSWAP,
}