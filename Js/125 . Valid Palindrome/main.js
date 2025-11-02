console.log(isPalindrome("A man, a plan, a canal: Panama"))
console.log(isPalindrome("race a car"))

function isPalindrome(s) {
    let regex = s.toLowerCase().match(/[a-zA-Z0-9]/g)
    for (let i=0 ; i < regex.length/2 ;i++) {
        if (regex[i] != regex[regex.length-1-i]) {
            return false
        }
    }
    return true
}
