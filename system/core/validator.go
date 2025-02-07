package core

import (
	"app.modules/core/customerror"
	"app.modules/core/myfirestore"
	"strconv"
)

func (s *System) ValidateCommand(command CommandDetails) customerror.CustomError {
	switch command.CommandType {
	case In:
		return s.ValidateIn(command)
	case Out:
		return customerror.NewNil()
	case Info:
		return s.ValidateInfo(command)
	case My:
		return s.ValidateMy(command)
	case Change:
		// seatStateに依存するためChange()の中で行う。
		return customerror.NewNil()
	case Seat:
		return customerror.NewNil()
	case Report:
		return s.ValidateReport(command)
	case Kick:
		return s.ValidateKick(command)
	case Check:
		return s.ValidateCheck(command)
	case More:
		return s.ValidateMore(command)
	case Break:
		return s.ValidateBreak(command)
	case Resume:
		return s.ValidateResume(command)
	case Rank:
		return customerror.NewNil()
	}
	return customerror.NewNil()
}

func (s *System) ValidateIn(command CommandDetails) customerror.CustomError {
	if command.CommandType != In {
		return customerror.InvalidParsedCommand.New("this is not a In command.")
	}
	
	// 作業時間の値
	inputWorkMin := command.InOption.MinutesAndWorkName.DurationMin
	if inputWorkMin != 0 {
		expect := s.Configs.Constants.MinWorkTimeMin <= inputWorkMin && inputWorkMin <= s.Configs.Constants.MaxWorkTimeMin
		if !expect {
			return customerror.InvalidCommand.New("作業時間（分）は" + strconv.Itoa(s.Configs.Constants.MinWorkTimeMin) + "～" + strconv.Itoa(s.Configs.Constants.MaxWorkTimeMin) + "の値にしてください。")
		}
	}
	// 席番号
	if command.InOption.IsSeatIdSet {
		if command.InOption.SeatId < 0 {
			return customerror.InvalidCommand.New("座席番号は0以上の値にしてください。")
		}
	}
	
	// 作業名は特に制限はない
	// pass
	
	return customerror.NewNil()
}

func (s *System) ValidateInfo(command CommandDetails) customerror.CustomError {
	if command.CommandType != Info {
		return customerror.InvalidParsedCommand.New("this is not a Info command.")
	}
	
	// pass
	
	return customerror.NewNil()
}

func (s *System) ValidateMy(command CommandDetails) customerror.CustomError {
	if command.CommandType != My {
		return customerror.InvalidParsedCommand.New("this is not a My command.")
	}
	
	var isRankVisibleSet, isDefaultStudyMinSet, isFavoriteColorSet bool
	
	for _, option := range command.MyOptions {
		switch option.Type {
		case RankVisible:
			if isRankVisibleSet {
				return customerror.InvalidParsedCommand.New("more than 2 RankVisible options.")
			}
			isRankVisibleSet = true
		case DefaultStudyMin:
			if isDefaultStudyMinSet {
				return customerror.InvalidParsedCommand.New("more than 2 DefaultStudyMin options.")
			}
			inputDefaultStudyMin := option.IntValue
			if inputDefaultStudyMin != 0 {
				expect := s.Configs.Constants.MinWorkTimeMin <= inputDefaultStudyMin && inputDefaultStudyMin <= s.Configs.Constants.MaxWorkTimeMin
				if !expect {
					return customerror.InvalidCommand.New("作業時間（分）は" + strconv.Itoa(s.Configs.Constants.MinWorkTimeMin) + "～" + strconv.Itoa(s.Configs.Constants.MaxWorkTimeMin) + "の値にしてください。")
				}
			}
			isDefaultStudyMinSet = true
		case FavoriteColor:
			if isFavoriteColorSet {
				return customerror.InvalidParsedCommand.New("more than 2 FavoriteColor options.")
			}
			// 「color=」、つまり空欄の場合は-1として扱う。
			expect := option.IntValue == -1 || 0 <= option.IntValue
			if !expect {
				return customerror.InvalidParsedCommand.New(FavoriteColorMyOptionPrefix + "の値は設定したい色になる累計時間（0以上）を指定してください。リセットする場合は空欄にしてください。")
			}
			isFavoriteColorSet = true
		default:
			return customerror.InvalidParsedCommand.New("there is an unknown option in command.MyOptions")
		}
	}
	return customerror.NewNil()
}

