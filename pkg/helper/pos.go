package helper

func GetCenterPos(windowWidth, windowHeight, width, height int32) (posX int32, posY int32) {
	return (windowWidth - width) / 2, (windowHeight - height) / 2
}
