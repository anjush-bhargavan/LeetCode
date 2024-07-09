func garbageCollection(garbage []string, travel []int) int {
	M, G, P := 0, 0, 0
    travels := 0
    mTravel,pTravel,gTravel := 0,0,0

	for i, item := range garbage {
		mFlag, pFlag, gFlag := false, false, false
		for _, s := range item {
			switch s {
			case 'M':
				M++
				mFlag = true
			case 'G':
				G++
				gFlag = true
			case 'P':
				P++
				pFlag = true
			}
		}
		if i-1 >= 0 {
            travels += travel[i-1]
			if mFlag {
				mTravel= travels
			}

			if gFlag {
				gTravel = travels
			}

			if pFlag {
				pTravel = travels
			}
		}
	}
	return M + G + P + mTravel + gTravel + pTravel
}