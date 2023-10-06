package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n > 0 {
		return n * n * n, nil
	}
	return 0, errors.New("edge length less than 1 is not allowed")

}
