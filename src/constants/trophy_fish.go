package constants

type TrophyFish struct {
	DisplayName string            `json:"display_name"`
	Description string            `json:"description"`
	Textures    map[string]string `json:"textures"`
}

var TROPHY_FISH = map[string]TrophyFish{
	"BLOBFISH": {
		DisplayName: "Blobfish",
		Description: "§7Caught everywhere.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/3e77a61f2a25f19bb047be985b965f069d857502881bea3f9682d00bfd5cc3e7",
			"silver":  "http://localhost:8080/api/head/40e6628e5c358e080d3180c7424b70371e67e5724020b68a64d60ab41fe70bae",
			"gold":    "http://localhost:8080/api/head/2e93d72c19e7b917f233dc291b749391f2d870ef30cafc5f17281023ae17c2b9",
			"diamond": "http://localhost:8080/api/head/d9339dfceaa5d99a5571ec4743b004d827b229eaf00eb703ced03a86029f3ad",
		},
	},
	"FLYFISH": {
		DisplayName: "Flyfish",
		Description: "§7Caught from §a8 §7blocks or higher above lava in the §6Blazing Volcano§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/5140c42fc3a1ba2fe77c45b975fa87e8f54a1b1833bc76e6339b4c262304011d",
			"diamond": "http://localhost:8080/api/head/1516c233797626ce5e0132166062ae1cb435c561c00a6fa11b37e2295f8c7c5b",
			"gold":    "http://localhost:8080/api/head/6051fc3ea4b8459df839e67c157e9d1a753be0603cbf18d3c1352ebf61b58581",
			"silver":  "http://localhost:8080/api/head/adc59df19336fa7a84e4637519d7f843c38e7673ffcc3f39edeed950e28d8b57",
		},
	},
	"GOLDEN_FISH": {
		DisplayName: "Golden Fish",
		Description: "§7Has a chance to spawn after §a8 §7minutes of fishing, increasing linearly until reaching §f100% §7at §a12 §7minutes. The §6Golden Fish §7is rolled when your fishing hook is thrown out, regardless of if you catch a fish or not. You are not required to Trophy Fish to catch this one.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/fcfa31d81eae819936834e6664430daf8affbff30b48a1daa7ca1b4c20e3fe7d",
			"diamond": "http://localhost:8080/api/head/f46cb8cbb60cda60092b3051648dd2db7eec49ae87f3f524c6182797b4109a12",
			"gold":    "http://localhost:8080/api/head/38175538210af6a4d1a443e08ee321abad837c6c927c7803ab361d29e3f8e509",
			"silver":  "http://localhost:8080/api/head/98ac13a462e9e9ffc3569c6beb44e4b28be5a15e7628c8957c5cddfaff26c285",
		},
	},
	"GUSHER": {
		DisplayName: "Gusher",
		Description: "§7Caught within §a7-16 §7minutes after a §6Blazing Volcano §7eruption.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/afb8c51bfcd996840010bcce2a3575ae26352e00446d3ec313fcbf1f88108512",
			"diamond": "http://localhost:8080/api/head/577322d86e61df34b2dcbb3fa7f4d03d0e3be56bddac6bdf2e7de61e21f718eb",
			"gold":    "http://localhost:8080/api/head/d550a51e76f32aaa556b98d817db6ae71aadb2af7d76c34b903dad5ee2f90439",
			"silver":  "http://localhost:8080/api/head/e5c2828e950c1fd7f73a6b7879d8b562ecff894b6cb22afb9c2660627b3b8dfe",
		},
	},
	"KARATE_FISH": {
		DisplayName: "Karate Fish",
		Description: "§7Caught in the lava pools near the §eDojo§7. Note - Half of the lava pools do not actually count as being in the §eDojo §7area. If you stand in the same place as your bobber and do not see '§eDojo§7' in the sidebar, you cannot catch the §5Karate Fish §7there.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/901ef47164aba899674cac1e04dda29895ba670807a162e888681c6428e42a83",
			"diamond": "http://localhost:8080/api/head/fa57e62e393f80cd4bb23109c7c869b1f66621ceccae4b7601afb3269120294d",
			"gold":    "http://localhost:8080/api/head/ac28a1f0f3908274e5860e1ac1b122b7323327aeb376e8378a036c9fa8bf35c5",
			"silver":  "http://localhost:8080/api/head/e56227d0e98b7541b2922e9cd9113155d13028978d4c3f68775199ea27299fd4",
		},
	},
	"LAVA_HORSE": {
		DisplayName: "Lavahorse",
		Description: "§7Caught everywhere.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/1176ea86635c4e849469ed694b3c0c3f7ec7677ed0682bee5ef6d59ea669677f",
			"diamond": "http://localhost:8080/api/head/67e685788792dc90e7d19c2932f3dc1dbe396e6cac068260997ea0b64ffd2bf8",
			"gold":    "http://localhost:8080/api/head/a4d8850536ca5d3b735691be439c7f8619ebc14ad0aef78055fd86a37dd2adf1",
			"silver":  "http://localhost:8080/api/head/124ab405566c448c7484132f99a7bb79bdf547ae50713c0dd866a2375156b2c7",
		},
	},
	"MANA_RAY": {
		DisplayName: "Mana Ray",
		Description: "§7Caught when you have at least §b1,200 ✎ Intelligence§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/ff357b0f4e13cc2013bf4a02a3d3351ab0d7856e73f9f0cf8b6f13e78d95b215",
			"diamond": "http://localhost:8080/api/head/6d6dacd67f0562980e597a8ba508b3e321002f4d358c85dd9a4a39bacaea63f8",
			"gold":    "http://localhost:8080/api/head/c515e96329992b2969c427a17ae68173ad0f8b86f9e3a0e3f9460385581edfe3",
			"silver":  "http://localhost:8080/api/head/dfd706157c2a3e2c8c3674cbc37f7a1206b87b7a6453d11e7408492c01c35634",
		},
	},
	"MOLDFIN": {
		DisplayName: "Moldfin",
		Description: "§7Caught in the §dMystic Marsh§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/54f33dc405ba447b35926b48d90600302aeebb140ad330d885886cb1029a8af",
			"diamond": "http://localhost:8080/api/head/9c99ddc1d711f482305d8f1ffcadd4b444fad1b0b07c5be2b73b1f08ee6cbe5e",
			"gold":    "http://localhost:8080/api/head/5113e386937a8ed25fe40fdbe6b88712ae73411f6a0ae927420ffcb9d778226b",
			"silver":  "http://localhost:8080/api/head/6a52dd3040aca257bf361c15cc4a0114f850b2a91867ee49d26cebca5203aab4",
		},
	},
	"OBFUSCATED_FISH_1": {
		DisplayName: "Obfuscated 1",
		Description: "§7Caught with Corrupted Bait or dropped from corrupted Sea Creatures.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/e1f4d91e1bf8d3c4258fe0f28ec2fa40670e25ba06ac4b5cb1abf52a83731a9c",
			"diamond": "http://localhost:8080/api/head/caa0b4b4f443257e83176df4ffd550de7ee89867e506b9c1ca53f33611327929",
			"gold":    "http://localhost:8080/api/head/8a2a44913ee1d5babc172f374351ea7ad1516ca256d16a4ef72d8a092b519cd1",
			"silver":  "http://localhost:8080/api/head/479b52391ff0cd3c83db1c1c218a02ab13a2c6a4aaa1cf126c7912bd377e8fbf",
		},
	},
	"OBFUSCATED_FISH_2": {
		DisplayName: "Obfuscated 2",
		Description: "§7Caught whilst using Obfuscated 1 as bait.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/8321e19aa4b3163c8990b066b1cd0895c3c97a799057327507db0ffc80d90575",
			"gold":    "http://localhost:8080/api/head/4631953a0351988029b90e838181e4e563d782e470ea33b8c612756f730625c2",
			"silver":  "http://localhost:8080/api/head/cb12de47e0b48ab8f7d8f500fdc5d7869b7f2192f823620088582a56afcf68fb",
			"diamond": "http://localhost:8080/api/head/cdca1057973e87f875722d7cf5c7b3de2aa4831ced2aa4259c0e4bec7b499245",
		},
	},
	"OBFUSCATED_FISH_3": {
		DisplayName: "Obfuscated 3",
		Description: "§7Caught with Obfuscated 2 as bait.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/df478663687d16f79b9c32546a7c3ec2736e87ce69779991e52deaf622abd8c2",
			"silver":  "http://localhost:8080/api/head/3c800c71b925587213382eeaaa426ed63112503e278ff9f5b3d39dbdb95d31d0",
			"gold":    "http://localhost:8080/api/head/97f71c13302401772e611a2a508f23df54b778be725ce231662f3fc810d258a1",
			"diamond": "http://localhost:8080/api/head/665a04023ce40813abee55061fe802d2f8195fcdee3570388bba072073ecef3a",
		},
	},
	"SKELETON_FISH": {
		DisplayName: "Skeleton Fish",
		Description: "§7Caught in the §eBurning Desert§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/923e0a25048b60a2cc092f72943159ec170063bb235aa79690ef34ab181d691",
			"diamond": "http://localhost:8080/api/head/ed01389874c7be1165d5df633daf27d936bfaf553143cfcbaa50c93c4746f9f3",
			"gold":    "http://localhost:8080/api/head/639dd2fe302e2ac4d9e8d31876489258b04e58a8e8754714669145a89a82d2e0",
			"silver":  "http://localhost:8080/api/head/a39846ad1c8fd3420febff458401bacd91d1f4fd444940f7f7cb9e2a490ca4dd",
		},
	},
	"SLUGFISH": {
		DisplayName: "Slugfish",
		Description: "§7Caught when the bobber has been active for at least §a20 §7seconds. The §6Slug Pet §7reduces this time by up to §a50%§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/c1de9e49ecc8d6209c783bfd1684a89e624a4e483a86023c6df57f77d5b75890",
			"diamond": "http://localhost:8080/api/head/a5d717aa5c9063181283811d265bfd0ffdc7eda09a0984cee59578b4a5efd4a1",
			"gold":    "http://localhost:8080/api/head/289d72f3750a5cd244cbf09af9478453b6575df591d720b6ec23d84a165c65f2",
			"silver":  "http://localhost:8080/api/head/d82efd885e6e2a964efb857de724b2ef043f1dcbbe618f10ec3742c6f2cecab",
		},
	},
	"SOUL_FISH": {
		DisplayName: "Soul Fish",
		Description: "§7Caught in the §5Stronghold§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/7fe554d346c20c161aa85cfdc1b89779c9f64e726a5bb28ace8078e6594052d7",
			"diamond": "http://localhost:8080/api/head/b63912df6540359774fb5cb4546a2eea1736f3fc7cf2848421697c1be8a5361",
			"gold":    "http://localhost:8080/api/head/3c9548a25e68332e6dde60816d811242fa40da58a5e24a5233e0079d9c57f779",
			"silver":  "http://localhost:8080/api/head/d086c5700ec707d703cfb45ab8afee5269a59eedda516a2c68f3e2aef7fa6a94",
		},
	},
	"STEAMING_HOT_FLOUNDER": {
		DisplayName: "Steaming-Hot Flounder",
		Description: "§7Caught when the bobber is within §a2 §7blocks of a Geyser in the §6Blazing Volcano§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/8b88f88f3053c434660eeb4c7b2344bc21ab52596cea5a66d0f9db8c0e050209",
			"diamond": "http://localhost:8080/api/head/c6602a15cf491f76584179221ed1da25fe6918f9100b864b39ea6493734809d1",
			"gold":    "http://localhost:8080/api/head/305cb6623837195113c04409658f9e8872a05e2321b79dfd1ce48eda6f749b90",
			"silver":  "http://localhost:8080/api/head/6887f3db9e1f28f08fed4fdc8be40dbeef7a04f026d5e30041ecd49a78558efb",
		},
	},
	"SULPHUR_SKITTER": {
		DisplayName: "Sulphur Skitter",
		Description: "§7Caught when standing within §a4 §7blocks of a Sulphur Ore.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/4fbf7111609f2ec23d9b3f285e1755b62193bd7c3d770576e2b18c48afeb0e29",
			"diamond": "http://localhost:8080/api/head/4c6eac56808a85b59d48aff59a262922a57cfa766f6c56f69c7d91fea230fa",
			"gold":    "http://localhost:8080/api/head/a79c52c2bb808d2e46e8e4e4db506f9406e9dfa20aee419ad90eacfb0216c169",
			"silver":  "http://localhost:8080/api/head/ba6e3560712f7213a428518deb66c0638269d17e90d8f31d4ade2a0acb91fd80",
		},
	},
	"VANILLE": {
		DisplayName: "Vanille",
		Description: "§7Caught when using a §aStarter Lava Rod §7with no Enchantments.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/57120222cce38d53ba69fc6540e97bff9abdbe22ba6068d4ee9af52ecc56842f",
			"diamond": "http://localhost:8080/api/head/bd6519e85f7fd69cb2d7bbd8d77018d907cab7b9ec1309eb933de78df00b63c1",
			"gold":    "http://localhost:8080/api/head/7684b553c9b045a429a775881130c5e6daa547d314f1f0255135f5bd46870e85",
			"silver":  "http://localhost:8080/api/head/7313271354dc2e5b1a720c6668f03ca7106ee439f874907f2a54083f0cb57721",
		},
	},
	"VOLCANIC_STONEFISH": {
		DisplayName: "Volcanic Stonefish",
		Description: "§7Caught in the §6Blazing Volcano§7.",
		Textures: map[string]string{
			"bronze":  "http://localhost:8080/api/head/38f89cbaa61ecc99a8321d53f070cef8414efc3eac47bf6fe143056fed9ee8",
			"diamond": "http://localhost:8080/api/head/48bec97138419aff0af6fb445cd5f8d68e30698facf46ae956cbda2331fb2284",
			"gold":    "http://localhost:8080/api/head/491156fede270a2fe49acf80d5956a0a9f312e37a0be669c3da54732a4c169c2",
			"silver":  "http://localhost:8080/api/head/f868b5727d27ace0beca1ca19c783b5375a88ce3d3956a0c589b4ba167e6c18f",
		},
	},
}

var TROPHY_FISH_STAGES = []string{
	"Bronze Hunter",
	"Silver Hunter",
	"Gold Hunter",
	"Diamond Hunter",
}

var TROPHY_FISH_TIERS = []string{
	"bronze",
	"silver",
	"gold",
	"diamond",
}
