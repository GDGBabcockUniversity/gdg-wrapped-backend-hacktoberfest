package util

import "github.com/philip-edekobi/wrapped/types"

var MemberPerfMap = map[int]string{
	2078: "Oladosu Emmanuel - Dosu",
	1082: "Edekobi Philip - luxury.dev",
	682:  "Abolo Samuel - That Guy",
	637:  "Ajao Rotimi Favour - Ajao Rotimi",
	571:  "Onuada Alfred - M.I",
	508:  "Ikeaba Adrian - GHØßT",
	493:  "Damilola Soji-Oderinde - Soji jr",
	476:  "Igwedinma Divine - 🐼.is.Him",
	372:  "Akpotohwo Neku - Chukwuneku",
	297:  "Onofiok Lilian - Onofiok",
}

var MemberMessageMap = map[string]string{
	"BACKEND":       "Edekobi Philip (luxury.dev)",
	"BLOCKCHAIN":    "Ifeanyichukwu John (coder)",
	"CLC":           "Damilola Soji-Oderinde (Soji jr)",
	"CLOUD":         "Iyalla Ibitein (tein)",
	"CYBERSECURITY": "Oderinde Tolu (Issa)",
	"DESIGN":        "Asoegwu Stephanie (Stephanie)",
	"GENERAL":       "Oladosu Emmanuel (Dosu)",
	"ML":            "Abolo Samuel (That Guy)",
	"MOBILE":        "Okpechi Michael (Unknown)",
	"PM":            "Idumu Emmanuel (Emmy)",
	"WEB":           "Onuada Alfred (M.I)",
}

func AdaptMembers(members types.MemStruct) []types.ReadableMem {
	nums := []int{
		members.A,
		members.B,
		members.C,
		members.D,
		members.E,
		members.F,
		members.G,
		members.H,
		members.I,
		members.J,
	}
	arr := []types.ReadableMem{}

	for _, num := range nums {
		arr = append(arr, types.ReadableMem{
			Name:     MemberPerfMap[num],
			Messages: num,
		})
	}

	return arr
}

func AdaptMemsPerTrack(mems []types.ActiveMember) []types.ReadableActiveMember {
	arr := []types.ReadableActiveMember{}

	for _, mem := range mems {
		arr = append(arr, types.ReadableActiveMember{
			Name:         MemberMessageMap[mem.GC],
			MessageCount: mem.MessageCount,
			Phone:        mem.Phone,
			GC:           mem.GC,
		})
	}

	return arr
}

func TransFormGeneralInfo(info *types.GeneralInfo) *types.ReadableGeneralInfo {
	newinfo := &types.ReadableGeneralInfo{
		MostActiveGroup:           info.MostActiveGroup,
		MostActiveTime:            info.MostActiveTime,
		MostActiveMembers:         AdaptMembers(info.MostActiveMembers),
		MostActiveMembersPerTrack: AdaptMemsPerTrack(info.MostActiveMembersPerTrack),
	}

	return newinfo
}
