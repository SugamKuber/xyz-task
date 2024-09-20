package utils

func ProcessImage(image []byte) (string, string, string, string) {
	return "SUV", "Red", "Toyota", "RAV4"
}

func GenerateCaption(carType, color, make, model string) string {
	return "Get this " + color + " " + carType + " " + make + " " + model + " now!"
}
