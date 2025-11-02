console.log(strStr("adbutsad","sad"))
console.log(strStr("leetcode","leeto"))

function strStr(haystack, needle) {
    for (let i=0 ; i <= haystack.length-needle.length; i++) {
        console.log(haystack.slice(i,i+needle.length), '--',needle)
        if (haystack.slice(i,i+needle.length) === needle) {
            return i
        }
    }
    return -1
};