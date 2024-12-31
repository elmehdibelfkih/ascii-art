package config

const Content = ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_'abcdefjhijklmnopqrstuvwxyz{|}~`
var Standard =  make(map[rune][]string)
var Shadow =  make(map[rune][]string)
var Thinkertoy =  make(map[rune][]string)
const CharacterHeight = 8