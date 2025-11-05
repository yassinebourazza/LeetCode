function climbStairs(n) {
	let save1= 1
	let save2 = 2
	let result;
	if (n === 1) {
		return 1
	} else if (n === 2 ) {
		return 2 
	}
    for (let i = 3 ; i <= n; i++) {
		result = save1 + save2
		save1 = save2
		save2 = result
	}
	return result
};
