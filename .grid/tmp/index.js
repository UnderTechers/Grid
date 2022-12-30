var jsDiff = require("diff")
const str1 = 'guanlanluditie';
const str2 = 'smartguanlanluditie';
const diffArr = jsDiff.diffChars(str1, str2);

console.log(diffArr)