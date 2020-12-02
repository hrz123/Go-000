package Week02

func DealWithUser(userID int64) (*UserInfo, error) {
	userInfo, err := GetUserFromDB(userID)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
