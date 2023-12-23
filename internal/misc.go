package internal

type State int

func GetLastKey(m map[string]string) string {
	i := 0

	for k := range m {
		if i == len(m)-1 {
			return k
		}

		i++
	}

	return ""
}
