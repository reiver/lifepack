package main



var fontSpace = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' '}


var fontPeriod = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' '}


var fontComma = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X',' ',' ',' ',' '}


var fontColon = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' '}


var fontSemicolon = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X',' ',' ',' ',' '}





var fontA = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X',' ','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X','X','X','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' '}


var fontB = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X','X','X',' ',' '}


var fontC = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var fontD = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X','X','X',' ',' '}


var fontE = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' '}


var fontF = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' '}


var fontG = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X','X',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ','X','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X','X',' '}


var fontH = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ','X',
	'X',' ',' ',' ',' ','X',
	'X',' ',' ',' ',' ','X',
	'X','X','X','X','X','X',
	'X',' ',' ',' ',' ','X',
	'X',' ',' ',' ',' ','X',
	'X',' ',' ',' ',' ','X'}


var fontI = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X','X','X',' ',' '}


var fontJ = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var fontK = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ','X',' ',' ',
	'X',' ','X',' ',' ',' ',
	'X','X',' ',' ',' ',' ',
	'X',' ','X',' ',' ',' ',
	'X',' ',' ','X',' ',' ',
	'X',' ',' ',' ','X',' '}


var fontL = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' '}


var fontM = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X','X',' ','X','X',' ',
	'X',' ','X',' ','X',' ',
	'X',' ','X',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' '}


var fontN = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X',' ',' ','X',' ',
	'X',' ','X',' ','X',' ',
	'X',' ',' ','X','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' '}


var fontO = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var fontP = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' '}


var fontQ = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ','X',' ','X',' ',
	'X',' ',' ','X',' ',' ',
	' ','X','X',' ','X',' '}


var fontR = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X','X','X','X',' ',' ',
	'X',' ','X',' ',' ',' ',
	'X',' ',' ','X',' ',' ',
	'X',' ',' ',' ','X',' '}


var fontS = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	' ',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var fontT = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' '}


var fontU = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var fontV = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X',' ','X',' ',' ',
	' ',' ','X',' ',' ',' '}


var fontW = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ','X',' ','X',' ',
	'X',' ','X',' ','X',' ',
	'X','X',' ','X','X',' ',
	'X',' ',' ',' ','X',' '}


var fontX = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X',' ','X',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X',' ','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' '}


var fontY = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X',' ','X',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' '}


var fontZ = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ','X',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' '}





var font0 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ','X','X',' ',
	'X',' ','X',' ','X',' ',
	'X','X',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var font1 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','x','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X','X','X',' ',' '}


var font2 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ','X','X',' ',' ',
	' ','X',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' '}


var font3 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ','X',' ',' ',
	' ',' ','X','X',' ',' ',
	' ',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var font4 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ',' ','X',' ',' ',
	' ',' ','X','X',' ',' ',
	' ','X',' ','X',' ',' ',
	'X',' ',' ','X',' ',' ',
	'X','X','X','X','X',' ',
	' ',' ',' ','X',' ',' ',
	' ',' ',' ','X',' ',' '}


var font5 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var font6 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ',' ','X','X','X',' ',
	' ','X',' ',' ',' ',' ',
	'X',' ',' ',' ',' ',' ',
	'X','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var font7 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	'X','X','X','X','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ','X',' ',' ',
	' ',' ','X',' ',' ',' ',
	' ','X',' ',' ',' ',' ',
	' ','X',' ',' ',' ',' ',
	' ','X',' ',' ',' ',' '}


var font8 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X',' ',' '}


var font9 = [48]rune{
	' ',' ',' ',' ',' ',' ',
	' ','X','X','X',' ',' ',
	'X',' ',' ',' ','X',' ',
	'X',' ',' ',' ','X',' ',
	' ','X','X','X','X',' ',
	' ',' ',' ',' ','X',' ',
	' ',' ',' ','X',' ',' ',
	'X','X','X',' ',' ',' '}