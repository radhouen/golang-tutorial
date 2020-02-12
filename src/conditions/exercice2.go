package main

func profile(exp string) (l string) {
	switch exp {
	case "j1", "j2", "j3":
		return "you are a junior profile"
		break
	case "c1", "c2", "c3":
		return "you are a confirmed profile"
		break
	case "s1", "s2", "s3":
		return "you are a senior profile"
		break
	case "m1", "m2", "m3":
		return "you are a manager profile"
		break
	case "d1", "d2", "d3":
		return "you are a Director  profile"
		break
	default:
		return ""

	}
	return
}