func (s *System) ValidateKick(command CommandDetails) customerror.CustomError {
	if command.CommandType != Kick {
		return customerror.InvalidParsedCommand.New("this is not a Kick command.")
	}
	
	// 指定座席番号
	if command.KickOption.SeatId <= 0 {
		return customerror.InvalidCommand.New("席番号は1以上にしてください。")
	}
	
	return customerror.NewNil()
}

func (s *System) ValidateCheck(command CommandDetails) customerror.CustomError {
	if command.CommandType != Check {
		return customerror.InvalidParsedCommand.New("this is not a Check command.")
	}
	
	// 指定座席番号
	if command.CheckOption.SeatId <= 0 {
		return customerror.InvalidCommand.New("席番号は1以上にしてください。")
	}
	
	return customerror.NewNil()
}

func (s *System) ValidateReport(command CommandDetails) customerror.CustomError {
	if command.CommandType != Report {
		return customerror.InvalidParsedCommand.New("this is not a Report command.")
	}
	
	// 空欄でないか
	if command.ReportOption.Message == "" {
		return customerror.InvalidCommand.New(ReportCommand + "の右にスペースを空けてメッセージを書いてください。")
	}
	
	return customerror.NewNil()
}

func (s *System) ValidateChange(command CommandDetails, seatState myfirestore.SeatState) customerror.CustomError {
	if command.CommandType != Change {
		return customerror.InvalidParsedCommand.New("this is not a Change command.")
	}
	
	// オプションが1つ以上指定されているか
	if command.ChangeOption.NumOptionsSet() == 0 {
		return customerror.InvalidCommand.New("オプションを指定してください")
	}
	
	switch seatState {
	case myfirestore.WorkState:
		// 作業内容
		// pass
		
		// 入室時間
		if command.ChangeOption.IsDurationMinSet {
			inputDurationMin := command.ChangeOption.DurationMin
			expect := s.Configs.Constants.MinWorkTimeMin <= inputDurationMin && inputDurationMin <= s.Configs.Constants.MaxWorkTimeMin
			if !expect {
				return customerror.InvalidCommand.New("作業時間（分）は" + strconv.Itoa(s.Configs.Constants.MinWorkTimeMin) + "～" + strconv.Itoa(s.Configs.Constants.MaxWorkTimeMin) + "の値にしてください。")
			}
		}
	case myfirestore.BreakState:
		// 休憩内容
		// pass
		
		// 休憩時間
		if command.ChangeOption.IsDurationMinSet {
			inputDurationMin := command.ChangeOption.DurationMin
			expect := s.Configs.Constants.MinBreakDurationMin <= inputDurationMin && inputDurationMin <= s.Configs.Constants.MaxBreakDurationMin
			if !expect {
				return customerror.InvalidCommand.New("休憩時間（分）は" + strconv.Itoa(s.Configs.Constants.MinBreakDurationMin) + "～" + strconv.Itoa(s.Configs.Constants.MaxBreakDurationMin) + "の値にしてください。")
			}
		}
	}
	
	return customerror.NewNil()
}

func (s *System) ValidateMore(command CommandDetails) customerror.CustomError {
	if command.CommandType != More {
		return customerror.InvalidParsedCommand.New("this is not a More command.")
	}
	
	// 時間オプション
	if command.MoreOption.DurationMin <= 0 {
		return customerror.InvalidCommand.New("延長時間（分）は1以上の値にしてください。")
	}
	
	return customerror.NewNil()
}

func (s *System) ValidateBreak(command CommandDetails) customerror.CustomError {
	if command.CommandType != Break {
		return customerror.InvalidParsedCommand.New("this is not a More command.")
	}
	
	// 休憩内容
	// pass
	
	// 休憩時間
	if command.BreakOption.IsDurationMinSet {
		inputDurationMin := command.BreakOption.DurationMin
		expect := s.Configs.Constants.MinBreakDurationMin <= inputDurationMin && inputDurationMin <= s.Configs.Constants.MaxBreakDurationMin
		if !expect {
			return customerror.InvalidCommand.New("休憩時間（分）は" + strconv.Itoa(s.Configs.Constants.MinBreakDurationMin) + "～" + strconv.Itoa(s.Configs.Constants.MaxBreakDurationMin) + "の値にしてください。")
		}
	}
	
	return customerror.NewNil()
}

func (s *System) ValidateResume(command CommandDetails) customerror.CustomError {
	if command.CommandType != Resume {
		return customerror.InvalidParsedCommand.New("this is not a Resume command.")
	}
	
	// 作業名
	// pass
	
	return customerror.NewNil()
}
