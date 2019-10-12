package namecalculator

import "image"

type BasicNameCalculator struct{}

func (bnc *BasicNameCalculator) Rename(img image.Image) (string, error) {
	// h := md5.New()
	// parts := strings.Split(oldName, ".")
	// name := strings.Join(parts[:len(parts)-1], ".")
	// extension := parts[len(parts)-1]

	// imgBytes, err := ioutil.ReadFile(path + oldName)
	// if err != nil {
	// 	fmt.Printf("%s%s.%s not found\n", path, name, extension)
	// 	return "", err
	// }

	// io.WriteString(h, string(imgBytes))
	// hashString := hex.EncodeToString(h.Sum(nil))
	// return hashString, nil

	return "", nil
}
