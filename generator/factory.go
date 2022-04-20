package generator

func NewPasswordAllValues(length int) (string, error) {
	//Starting at first ascii (33 or !) up to the 94 ascii code numbers that exist
	pwd := Password{0, 94, 33, length}
	return newPassword(&pwd)
}

func newPassword(pass *Password) (string, error) {
	password, err := pass.GeneratePassword()
	if err != nil {
		return "", err
	}
	return password, nil
}
